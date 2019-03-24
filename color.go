package main

type Color int

const (
	Red Color = iota + 1
	Blue
	Yellow
)

var colorText = map[Color]string{
	Red:    "赤",
	Blue:   "青",
	Yellow: "黄",
}

func (c Color) String() string {
	return colorText[c]
}
