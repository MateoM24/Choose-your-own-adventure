package infra

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func LoadPlotFileToMap() map[string]map[string]interface{} {
	plotFile := loadPlot()
	return convertJSONToMap(plotFile)
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
