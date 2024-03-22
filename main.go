package main

import (
	"net/http"
	"main.go/helpers"
)

var port = ":1357"

func main() {
	http.HandleFunc("/", helpers.StatusUpdateFromRandom)
	http.ListenAndServe(port, nil)
}
