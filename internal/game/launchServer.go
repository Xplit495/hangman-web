package launchServer

import (
	"fmt"
	"net/http"
)

func launchSever() {
	// Définis le gestionnaire
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Bonjour, le serveur fonctionne !")
	})

	// Démarre le serveur HTTP sur le port 8080
	fmt.Println("Le serveur démarre sur le port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur :", err)
	}
}
