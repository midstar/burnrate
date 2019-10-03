package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/GeertJohan/go.rice"
)

var box *rice.Box

func main() {
	var port string
	flag.StringVar(&port, "p", "8080", "Network port to listen to")
	flag.Parse()

	fmt.Println("Starting BurnRate server on port " + port)

	http.Handle("/", http.FileServer(rice.MustFindBox("templates").HTTPBox()))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error! %s\n", err)
	}
}
