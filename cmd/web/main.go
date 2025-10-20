package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	addr := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()
	//file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on :%s", *addr)
	err := http.ListenAndServe(":"+*addr, mux)
	log.Fatal(err)
}
