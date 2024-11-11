package utils

import "github.com/muesli/termenv"

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getHeight() int {
	_, height, err := termenv.Size()
	if err != nil {
		panic(err)
	}

	return height
}
