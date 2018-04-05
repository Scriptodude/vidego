package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type VideoType int

const (
	youtube VideoType = iota
	other             = iota
)

func HandleVideoRequest(w http.ResponseWriter, req *http.Request) {
	log.Printf("New Video Request %v", req)

	formValue := req.FormValue("v")

	videoType := getVideoType(formValue)
	switch videoType {
	case youtube:
		log.Println("Going for youtube frame")
		formValue = formatYoutubeUri(formValue)
		writeInIframe(formValue, w)
	case other:
		log.Printf("Still to come...")
	default:
		log.Printf("Unknown type : %d", videoType)
	}
}

func getVideoType(value string) VideoType {
	isYoutube, _ := regexp.MatchString("youtube.com", value)

	if isYoutube {
		return youtube
	} else {
		return other
	}
}

func formatYoutubeUri(uri string) string {
	return strings.Replace(uri, "watch?v=", "embed/", -1)
}

func writeInIframe(uri string, w http.ResponseWriter) {
	// Todo, make the frame window-sized
	log.Printf("Writing %s using iframe", uri)
	iframe := fmt.Sprintf("<iframe src='%s?autoplay=1' style='position:fixed; top:0px; left:0px; bottom:0px; right:0px; width:100%%; height:100%%; border:none; margin:0; padding:0; overflow:hidden; z-index:999999;'></iframe>", uri)
	io.WriteString(w, fmt.Sprintf("<html><head></head><body>%s</body></html>", iframe))
}
