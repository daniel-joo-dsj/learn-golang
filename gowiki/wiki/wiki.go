//go:build ignore

package main

import (
	"log"
	"net/http"

	"github.com/daniel-joo-dsj/learn-golang/gowiki/handler"
)

func main() {
	http.HandleFunc("/view/", handler.ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
