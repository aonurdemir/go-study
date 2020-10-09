package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	gif := make([][]uint8, dy)

	for y := range gif {
		gif[y] = make([]uint8, dx)

		for x := range gif[y] {

			gif[y][x] = uint8((x + y) / 2)
		}

	}

	return gif
}

func main() {
	pic.Show(Pic)
}
