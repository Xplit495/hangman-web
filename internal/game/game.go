package game

import (
	"fmt"
	"github.com/Xplit495/hangman-classic/util"
)

func CheckProposition(choice string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string) ([]string, []string, []string) {
	var (
		inputValidation      bool
		letterFind           bool
		wordFind             bool
		choiceToLowerStrings []string
	)
	inputValidation, letterHistory, wordHistory, choiceToLowerStrings = util.StartGame(choice, wordPartiallyReveal, letterHistory, wordHistory)
	if inputValidation {
		wordPartiallyReveal, letterFind, wordFind = util.UpdateWord(arrSelectWord, wordPartiallyReveal, choiceToLowerStrings)
		fmt.Println(letterFind)
		fmt.Println(wordFind)
	}
	return letterHistory, wordHistory, wordPartiallyReveal
}
