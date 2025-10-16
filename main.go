package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Display a specific snippet..."))
	fmt.Fprintf(w, "Display a specific snippet with the ID: %d\n", id)
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	//manupulating header map
	w.Header().Set("Cache-Control", "Public, max-age: 34555")
	w.Header().Del("Cache-Control")
	log.Print(w.Header().Get("Content-Type"))
	log.Print(w.Header().Values("Content-Type"))
	// suppresing system generated headers
	w.Header()["Date"] = nil
	w.Write([]byte("Create a snippet..."))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet/view", snippetView)

	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on port: 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
