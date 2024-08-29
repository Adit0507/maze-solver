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

	for {
		select {
		case <-s.quit:
			log.Println("the treasure has been found, stopping worker")
			return

		case p := <-s.pathsToExplore:
			wg.Add(1)
			go func(p *path) {
				defer wg.Done()
				s.explore(p)
			}(p)
		}
	}
}

func (s *Solver) explore(pathToBranch *path) {
	if pathToBranch == nil {
		return
	}

	pos := pathToBranch.at
	for {
		select {
		case <-s.quit:
			return
		default:
		}

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
					close(s.quit)
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
			select {
			case <- s.quit:
				log.Printf("I'm an unlucky branch, someone else found the treasure, I give up at position %v.", pos)
				return
			case s.pathsToExplore <- branch:
				// contuinue execution after select block
			}
		}

		pathToBranch = &path{previousStep: pathToBranch, at: candidates[0]}
		pos = candidates[0]
	}

}
