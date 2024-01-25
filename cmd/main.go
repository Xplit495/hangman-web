package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		difficultyFile string
	)

	if len(os.Args) > 2 {
		fmt.Println("Trop d'arguments, fin du programme.")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		difficultyFile = os.Args[1]
	} else {
		fmt.Println("Pas de difficulté choisie, fin du programme.")
		os.Exit(1)
	}

	if difficultyFile == "easy.txt" {
		fmt.Println("Difficulté facile")
	} else if difficultyFile == "medium.txt" {
		fmt.Println("Difficulté moyenne")
	} else if difficultyFile == "hard.txt" {
		fmt.Println("Difficulté difficile")
	}
}
