package solver

import "image"

// returns an array of 4 neighbours of the pixel
func neighbours(p image.Point) []image.Point {
	return []image.Point{
		{p.X, p.Y + 1},
		{p.X, p.Y - 1},
		{p.X + 1, p.Y},
		{p.X - 1, p.Y},
	}
}
