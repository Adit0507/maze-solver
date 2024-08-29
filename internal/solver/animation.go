package solver

import (
	"image"
	plt "image/color/palette"
	"golang.org/x/image/draw"
)

func (s *Solver) countExplorablePixels() int {
	explorablePixels := 0

	for row := 0; row < s.maze.Bounds().Dy(); row++ {
		for col := 0; col < s.maze.Bounds().Dx(); col++ {
			if s.maze.RGBAAt(col, row) != s.pallete.wall {
				explorablePixels++
			}
		}
	}

	return explorablePixels
}

// registers positions as explored on the image,
// and, if a threshold is reached, adds the frame to the output GIF.
func (s *Solver) registerExploredPixels() {
	const totalExpectedFrames = 30

	explorablePixels := s.countExplorablePixels()
	pixelsExplored := 0

	for {
		select {
		case <-s.quit:
			return

		case pos := <-s.exploredPixels:
			s.maze.Set(pos.X, pos.Y, s.pallete.explored)
			pixelsExplored++
			if pixelsExplored%(explorablePixels/totalExpectedFrames) == 0 {
				s.drawCurrentFrametoGIF()
			}
		}
	}
}

// drawCurrentFrameToGIF adds the current state of the maze as a frame of the animation.
func (s *Solver) drawCurrentFrametoGIF() {
	const (
		gifWidth      = 500
		frameDuration = 20
	)

	frame := image.NewPaletted(image.Rect(0, 0, gifWidth, gifWidth*s.maze.Bounds().Dy()/s.maze.Bounds().Dx()), plt.Plan9)

	draw.NearestNeighbor.Scale(frame, frame.Rect, s.maze, s.maze.Bounds(), draw.Over, nil)
	
	s.animation.Image = append(s.animation.Image, frame)
	s.animation.Delay = append(s.animation.Delay, frameDuration)
}
