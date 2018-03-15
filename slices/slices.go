package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
	}
	
	// Do Image
	for i := range pic {
		for j := range pic[i] {
			switch {
				case j % 5 == 0:
					pic[i][j] = 255
			}
		}
	}
	
	return pic
}

func main() {
	pic.Show(Pic)
}