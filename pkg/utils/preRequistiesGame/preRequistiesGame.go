package preRequistiesGame

import (
	"github.com/Xplit495/hangman-classic/util"
)

func PreRequistiesGame(difficulty string, username string) ([]string, []string) {
	util.ClearTerminal()
	absolutePath := util.SelectDictionnaryPath(difficulty)
	arrSelectWord := util.SelectRandomWordIntoDictionnary(absolutePath)
	clues := util.GenerateWordClue(arrSelectWord)
	wordPartiallyReveal := util.AssociateClueToWord(clues, arrSelectWord)
	return wordPartiallyReveal, arrSelectWord
}
