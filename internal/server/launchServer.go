package server

import (
	"fmt"
	"github.com/Xplit495/hangman-classic/util"
	"hangman-web/pkg/utils/preRequistiesGame"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func LaunchServer() {

	var wordPartiallyReveal []string

	type TemplateData struct {
		DynamicText string
	}

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
			wordPartiallyReveal = preRequistiesGame.PreRequistiesGame(difficulty, username)
			fmt.Println("Le mot à trouver est: ", wordPartiallyReveal)
			http.Redirect(w, r, "html/gamePage.html", http.StatusSeeOther)
		}
	})

	tmpl, err := template.ParseFiles(wd + "\\web\\html\\index.html")
	if err != nil {
		log.Fatal("Erreur lors de l'analyse du template :", err)
	}

	// Modifie le handler pour la route "/"
	http.HandleFunc("/gamePage.html", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/gamePage.html" {
			// Crée les données dynamiques à afficher
			data := TemplateData{
				DynamicText: strings.Join(wordPartiallyReveal, ""),
			}
			// Exécute le template avec les données dynamiques
			err1 := tmpl.Execute(w, data)
			if err1 != nil {
				http.Error(w, "Erreur lors de l'exécution du template", http.StatusInternalServerError)
			}
		} else {
			fileServer.ServeHTTP(w, r)
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
