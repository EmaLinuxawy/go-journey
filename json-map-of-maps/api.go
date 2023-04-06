package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

var Response = []byte(`{
	"data": [
			{
					"ID State": "04000US01",
					"State": "Alabama",
					"ID Year": 2020,
					"Year": "2020",
					"Population": "4893186",
					"Slug State": "alabama"
			},
			{
					"ID State": "04000US01",
					"State": "Alabama",
					"ID Year": 2019,
					"Year": "2019",
					"Population": "4893186",
					"Slug State": "alabama"
			},
			{
				"ID State": "04000US02",
				"State": "Alaska",
				"ID Year": 2020,
				"Year": "2020",
				"Population": 736990,
				"Slug State": "alaska"
		
			}
	]
}`)

func Transformation(res []byte) (map[string]any, error) {
	result := make(map[string]any)

	var data map[string]any
	err := json.Unmarshal(res, &data)
	if err != nil {
		return result, err
	}

	population := 0
	for _, item := range data {
		for _, el := range item.([]any) {
			element := el.(map[string]any)
			if element["State"] == "Alabama" {
				p, _ := strconv.Atoi(element["Population"].(string))
				population += p
			}
		}
	}
	return map[string]any{"Total Population": population}, nil
}

func main() {
	result, err := Transformation(Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
