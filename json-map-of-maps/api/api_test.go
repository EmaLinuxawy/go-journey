package api

import (
	"fmt"
	"testing"
)

func TestTransformation(t *testing.T) {
	result, err := Transformation(Response)
	if err != nil {
		t.Log(err)
		t.Error("faild to transform")
	}
	fmt.Println(result)

	_, ok := result["Alabama"]
	if !ok {
		fmt.Println(ok)
		t.Error("faild to get state")
	}
}

func TestSearch(t *testing.T) {
	index, err := Transformation(Response)
	if err != nil {
		t.Log(err)
		t.Error("faild to index")
	}
	s, err := Search(index, "Alabama", "2020")
	if err != nil {
		t.Error(err)
	}
	if s != "4893186" {
		t.Error("Invalid Population")
	}
	_, err = Search(index, "Cancun", "2020")
	if err == nil {
		t.Error(err)
	}
	_, err = Search(index, "Alabama", "2021")
	if err == nil {
		t.Error(err)
	}

}
