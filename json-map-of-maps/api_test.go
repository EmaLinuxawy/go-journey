package main

import (
	"fmt"
	"testing"
)

func TestApi(t *testing.T) {
	result, err := Transformation(Response)
	if err != nil {
		t.Log(err)
		t.Error("faild to transform")
	}
	fmt.Println(result)

	_, ok := result["Total Population"]
	if !ok {
		fmt.Println(ok)
		t.Error("faild to get state")
	}
}
