package main

import (
	"fmt"
	"github.com/scriptodude/vidego/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func VidegoRootHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Got new request %v", req)
	writeIndexOrNotFound(req.RequestURI, w)
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

func writeIndexOrNotFound(path string, response http.ResponseWriter) {
	htmlPath := config.GetHtmlBaseFolder() + path
	content, err := ioutil.ReadFile(htmlPath + "/index.html")

	if err != nil {
		log.Printf("Could not find file %s/index.html, using default 404\n")
		response.WriteHeader(http.StatusNotFound)
		io.WriteString(response, "<html> <head></head><body><h1>404 not found</h1></body> </html>")
	} else {
		response.WriteHeader(http.StatusOK)
		io.WriteString(response, string(content))
	}
}
