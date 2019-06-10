package main

import (
	"encoding/binary"
	"github.com/karalabe/hid"
)

// rowToByte takes single row of int flags (as binary representation) and converts
// them into the corresponding byte value. Note that the values in row are reversed!
func rowToByte(row []int) uint64 {
	value := uint64(0)
	for i := uint64(len(row)); i > 0; i-- {
		if row[i-1] != 0 {
			value |= 1 << (i - 1) // values are displayed right-to-left
		}
	}
	return value
}

// convertRows converts grid to bytes which can be sent to the device
func convertRows(grid [][]int) []uint64 {

	// The USB report has the size of 8 bytes. The packets are defined in the following format:
	//
	// byte 0 = brightness
	//      1 = row #
	//      2 = data 1st row (LEDs 17 - 21)
	//      3 = data 1st row (LEDs 9 - 16)
	//      4 = data 1st row (LEDs 1 - 8)
	//      5 = data 2nd row (LEDs 17 - 21)
	//      6 = data 2nd row (LEDs 9 - 16)
	//      7 = data 2nd row (LEDs 1 - 8)
	//
	// Remarks:
	//  a. The value of brightness ranged from 0 to 2. 0 is the maximum brightness.
	//  b. The row number should be 0, 2, 4, 6.
	//  c. As each packet contains two rows’ data, so the packet with Row # = 0 includes the
	//     data from Row 1 and Row 2.
	//  d. The LED will turn on shortly after receiving the packet. So the software should
	//     keep refreshing the device.
	//  e. For row data, 1=off and 0=on.
	//
	// For every two rows of the 21 LEDs the data will be stored in 8 bytes as follows:
	//
	// led   01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21
	// bit    0  1  2  3  4  5  6  7  0  1  2  3  4  5  6  7  0  1  2  3  4  5  6  7
	// row 1 <------ byte 4 -------> <------ byte 3 -------> <-- byte 2 -->  1  1  1
	// row 2 <------ byte 7 -------> <------ byte 6 -------> <-- byte 5 -->  1  1  1

	rows := make([]uint64, 4)
	for idx, row := range grid {
		// device will interpret 1 as off and 0 as on so we flip all bits
		converted := 0x1FFFFF ^ rowToByte(row)
		// add padding bits (111) for missing LEDs 22 to 24
		converted |= 0xe00000
		if idx%2 == 0 {
			rows[idx/2] |= converted << 24
		} else {
			rows[idx/2] |= converted
		}
	}
	// add brightness and row numbers
	rows[0] |= 0x0200000000000000
	rows[1] |= 0x0202000000000000
	rows[2] |= 0x0204000000000000
	rows[3] |= 0x0206000000000000
	return rows
}

// DisplayGrid displays an 21 x 7 grid on a Dream Cheeky Message board (device)
func DisplayGrid(grid [][]int, device *hid.Device) error {
	// send rows to device as bytes
	rows := convertRows(grid)
	for _, row := range rows {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, row)
		_, err := device.Write(buffer)
		if err != nil {
			return err
		}
	}
	return nil // everything went smoothly
}
