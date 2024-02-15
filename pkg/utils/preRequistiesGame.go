package utils

// Defines a package for utility functions used across the game.

import (
	"github.com/Xplit495/hangman-classic/util" // Imports the utility package from the hangman-classic repository.
)

// PreRequistiesGame initializes the game based on the selected difficulty.
func PreRequistiesGame(difficulty string) ([]string, []string) {
	util.ClearTerminal()                                                  // Clears the terminal for a clean start.
	absolutePath := util.SelectDictionnaryPath(difficulty)                // Determines the path to the dictionary based on difficulty.
	arrSelectWord := util.SelectRandomWordIntoDictionnary(absolutePath)   // Selects a random word from the specified dictionary.
	clues := util.GenerateWordClue(arrSelectWord)                         // Generates initial clues for the selected word.
	wordPartiallyReveal := util.AssociateClueToWord(clues, arrSelectWord) // Associates clues with the selected word for display.
	return wordPartiallyReveal, arrSelectWord                             // Returns the partially revealed word and the array of the selected word.
}
