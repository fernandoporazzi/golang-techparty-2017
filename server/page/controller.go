package page

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"page/index.html",
))

// Index renders landing page
func Index(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("http://localhost:8001/movie")

	if err != nil {
		panic(err)
	}

	m, _ := ioutil.ReadAll(res.Body)

	templates.ExecuteTemplate(w, "index.html", template.JS(m))
}
