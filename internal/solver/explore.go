package solver

import (
	"image"
	"log"
	"sync"
)

// starts a new goroutine for each message in the channel
func (s *Solver) listenToBranches() {
	// adds one tracker to wait group before spinning a new gorutine
	// goroutine should then tell wait group when it is done with its work
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for p := range s.pathsToExplore {
		wg.Add(1)
		go func (path *path)  {
			defer wg.Done()
			s.explore(path)
		}(p)

		if s.solutionFound() {
			return
		}
	}
}

// returns whether solution is found or not
func (s *Solver) solutionFound() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.solution != nil
}

func (s *Solver) explore(pathToBranch *path) {
	if pathToBranch == nil {
		return
	}

	pos := pathToBranch.at
	for !s.solutionFound() {
		candidates := make([]image.Point, 0, 3)
		for _, n := range neighbours(pos) {
			if pathToBranch.isPreviousStep(n) {
				continue
			}

			switch s.maze.RGBAAt(n.X, n.Y) {
			case s.pallete.treasure:
				s.mutex.Lock()
				defer s.mutex.Unlock()
				if s.solution == nil {
					s.solution = &path{previousStep: pathToBranch, at: n}
					log.Printf("Treasure found at %v", s.solution.at)
				}
				return

			case s.pallete.path:
				candidates = append(candidates, n)
			}
		}

		if len(candidates) == 0 {
			log.Printf("I must have taken the wrong turn at position %v.", pos)
			return
		}

		for _, candidate := range candidates[1:] {
			branch := &path{previousStep: pathToBranch, at: candidate}
			s.pathsToExplore <- branch
		}

		pathToBranch = &path{previousStep: pathToBranch, at: candidates[0]}
		pos = candidates[0]
	}

}
