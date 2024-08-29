package solver

import "image/color"

type pallete struct {
	wall     color.RGBA
	path     color.RGBA
	entrance color.RGBA
	treasure color.RGBA
	solution color.RGBA
	explored color.RGBA
}

func defaultPallete() pallete {
	return pallete{
		wall:     color.RGBA{R: 0, G: 0, B: 0, A: 255},
		path:     color.RGBA{R: 255, G: 255, B: 255, A: 255},
		entrance: color.RGBA{R: 0, G: 191, B: 255, A: 255},
		treasure: color.RGBA{R: 255, G: 0, B: 128, A: 255},
		solution: color.RGBA{R: 255, G: 0, B: 0, A: 0},
		explored: color.RGBA{R:0 , G:100, B:100, A:220},
	}
}
