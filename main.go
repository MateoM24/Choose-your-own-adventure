package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/MateoM24/Choose-your-own-adventure/model"
)

type Ops struct {
	Title string
}

func main() {
	plotFile := loadPlot()
	plotMap := convertJSONToMap(plotFile)
	adventure := model.ParseToStories(plotMap)
	adventure.Start()
	// printing example how to traverse stories
	// fmt.Println(adventure.GetStoryNode().Title)
	// fmt.Println(adventure.GetStoryNode().Story)
	fmt.Println(adventure.GetStoryNode().Title)

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

func loadPlot() []byte {
	plotFile := "plot.json"
	fileBytes, err := ioutil.ReadFile(plotFile)
	if err != nil {
		log.Fatalln("Cannot open file with story plot definition", plotFile)
	}
	return fileBytes
}

func convertJSONToMap(file []byte) map[string]map[string]interface{} {
	plotMap := new(map[string]map[string]interface{})
	err := json.Unmarshal(file, plotMap)
	if err != nil {
		log.Fatalln("Cannot unmarshal json file to expected type", err)
	}
	return *plotMap
}

func loadTemplates() *template.Template {
	templates := template.New("templates")
	const basePath = "templates"
	template.Must(templates.ParseGlob(basePath + "/*.html"))
	return templates
}
