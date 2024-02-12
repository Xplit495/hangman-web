package game

import (
	"fmt"
	"github.com/Xplit495/hangman-classic/util"
	"os"
)

func CheckProposition(liveJose int, choice string, arrSelectWord []string, wordPartiallyReveal []string, letterHistory []string, wordHistory []string, goodWarning string, badWarning string) ([]string, []string, []string, int, string, string) {
	var (
		inputValidation      bool
		letterFind           bool
		wordFind             bool
		letterAlreadyUse     bool
		wordAlreadyUse       bool
		choiceToLowerStrings []string
	)
	badWarning = ""
	goodWarning = ""
	inputValidation, letterHistory, wordHistory, choiceToLowerStrings = util.StartGame(choice, wordPartiallyReveal, letterHistory, wordHistory)

	if inputValidation {
		wordPartiallyReveal, letterFind, wordFind = util.UpdateWord(arrSelectWord, wordPartiallyReveal, choiceToLowerStrings)
		letterHistory, wordHistory, letterAlreadyUse, wordAlreadyUse = util.UpdateHistroy(letterHistory, wordHistory, choiceToLowerStrings)
		if letterAlreadyUse || wordAlreadyUse {
			badWarning = "Attention vous avez déjà essayé ce mot ou cette lettre !"
		} else {
			if wordFind {
				os.Exit(0)
			}
			if !wordFind && len(choiceToLowerStrings) > 1 {
				liveJose -= 2
			} else if !letterFind {
				liveJose--
			} else {
				goodWarning = "Correct !"
			}
		}
		fmt.Println(liveJose)
		fmt.Println(letterFind)
	} else {
		badWarning = "Merci de saisir un mot de la même longueur OU une lettre de l'alphabet !"
	}
	return letterHistory, wordHistory, wordPartiallyReveal, liveJose, goodWarning, badWarning
}
