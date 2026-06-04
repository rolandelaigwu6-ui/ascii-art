package main

import (
	//"html/template"
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// send a Response to the client
	fmt.Fprintf(w, "Welcome to my go server")
}

func main() {
	// register the handler function for the root route
	http.HandleFunc("/", homeHandler)

	fmt.Println("server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error: ", err)
	}
}

// type PageData struct {
// 	Text   string
// 	Output string
// }

// var page = template.Must(template.ParseFiles("templates/index.html"))

// func main() {
// 	http.HandleFunc("/", homeHandler)
// 	http.HandleFunc("/ascii-art", asciiArtHandler)
// 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

// 	log.Println("listening on http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	if r.Method != http.MethodGet {
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	renderPage(w, PageData{})
// }

// func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	text := r.FormValue("text")
// 	renderPage(w, PageData{
// 		Text:   text,
// 		Output: RenderMiniASCII(text),
// 	})
// }

// func renderPage(w http.ResponseWriter, data PageData) {
// 	if err := page.Execute(w, data); err != nil {
// 		http.Error(w, "template error", http.StatusInternalServerError)
// 	}
// }

// func RenderMiniASCII(text string) string {
// 	if text == "" {
// 		return ""
// 	}

// 	lines := strings.Split(text, "\n")
// 	var blocks []string

// 	for _, line := range lines {
// 		blocks = append(blocks, renderLine(line))
// 	}

// 	return strings.Join(blocks, "\n")
// }

// func renderLine(text string) string {
// 	rows := []string{"", "", ""}

// 	for _, char := range text {
// 		block := miniBlock(char)
// 		for i := 0; i < len(rows); i++ {
// 			rows[i] += block[i] + " "
// 		}
// 	}

// 	return strings.Join(rows, "\n")
// }

// func miniBlock(char rune) []string {
// 	if char == ' ' {
// 		return []string{"   ", "   ", "   "}
// 	}

// 	return []string{
// 		strings.Repeat(string(char), 3),
// 		string(char) + " " + string(char),
// 		strings.Repeat(string(char), 3),
// 	}
// }
