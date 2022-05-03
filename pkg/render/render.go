package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var funtions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	cache, err := CreateTemplateCache()

	if err != nil {
		fmt.Print("An error ocuured", err)
	}

	t, ok := cache[tmpl]

	if !ok {
		log.Fatal(err)
	}

	buff := new(bytes.Buffer)
	_ = t.Execute(buff, nil)

	_, err = buff.WriteTo(w)

	if err != nil {
		fmt.Println("Could not write file to the browswer")
	}

	fmt.Println(cache)

	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	// err := parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("Error parsing template:", err)
	// 	return
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, nil
	}

	for _, page := range pages {
		// fmt.Println(page)
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(funtions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCache, nil
		}

		if len(matches) > 0 {
			_, err := ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return myCache, nil
			}
		}

		fmt.Println(name)
		myCache[name] = ts
		// ts, err := template.New(name)
	}

	return myCache, nil
}
