package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	mux := http.NewServeMux()

	//flag
	addr := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()

	//file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	//custom log info
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	//custom log error
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	// Initialize a new instance of our application struct, containing the
	// dependencies.
	app := application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	//http.server stuct
	srv := &http.Server{
		Addr:     ":" + *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Starting server on :%s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
