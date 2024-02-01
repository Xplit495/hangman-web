package server

import (
	"fmt"
	"github.com/Xplit495/hangman-classic/util"
	"log"
	"net/http"
	"os"
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
			//username := r.FormValue("username")
			//difficulty := r.FormValue("difficulty")
		}
	})

	fmt.Println("Pour accéder au jeu -> http://localhost:8080/")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
