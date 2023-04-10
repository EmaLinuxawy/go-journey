package api

import (
	"encoding/json"
	"errors"
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
					"Population": "4876250",
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
			return output, errors.New("Expected State")
		}
		year, ok := element["Year"].(string)
		if !ok {
			return output, errors.New("Expected Year")
		}
		population, ok := element["Population"].(string)
		if !ok {
			return output, errors.New("Expected Population")
		}
		var stateData []map[string]string
		yearToPopulationMap := make(map[string]string)

		yearToPopulationMap[year] = population
		//yearToPopulationMap["Population"] = population
		stateData = append(output[state], yearToPopulationMap)
		output[state] = stateData
	}

	return output, nil
}

func Search(index map[string][]map[string]string, state, year string) (string, error) {
	for _, val := range index[state] {
		for k, v := range val {
			if k == year {
				return v, nil
			} else {
				return "", fmt.Errorf("Population doesn't exist of the year %s", year)
			}
		}
	}

	return "", fmt.Errorf("state not found: %s", state)

}
