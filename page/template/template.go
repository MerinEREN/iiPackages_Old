package template

import (
	"github.com/MerinEREN/iiPackages/page/content"
	// usr "github.com/MerinEREN/iiPackages/user"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	index = template.Must(template.
		ParseFiles("static/templates/index.html"))
	layout = template.Must(template.ParseFiles("static/templates/layout.html"))
	// layout = template.Must(template.New("layout.html").Funcs(initFuncMap()).
	// ParseFiles("static/templates/layout.html"))
	layoutClone     = template.Must(layout.Clone())
	layoutAndBlocks = template.Must(layoutClone.
			ParseGlob("static/templates/blocks/*.html"))
	templates             map[string]*template.Template
	RenderIndex           = renderTemplate("index")
	RenderAccount         = renderTemplate("account")
	RenderRoles           = renderTemplate("roles")
	RenderUserSettings    = renderTemplate("userSettings")
	RenderAccountSettings = renderTemplate("accountSettings")
)

func renderTemplate(title string) func(w http.ResponseWriter, p *content.Page) {
	return func(w http.ResponseWriter, p *content.Page) {
		var err error
		if title == "index" {
			err = index.Execute(w, p)
		} else {
			// Parsing all templates only once =)
			if len(templates) == 0 {
				templates, err = parseTemplates()
				// log.Println(templates)
				if err != nil {
					http.Error(w, err.Error(),
						http.StatusInternalServerError)
					log.Fatalln(err)
				}
			}
			err = templates[title+".html"].Execute(w, p)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err)
		}
	}
}

func parseTemplates() (map[string]*template.Template, error) {
	result := make(map[string]*template.Template)
	dir, err := os.Open("static/templates/pages")
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	var fis []os.FileInfo
	fis, err = dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	var tmpl *template.Template
	f := new(os.File)
	var content []byte
	// log.Println(index.DefinedTemplates())
	// log.Println(layout.DefinedTemplates())
	// log.Println(layoutAndBlocks.DefinedTemplates())
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		f, err = os.Open("static/templates/pages/" + fi.Name())
		if err != nil {
			return nil, err
		}
		content, err = ioutil.ReadAll(f)
		f.Close()
		if err != nil {
			return nil, err
		}
		tmpl, err = layoutAndBlocks.Clone()
		if err != nil {
			return nil, err
		}
		tmpl.Parse(string(content))
		result[fi.Name()] = tmpl
		// log.Println(i, fi.Name(), tmpl.DefinedTemplates())
	}
	return result, nil
}

/* func initFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		"isAdmin":         usr.User.IsAdmin,
		"isContentEditor": usr.User.IsContentEditor,
	}
	return funcMap
} */
