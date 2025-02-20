package main

import (
	"fmt"
	// "net/http"

	"github.com/daniel-joo-dsj/learn-golang/gowiki/page"
)

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/view/"):]
// 	p, _ := page.LoadPage(title)
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

func main() {
	p1 := &page.Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.Save()
	p2, _ := page.LoadPage("TestPage")
	fmt.Println(string(p2.Body))
}
