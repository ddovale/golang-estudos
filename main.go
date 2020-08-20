package main

import (
	"fmt"

	"github.com/ddovale/golang-estudos/model"
)

func main() {

	p := model.Person{
		Name:    "Daniel",
		Age:     22,
		Married: false,
	}

	fmt.Println(p)

}
