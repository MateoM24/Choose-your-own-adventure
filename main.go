package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("App is starting...")
	plotFile := loadPlot()
	plotMap := convertJSONToMap(plotFile)
	fmt.Println(plotMap)
}

func loadPlot() []byte {
	plotFile := "plot.json"
	fileBytes, err := ioutil.ReadFile(plotFile)
	if err != nil {
		log.Fatalln("Cannot open file with story plot definition", plotFile)
	}
	return fileBytes
}

func convertJSONToMap(file []byte) map[string]interface{} {
	plotMap := new(map[string]interface{})
	err := json.Unmarshal(file, plotMap)
	if err != nil {
		log.Fatalln("Cannot unmarshal json file to expected type", err)
	}
	return *plotMap
}

// check how to use it instead of too generic interface{}
type plot struct {
	title   string
	story   []string
	options []option
}

type option struct {
	text string
	arc  string
}

// func loadTemplates() *template.Template {
// 	templates := template.New("templates")
// 	const basePath = "templates"
// 	template.Must(templates.ParseGlob(basePath + "/*.html"))
// 	return templates
// }
