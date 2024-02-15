package server

import (
	"github.com/Xplit495/hangman-classic/util"
	"hangman-web/internal/game"
	"hangman-web/pkg/utils"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type TemplateData struct {
	Word          string
	Username      string
	GoodWarning   string
	BadWarning    string
	ImagePath     string
	LetterHistory string
	WordHistory   string
	ArrSelectWord string
}

var (
	wordPartiallyReveal []string
	arrSelectWord       []string
	letterHistory       []string
	wordHistory         []string
	username            string
	goodWarning         string
	badWarning          string
	imagePath           string
	liveJose            = 10
)

func LaunchServer() {
	util.ClearTerminal() // Clears the console for clarity before server starts.

	wd, _ := os.Getwd() // Gets the current working directory, ignoring errors.

	fileServer := http.FileServer(http.Dir(wd + "\\web")) // Serves static files from the 'web' directory.

	// Handles HTTP requests to the root ("/") path.
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" {
			http.ServeFile(writer, request, wd+"\\web\\html\\index.html") // Serves the main page.
		} else {
			fileServer.ServeHTTP(writer, request) // Serves static files for any other request.
		}
	})

	http.HandleFunc("/submit", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost { // Ensures that the form submission is a POST request.
			err := request.ParseForm() // Parses the form values from the request.
			if err != nil {
				return // Exits the function if there's an error parsing the form.
			}
			username = request.FormValue("username")      // Retrieves the username from the form submission.
			difficulty := request.FormValue("difficulty") // Retrieves the game difficulty selected by the user.
			// Initializes the game based on the selected difficulty and updates global variables.
			wordPartiallyReveal, arrSelectWord = utils.PreRequistiesGame(difficulty)
			// Redirects the user to the game page after processing the form submission.
			http.Redirect(writer, request, "/gamePage.html", http.StatusSeeOther)
		}
	})

	http.HandleFunc("/gamePage.html", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost { // Process only POST requests to update game state based on user's guess.
			err := request.ParseForm() // Parse submitted form data.
			if err != nil {
				return // Exit if there's an error parsing the form.
			}
			proposition := request.FormValue("proposition") // Get user's guess from the form.
			// Update game state based on the proposition and current game data.
			letterHistory, wordHistory, wordPartiallyReveal, liveJose, goodWarning, badWarning = game.CheckProposition(liveJose, proposition, arrSelectWord, wordPartiallyReveal, letterHistory, wordHistory, goodWarning, badWarning)
		}

		counter := 0 // Count correctly guessed letters.
		for i := 0; i < len(arrSelectWord); i++ {
			if wordPartiallyReveal[i] == arrSelectWord[i] {
				counter++
			}
		}
		// Redirect to winning screen if all letters are correctly guessed and lives are remaining.
		if counter == len(arrSelectWord) && liveJose > 0 {
			http.Redirect(writer, request, "winningScreen.html", http.StatusSeeOther)
		}
		// Redirect to losing screen if no lives are remaining.
		if liveJose <= 1 {
			http.Redirect(writer, request, "/losingScreen.html", http.StatusSeeOther)
		}

		imagePath = game.ChoosePathJosePosition(liveJose) // Update hangman image based on remaining lives.
		// Prepare data for rendering the game page template.
		data := TemplateData{
			Word:          strings.Join(wordPartiallyReveal, " "),
			Username:      username,
			LetterHistory: strings.Join(letterHistory, ","),
			WordHistory:   strings.Join(wordHistory, ","),
			GoodWarning:   goodWarning,
			BadWarning:    badWarning,
			ImagePath:     imagePath,
		}

		tmpl, _ := template.ParseFiles(wd + "\\web\\html\\gamePage.html") // Parse the game page template.
		err := tmpl.Execute(writer, data)                                 // Render the template with current game state data.
		if err != nil {
			return // Exit if template execution fails.
		}
	})

	http.HandleFunc("/winningScreen.html", func(writer http.ResponseWriter, request *http.Request) {
		// Prepare data for the winning screen, including game state and outcome details.
		data := TemplateData{
			Word:          strings.Join(wordPartiallyReveal, " "), // Show the fully revealed word.
			Username:      username,                               // Display the player's username.
			LetterHistory: strings.Join(letterHistory, ","),       // Show history of guessed letters.
			WordHistory:   strings.Join(wordHistory, ","),         // Show history of guessed words.
			GoodWarning:   goodWarning,                            // Display any success message.
			BadWarning:    badWarning,                             // Display any warning message.
			ImagePath:     imagePath,                              // Path to the image for the current state of the hangman.
			ArrSelectWord: strings.Join(arrSelectWord, ""),        // The correct word as a single string.
		}

		// Parse the winning screen HTML template.
		tmpl, _ := template.ParseFiles(wd + "\\web\\html\\winningScreen.html")
		err := tmpl.Execute(writer, data) // Render the template with the prepared data.
		if err != nil {
			return // Exit the function if there's an error rendering the template.
		}
	})

	http.HandleFunc("/losingScreen.html", func(writer http.ResponseWriter, request *http.Request) {
		// Set up the data to be displayed on the losing screen, encapsulating the final game state.
		data := TemplateData{
			Word:          strings.Join(wordPartiallyReveal, " "), // Display the word with spaces between letters for readability.
			Username:      username,                               // Show the user who played the game.
			LetterHistory: strings.Join(letterHistory, ","),       // List all letters guessed during the game.
			WordHistory:   strings.Join(wordHistory, ","),         // List all words attempted during the game.
			GoodWarning:   goodWarning,                            // Any final messages for correct guesses.
			BadWarning:    badWarning,                             // Any final messages for incorrect guesses.
			ImagePath:     imagePath,                              // The final hangman image indicating loss.
			ArrSelectWord: strings.Join(arrSelectWord, ""),        // The correct word, shown as a single string.
		}

		// Parse the HTML template for the losing screen.
		tmpl, _ := template.ParseFiles(wd + "\\web\\html\\losingScreen.html")
		err := tmpl.Execute(writer, data) // Render the losing screen with the game's concluding data.
		if err != nil {
			return // Exit if there's an error in rendering the template.
		}
	})

	http.HandleFunc("/restart", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost { // Ensure this handler responds only to POST requests.
			// Reset the game state to start a new game.
			wordPartiallyReveal = nil
			arrSelectWord = nil
			letterHistory = nil
			wordHistory = nil
			liveJose = 10 // Reset lives.
			goodWarning = ""
			badWarning = ""
			// Surprise.SS
			err := utils.OpenBrowser("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
			if err != nil {
				return // Exit if unable to open the browser.
			}
			// Redirect the user back to the index page to start a new game.
			http.Redirect(writer, request, "/index.html", http.StatusSeeOther)
		}
	})

	http.HandleFunc("/restart1", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost { // Similar to "/restart" but might be intended for a different outcome or version.
			// Resets the game state for another attempt or version of the game.
			wordPartiallyReveal = nil
			arrSelectWord = nil
			letterHistory = nil
			wordHistory = nil
			liveJose = 10 // Reset the counter for lives.
			goodWarning = ""
			badWarning = ""
			// Surprise.
			err := utils.OpenBrowser("https://www.youtube.com/watch?v=4xnsmyI5KMQ")
			if err != nil {
				return // Exit if there's an issue opening the browser.
			}
			// Redirect to the starting page for a new game session.
			http.Redirect(writer, request, "/index.html", http.StatusSeeOther)
		}
	})

	// Attempt to open the game in the default web browser automatically on server start.
	err := utils.OpenBrowser("http://localhost:8080/")
	if err != nil {
		return // Exit if unable to open the browser automatically.
	}

	// Start listening for HTTP requests on port 8080.
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return // Exit if there's an error starting the HTTP server.
	}
}
