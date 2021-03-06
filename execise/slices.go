package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		ret[y] = make([]uint8, dx)
	    for x := 0; x < dx; x++ {
	      ret[y][x] = uint8((x*x + y*y)/2)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}

