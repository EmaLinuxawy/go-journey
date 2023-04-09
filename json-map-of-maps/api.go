package main

import (
	"encoding/json"
	"fmt"
)

var Response = []byte(`{
	"data": [
			{
					"ID State": "04000US01",
					"State": "Alabama",
					"ID Year": "2020",
					"Year": "2020",
					"Population": "4893186",
					"Slug State": "alabama"
			},
			{
					"ID State": "04000US01",
					"State": "Alabama",
					"ID Year": "2019",
					"Year": "2019",
					"Population": "4893186",
					"Slug State": "alabama"
			},
			{
				"ID State": "04000US02",
				"State": "Alaska",
				"ID Year": "2020",
				"Year": "2020",
				"Population": "736990",
				"Slug State": "alaska"
		
			}
	]
}`)

func Transformation(res []byte) (map[string][]map[string]string, error) {
	output := make(map[string][]map[string]string)
	var data map[string][]map[string]any

	err := json.Unmarshal(res, &data)
	if err != nil {
		return output, err
	}

	for _, element := range data["data"] {
		state, ok := element["State"].(string)
		if !ok {
			fmt.Println("Expected State as string type")
		}
		year, ok := element["Year"].(string)
		if !ok {
			fmt.Println("Expected Year as string type")
		}
		population, ok := element["Population"].(string)
		if !ok {
			fmt.Println("Expected Population as string type")
		}

		//fmt.Println("population of Year", year, "is:", population)
		var stateData []map[string]string

		yearToPopulationMap := make(map[string]string)
		yearToPopulationMap["Year"] = year
		yearToPopulationMap["Population"] = population
		stateData = append(stateData, yearToPopulationMap)

		output[state] = stateData
	}
	return output, nil
}

func main() {
	output, err := Transformation(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
