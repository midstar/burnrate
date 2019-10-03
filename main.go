package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/GeertJohan/go.rice"
)

var box *rice.Box

func main() {
	fmt.Println("Starting BurnRate server")
	box = rice.MustFindBox("templates")
	http.HandleFunc("/", serverHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error! %s\n", err)
	}
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Path
	if len(r.URL.Path) > 0 {
		fileName = r.URL.Path[1:] // Remove '/'
	}
	if fileName == "" {
		// Default is index page
		fileName = "index.html"
	}
	bytes, err := box.Bytes(fileName)
	if err != nil || len(bytes) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Unable to find: %s!", fileName)
	} else {
		if filepath.Ext(fileName) == ".html" {
			w.Header().Set("Content-Type", "text/html")
		} else {
			w.Header().Set("Content-Type", "image/png")
		}
		w.Write(bytes)
	}
}
