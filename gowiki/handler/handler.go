package handler

import (
	"fmt"
	"net/http"

	"github.com/daniel-joo-dsj/learn-golang/gowiki/page"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := page.LoadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
