package main

import (
	"fmt"
	"github.com/scriptodude/vidego/config"
	"io"
	"log"
	"net/http"
	"time"
)

func VidegoRootHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Got new request ")
	io.WriteString(w, "<html> <head> </head> <body> <h1> Root Handler </h1> </body> </html>")
}

func main() {
	config := config.GetConfigurations()
	host := fmt.Sprintf("%s:%d", config.IpAddress, config.Port)

	server := &http.Server{
		Addr:         host,
		ReadTimeout:  config.ReadTimeout * time.Second,
		WriteTimeout: config.WriteTimeout * time.Second,
	}

	log.Printf("Starting the webserver with configuration %+v", server)
	http.HandleFunc("/", VidegoRootHandler)
	log.Fatal(server.ListenAndServe())
}
