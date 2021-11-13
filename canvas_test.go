package main

import (
	"image"
	"testing"
)

func Test_pixelBufferLength(t *testing.T) {
	type args struct {
		r image.Rectangle
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "0x0 causes error",
			args:    args{r: image.Rect(0, 0, 0, 0)},
			want:    0,
			wantErr: true,
		},
		{
			name:    "0x5 causes error",
			args:    args{r: image.Rect(0, 0, 0, 5)},
			want:    0,
			wantErr: true,
		},
		{
			name:    "21x1",
			args:    args{r: image.Rect(0, 0, 21, 1)},
			want:    3,
			wantErr: false,
		},
		{
			name:    "21x7",
			args:    args{r: image.Rect(0, 0, 21, 7)},
			want:    19,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pixelBufferLength(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("pixelBufferLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pixelBufferLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
