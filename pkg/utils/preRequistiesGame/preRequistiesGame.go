package preRequistiesGame

import (
	"fmt"
	"github.com/Xplit495/hangman-classic/util"
)

func PreRequistiesGame(difficulty string, username string) []string {
	util.ClearTerminal()
	absolutePath := util.SelectDictionnaryPath(difficulty)
	randomWord := util.SelectRandomWordIntoDictionnary(absolutePath)
	clues := util.GenerateWordClue(randomWord)
	wordPartiallyReveal := util.AssociateClueToWord(clues, randomWord)
	fmt.Println("Bienvenue ", username)
	fmt.Println("Vous avez choisi la difficulté: ", difficulty)
	fmt.Println("Le est: ", randomWord)
	return wordPartiallyReveal
}
