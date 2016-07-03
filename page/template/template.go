package template

import (
	"github.com/MerinEREN/iiPackages/page/content"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// FIND A WAY TO ASSIGN funcMap's AFTER PARSING A TEMPLATE !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
var (
	// templates          = template.Must(template.ParseGlob("static/templates/*.html"))
	templates map[string]*template.Template
	RenderIndex        = renderTemplate("index")
	RenderAccount      = renderTemplate("account")
	RenderRoles        = renderTemplate("roles")
	RenderUserSettings = renderTemplate("userSettings")
)

func renderTemplate(title string) func(w http.ResponseWriter, p *content.Page) {
	return func(w http.ResponseWriter, p *content.Page) {
		var err error
		// Parsing all templates only once =)
		if templates == nil {
			templates, err = parseTemplates(w, title)
			// log.Println(templates)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Fatalln(err)
			}
		}
		// templates = templates.Funcs(initFuncMap(title))
		// err := templates.ExecuteTemplate(w, title+".html", p)
		err = templates[title+".html"].Execute(w, p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err)
		}
	}
}

func parseTemplates(w http.ResponseWriter, s string) (map[string]*template.Template, 
error) {
	layout, err := template.New("layout.html").Funcs(initFuncMap(s)).
		ParseFiles("static/templates/layout.html")
	result := make(map[string]*template.Template)
	if err != nil {
		return nil, err
	}
	dir := new(os.File)
	dir, err = os.Open("static/templates/blocks")
	defer dir.Close()
	var fis []os.FileInfo
	fis, err = dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		f := new(os.File)
		f, err = os.Open("static/templates/blocks/" + fi.Name())
		if err != nil {
			return nil, err
		}
		var content []byte
		content, err = ioutil.ReadAll(f)
		f.Close()
		if err != nil {
			return nil, err
		}
		var tmpl *template.Template
		tmpl, err = layout.Clone()
		if err != nil {
			return nil, err
		}
		tmpl.Funcs(initFuncMap(s)).Parse(string(content))
		result[fi.Name()] = tmpl
	}
	return result, nil
}

func initFuncMap(s string) template.FuncMap {
	funcMap := template.FuncMap{
		"IsAdmin": IsAdmin,
	}
	return funcMap
}

func IsAdmin(s string) bool {
	return s == "admin"
}

/* package template

import (
	"github.com/MerinEREN/iiPackages/page/content"
	"html/template"
	"log"
	"net/http"
)

var (
	// templates          = template.Must(template.ParseGlob("static/templates/*.html"))
	RenderIndex        = renderTemplate("index")
	RenderAccount      = renderTemplate("account")
	RenderRoles        = renderTemplate("roles")
	RenderUserSettings = renderTemplate("userSettings")
)

func renderTemplate(title string) func(w http.ResponseWriter, p *content.Page) {
	return func(w http.ResponseWriter, p *content.Page) {
		// templates = templates.Funcs(initFuncMap(title))
		// err := templates.ExecuteTemplate(w, title+".html", p)
		// FIND A WAY TO USE .Funcs WITH templates VARIABLE !!!!!!!!!!!!!!!!!!!!!!!
		file := "static/templates/" + title + ".html"
		tmpl, _ := template.New(title + ".html").Funcs(initFuncMap(title)).
		ParseFiles(file)
		err := tmpl.Execute(w, p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln(err)
		}
	}
}

func initFuncMap(s string) template.FuncMap {
	var funcMap template.FuncMap
	switch s {
	case "account":
		funcMap = template.FuncMap{
			"IsAdmin": IsAdmin,
		}
	}
	return funcMap
}

func IsAdmin(s string) bool {
	return s == "admin"
} */
