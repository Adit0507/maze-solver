package solver

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

// opens RGBA image from path
func openMaze(imagePath string) (*image.RGBA, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open image %s: %w", imagePath, err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("unable to load input image from %s: %w", imagePath, err)
	}

	rgbaImage, ok := img.(*image.RGBA)
	if ! ok {
		return nil, fmt.Errorf("expected RGBA image, got %T", img) 
	}

	return rgbaImage, nil
}

// saves the image as a PNG file with the solution path highlighted.
func (s *Solver) SaveSolution(outPath string) error {
	return nil
}