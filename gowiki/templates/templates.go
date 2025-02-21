package templates

import (
	"html/template"
	"net/http"

	"github.com/daniel-joo-dsj/learn-golang/gowiki/page"
)

var templatesPath = "../templates/"
var TEMPLATES = template.Must(template.ParseFiles(templatesPath+"edit.html", templatesPath+"view.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {
	err := TEMPLATES.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
