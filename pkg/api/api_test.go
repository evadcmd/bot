package api

import (
	"encoding/json"
	"fmt"
	"testing"
)

type obj struct {
	A *int
	B *int `json:"b,omitempty"`
}

func TestJson(t *testing.T) {
	x := obj{}
	dtoX, err := json.Marshal(x)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", dtoX)

	dtoXRef, err := json.Marshal(&x)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", dtoXRef)

	five := 5
	y := obj{A: &five}
	dtoY, err := json.Marshal(y)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+s\n", dtoY)

	dtoYRef, err := json.Marshal(&y)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+s\n", dtoYRef)

	z := obj{A: &five, B: &five}
	dtoZ, err := json.Marshal(z)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", dtoZ)
}
