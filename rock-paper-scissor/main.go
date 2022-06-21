package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	/* html := `<strong>Hello world</strong>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, html) */
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	roundResult := rps.PlayRound(playerChoice)
	log.Println("result:", roundResult)

	out, err := json.MarshalIndent(roundResult, "", " ")
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(out))
}

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/play", playRound)

	log.Println("Starting web server")
	http.ListenAndServe(":8082", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
