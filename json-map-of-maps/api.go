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
		//fmt.Println(state)
		var stateData []map[string]string

		for k, v := range element {
			v, ok := v.(string)
			if !ok {
				fmt.Printf("Expected value for key %q", k)
			}

			myMap := make(map[string]string)
			myMap[k] = v

			stateData = append(stateData, myMap)
		}
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
