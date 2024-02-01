package preRequistiesGame

import (
	"fmt"
	"github.com/Xplit495/hangman-classic/util"
)

func PreRequistiesGame(difficulty string, username string) []string {
	absolutePath := util.SelectDictionnaryPath(difficulty)
	randomWord := util.SelectRandomWordIntoDictionnary(absolutePath)
	clues := util.GenerateWordClue(randomWord)
	wordPartiallyReveal := util.AssociateClueToWord(clues, randomWord)
	fmt.Println("Bienvenue ", username)
	fmt.Println("Vous avez choisi la difficulté: ", difficulty)
	fmt.Println("Le est: ", randomWord)
	fmt.Println("Le mot à trouver est: ", wordPartiallyReveal)
	return wordPartiallyReveal
}
