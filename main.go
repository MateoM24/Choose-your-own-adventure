package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/MateoM24/Choose-your-own-adventure/infra"
	"github.com/MateoM24/Choose-your-own-adventure/model"
)

func main() {
	plotMap := infra.LoadPlotFileToMap()
	const startingNodeName string = "intro"
	adventure := model.ParseToStories(plotMap, startingNodeName)
	templates := loadTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		var template *template.Template
		if requestedFile == "" {
			template = templates.Lookup("home.html")
		} else {
			template = templates.Lookup(requestedFile + ".html")
		}
		if template == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			if r.Method == "POST" {
				r.ParseForm()
				nextNodeName := r.PostForm["next"][0]
				adventure.Next(nextNodeName)
			} else {
				adventure.Next(startingNodeName)
			}
			err := template.Execute(w, adventure)
			if err != nil {
				log.Fatalln("Cannot execute template for", requestedFile, ".html")
			}
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}

func loadTemplates() *template.Template {
	templates := template.New("templates")
	const basePath = "templates"
	template.Must(templates.ParseGlob(basePath + "/*.html"))
	return templates
}
