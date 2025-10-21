package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	//flag
	addr := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()

	//file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//logging and saving log in file
	f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//custom log info
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	//custom log error
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	//http.server stuct
	srv := &http.Server{
		Addr:     ":" + *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on :%s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
