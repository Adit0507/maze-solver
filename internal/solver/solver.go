package solver

import (
	"fmt"
	"image"
	"log"
	"sync"
)

type Solver struct {
	maze           *image.RGBA
	pallete        pallete
	pathsToExplore chan *path
	solution       *path
	mutex          sync.Mutex
}

// building the solver by opening the image
func New(imagePath string) (*Solver, error) {
	img, err := openMaze(imagePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open maze image: %w", err)
	}

	return &Solver{
		maze:           img,
		pallete:        defaultPallete(),
					// initialized it with 1, to make it buffered coz 
					// unbuffered channel cant be read from, as a
					// send operation on unbuffered channel blokcs the 
					// sending goroutine until corresponding recevie on the same channel at which point the valueis transmitted & both goroutines continue 
		pathsToExplore: make(chan *path, 1),
	}, nil
}

// finds path to the treasure
func (s *Solver) Solve() error {
	entrance, err := s.findEntrance()
	if err != nil {
		return fmt.Errorf("unable to find entrance: %w", err)
	}

	log.Printf("starting at %v", entrance)

	// writing in paths to explore before starting the chanel
	s.pathsToExplore <- &path{previousStep: nil, at: entrance}
	s.listenToBranches()

	return nil
}

// finds the one pixel that has entrance color
func (s *Solver) findEntrance() (image.Point, error) {
	for row := s.maze.Bounds().Min.Y; row < s.maze.Bounds().Max.X; row++ {
		for col := s.maze.Bounds().Min.X; col < s.maze.Bounds().Max.Y; col++ {
			if s.maze.RGBAAt(col, row) == s.pallete.entrance {
				return image.Point{X: col, Y: row}, nil
			}
		}
	}

	return image.Point{}, fmt.Errorf("entrance positon not found")
}
