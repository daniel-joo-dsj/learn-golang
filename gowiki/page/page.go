package page

import (
	"log"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	logger := log.New(os.Stdout, "INFO", log.Ldate|log.Ltime)
	if err != nil {
		logger.Println("Could not load page", err)
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
