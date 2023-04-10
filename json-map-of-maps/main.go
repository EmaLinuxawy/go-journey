package main

import (
	"fmt"

	"github.com/EmaLinuxawy/go-apps/api"
)

func main() {
	output, err := api.Transformation(api.Response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)

	index, err := api.Search(output, "Alabama", "2020")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(index)
}
