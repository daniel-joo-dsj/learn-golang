//go:build ignore

package main

import (
	"log"
	"net/http"

	"github.com/daniel-joo-dsj/learn-golang/gowiki/handler"
)

func main() {
	http.HandleFunc("/view/", handler.MakeHandler(handler.ViewHandler))
	http.HandleFunc("/edit/", handler.MakeHandler(handler.EditHandler))
	http.HandleFunc("/save/", handler.MakeHandler(handler.SaveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
