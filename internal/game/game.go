package game

// Defines the package for game logic.

import (
	"github.com/Xplit495/hangman-classic/util" // Imports utility functions for the game.
	"strconv"                                  // Used for converting integers to strings.
)

// CheckProposition evaluates the player's guess and updates the game state.
func CheckProposition(liveJose int, choice string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, goodWarning string, badWarning string) ([]string, []string, []string, int, string, string) {
	// Initialize variables to track the state of the guess.
	var (
		inputValidation      bool     // Checks if the input is valid.
		letterFind           bool     // Indicates if a letter guess is correct.
		wordFind             bool     // Indicates if a word guess is correct.
		letterAlreadyUse     bool     // Checks if a letter has already been guessed.
		wordAlreadyUse       bool     // Checks if a word has already been guessed.
		choiceToLowerStrings []string // Processes the guess to compare with the word.
	)
	// Reset warnings for each new proposition.
	badWarning = ""
	goodWarning = ""
	// Validates the input and updates history based on the guess.
	inputValidation, letterHistory, wordHistory, choiceToLowerStrings = util.StartGame(choice, wordPartiallyReveal, letterHistory, wordHistory)

	// Process the guess if it's valid.
	if inputValidation {
		// Update the game state based on the guess.
		wordPartiallyReveal, letterFind, wordFind = util.UpdateWord(arrSelectWord, wordPartiallyReveal, choiceToLowerStrings)
		// Update the guess histories and check if the guess was already made.
		letterHistory, wordHistory, letterAlreadyUse, wordAlreadyUse = util.UpdateHistroy(letterHistory, wordHistory, choiceToLowerStrings)
		// Set warning if the guess was already made.
		if letterAlreadyUse || wordAlreadyUse {
			badWarning = "Attention vous avez déjà essayé ce mot ou cette lettre !"
		} else {
			// Adjust lives based on the correctness of the guess.
			if !wordFind && len(choiceToLowerStrings) > 1 {
				liveJose -= 2 // Penalize wrong word guesses more heavily.
			} else if !letterFind {
				liveJose-- // Deduct a life for a wrong letter guess.
			} else {
				goodWarning = "Correct !" // Praise for a correct guess.
			}
		}
	} else {
		// Warn if the input is invalid.
		badWarning = "Merci de saisir un mot de la même longueur OU une lettre de l'alphabet !"
	}
	// Return the updated game state.
	return letterHistory, wordHistory, wordPartiallyReveal, liveJose, goodWarning, badWarning
}

// ChoosePathJosePosition determines the image path based on remaining lives.
func ChoosePathJosePosition(liveJose int) string {
	pathJose := "../resourcesWeb/" + strconv.Itoa(10-liveJose) + ".png" // Construct the path based on lives left.
	return pathJose
}
