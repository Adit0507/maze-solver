package solver

import "image"

type path struct {
	previousStep *path
	at           image.Point
}

func (p path) isPreviousStep(n image.Point) bool {
	return p.previousStep != nil && p.previousStep.at == n
}