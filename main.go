package main

import (
	"fmt"
	"log"
	"maze-solver/internal/solver"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	log.Printf("Solving maze %q and saving it as %q", inputFile, outputFile)

	s, err := solver.New(inputFile)
	if err != nil {
		exit(err)
	}

	err = s.Solve()
	if err != nil {
		exit(err)
	}

	err = s.SaveSolution(outputFile)
	if err != nil {
		exit(err)
	}
}

func exit(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Error: %s", err)
	os.Exit(1)
}

func usage() {
	_, _ = fmt.Fprintln(os.Stderr,"Usage: maze_solver input.png output.png")
	os.Exit(1)
}	
