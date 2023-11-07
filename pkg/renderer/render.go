package renderer

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var tc = make(map[string]*template.Template)

func RendererTemplate(w http.ResponseWriter, tmpl string) {
	//create a template cache
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Println(err)
	}

	//get requested template from cache
	t, ok := templateCache[tmpl]
	if !ok {

		log.Println(ok)
	}

	//optional final error check
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the files named *.page.html(or tmpl) from ./template
	pages, err := filepath.Glob("./template/*.page.html")
	if err != nil {
		return myCache, err
	}

	//range through all files ending with *.page.html (or tmpl)
	for _, page := range pages {
		//here page - is a full path to the file, and we need only name of the file
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./template/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./template/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateSet

	}
	return myCache, nil

}
