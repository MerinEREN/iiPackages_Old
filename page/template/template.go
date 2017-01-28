package template

import (
	"github.com/MerinEREN/iiPackages/page/content"
	// usr "github.com/MerinEREN/iiPackages/user"
	"html/template"
	"log"
	"net/http"
)

var (
	html        = template.Must(template.ParseFiles("../iiClient/src/index.html"))
	RenderIndex = renderTemplate("index")
)

func renderTemplate(title string) func(w http.ResponseWriter, p *content.Page) {
	return func(w http.ResponseWriter, p *content.Page) {
		if title == "index" {
			err := html.Execute(w, p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Fatalln(err)
			}
		}

	}
}
