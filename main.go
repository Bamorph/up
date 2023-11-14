package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "9000", "Port to listen on")
	// directory := flag.String("d", ".", "Directory to serve files from")
	flag.Parse()

	server := &http.Server{
		Addr:    ":" + *port,
		Handler: http.FileServer(http.Dir(".")),
	}

	log.Printf("Serving files from current directory on port %s\n", *port)
	log.Fatal(server.ListenAndServe())
}
