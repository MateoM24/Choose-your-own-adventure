package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("App is starting...")
	loadPlot()
}

func loadPlot() {
	plotFile := "plot.json"
	fileBytes, err := ioutil.ReadFile(plotFile)
	if err != nil {
		log.Fatalln("Cannot open file with story plot definition", plotFile)
	}
	fmt.Println(string(fileBytes))
}

// func loadTemplates() *template.Template {
// 	templates := template.New("templates")
// 	const basePath = "templates"
// 	template.Must(templates.ParseGlob(basePath + "/*.html"))
// 	return templates
// }
