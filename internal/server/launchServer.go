package server

import (
	"fmt"
	"github.com/Xplit495/hangman-classic/util"
	"hangman-web/internal/game"
	"hangman-web/pkg/utils/preRequistiesGame"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type TemplateData struct {
	DynamicText   string
	Username      string
	GoodWarning   string
	BadWarning    string
	LetterHistory []string
	WordHistory   []string
}

var (
	wordPartiallyReveal []string
	arrSelectWord       []string
	letterHistory       []string
	wordHistory         []string
	username            string
	goodWarning         string
	badWarning          string
	liveJose            = 10
)

func LaunchServer() {

	util.ClearTerminal()

	wd, _ := os.Getwd()

	fileServer := http.FileServer(http.Dir(wd + "\\web"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" {
			http.ServeFile(writer, request, wd+"\\web\\html\\index.html")
		} else {
			fileServer.ServeHTTP(writer, request)
		}
	})

	http.HandleFunc("/submit", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			request.ParseForm()
			username = request.FormValue("username")
			difficulty := request.FormValue("difficulty")
			wordPartiallyReveal, arrSelectWord = preRequistiesGame.PreRequistiesGame(difficulty, username)
			fmt.Println("Le mot est: ", arrSelectWord)
			http.Redirect(writer, request, "/gamePage.html", http.StatusSeeOther)
		}
	})

	http.HandleFunc("/gamePage.html", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			request.ParseForm()
			proposition := request.FormValue("proposition")
			letterHistory, wordHistory, wordPartiallyReveal, liveJose, goodWarning, badWarning = game.CheckProposition(liveJose, proposition, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, goodWarning, badWarning)
		}

		data := TemplateData{
			DynamicText:   strings.Join(wordPartiallyReveal, ""),
			Username:      username,
			LetterHistory: letterHistory,
			WordHistory:   wordHistory,
			GoodWarning:   goodWarning,
			BadWarning:    badWarning,
		}

		tmpl, _ := template.ParseFiles(wd + "\\web\\html\\gamePage.html")
		tmpl.Execute(writer, data)
	})

	openBrowser("http://localhost:8080/")

	http.ListenAndServe(":8080", nil)

}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
