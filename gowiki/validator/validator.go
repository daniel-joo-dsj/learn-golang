package validator

import (
	"errors"
	"net/http"
	"regexp"
)

var ValidPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func GetTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := ValidPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}
