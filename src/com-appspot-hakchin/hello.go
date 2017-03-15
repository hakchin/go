package hello

import (
	"fmt"
	"html/template"
	"net/http"
)

//func init() {
//	http.HandleFunc("/", handler)
//}
//
//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello, world!")
//}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}
