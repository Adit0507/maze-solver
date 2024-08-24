package main

import (
	"fmt"
	"log"
	"os"
)

func usage() {
	_, _ = fmt.Fprintln(os.Stderr,"Usage: maze-solver maze_10x10.png solution.png")
	os.Exit(1)
}	

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	log.Printf("Solving maze %q and saving it as %q", inputFile, outputFile)
}