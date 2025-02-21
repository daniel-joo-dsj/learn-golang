package handler

import (
	"net/http"

	"github.com/daniel-joo-dsj/learn-golang/gowiki/page"
	"github.com/daniel-joo-dsj/learn-golang/gowiki/templates"
	"github.com/daniel-joo-dsj/learn-golang/gowiki/validator"
)

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validator.ValidPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := page.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	templates.RenderTemplate(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := page.LoadPage(title)
	if err != nil {
		p = &page.Page{Title: title}
	}
	templates.RenderTemplate(w, "edit", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &page.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
