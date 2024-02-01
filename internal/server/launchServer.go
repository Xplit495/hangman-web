package server

import (
	"fmt"
	"github.com/Xplit495/hangman-classic/util"
	"hangman-web/pkg/utils/preRequistiesGame"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func LaunchServer() {
	util.ClearTerminal()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir(wd + "\\web"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, wd+"\\web\\html\\index.html")
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			username := r.FormValue("username")
			difficulty := r.FormValue("difficulty")
			wordPartiallyReveal := preRequistiesGame.PreRequistiesGame(difficulty, username)
			fmt.Println("Le mot à trouver est: ", wordPartiallyReveal)
			http.Redirect(w, r, "/web/html/gamePage.html", http.StatusSeeOther)
		}
	})

	err = openBrowser("http://localhost:8080/")
	if err != nil {
		fmt.Println("Impossible d'ouvrir le navigateur")
	}
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

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
