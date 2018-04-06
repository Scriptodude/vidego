package main

import (
	"fmt"
	"github.com/scriptodude/vidego/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

var cmd *exec.Cmd = nil
var started bool = false
var host string

func VidegoRootHandler(w http.ResponseWriter, req *http.Request) {
	killProcess()

	log.Printf("Got new request %v", req)

	writeIndexOrNotFound(req.RequestURI, w)
}

func VidegoVideoHandler(w http.ResponseWriter, req *http.Request) {
	killProcess()

	if cmd != nil {
		log.Println("We actually are opening the video link.")
		started = true
		HandleVideoRequest(w, req)
	} else {
		log.Println("The user wanted to watch a video.")
		io.WriteString(w, "<h1> Watch the remote ! </h1>")
		cmd = exec.Command("xdg-open", fmt.Sprintf("http://%s/%s", host, req.RequestURI))
		err := cmd.Run()
		if err != nil {
			log.Printf("Error running command : %s", err)
		}
	}
}

func killProcess() {
	if started && cmd != nil {
		log.Println("Killing the process")
		cmd.Process.Kill()
		started = false
		cmd = nil
	}
}

func main() {
	config := config.GetConfigurations()
	host = fmt.Sprintf("%s:%d", config.IpAddress, config.Port)

	server := &http.Server{
		Addr:         host,
		ReadTimeout:  config.ReadTimeout * time.Second,
		WriteTimeout: config.WriteTimeout * time.Second,
	}

	log.Printf("Starting the webserver with configuration %+v", server)
	http.HandleFunc("/watch", VidegoVideoHandler)
	http.HandleFunc("/", VidegoRootHandler)
	log.Fatal(server.ListenAndServe())
}

// Writes the content of the path/index.html to the response writer
// or writes a default Not Found page.
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
