package server

import (
	"fmt"
	"github.com/Xplit495/hangman-classic/util"
	"hangman-web/pkg/utils/preRequistiesGame"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type TemplateData struct {
	DynamicText string
	Username    string
}

var (
	wordPartiallyReveal []string
	username            string
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
			wordPartiallyReveal = preRequistiesGame.PreRequistiesGame(difficulty, username)
			fmt.Println("Le mot à trouver est: ", wordPartiallyReveal)
			http.Redirect(writer, request, "/gamePage.html", http.StatusSeeOther)
		}
	})

	http.HandleFunc("/gamePage.html", func(writer http.ResponseWriter, request *http.Request) {
		data := TemplateData{
			DynamicText: strings.Join(wordPartiallyReveal, ""),
			Username:    username,
		}
		tmpl, _ := template.ParseFiles(wd + "\\web\\html\\gamePage.html")
		tmpl.Execute(writer, data)

		if request.Method == http.MethodPost {
			request.ParseForm()
			proposition := request.FormValue("proposition")
			fmt.Println("La proposition est:" + proposition)
		}
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
