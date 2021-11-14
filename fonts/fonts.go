package fonts

type Font struct {
	Name  string
	Meta  map[string]string
	Chars map[int][]int
}
