package main

import (
	"fmt"
	"math"
)

func main() {

	//FOR loops
	for i := 0; i < 10; i++ {
		println(i)
	}

	j := 0
	for j < 10 {
		print(j)
		j++
	}

	// for {
	// 	println("loop infinito")
	// }

	//IFs
	if math.Pow(2, 2) < 4 {
		println("menor que 4")
	} else {
		println("maior/igual que 4")
	}

	if v := math.Pow(2, 2); v < 4 {
		fmt.Printf("v (%f) menor que 4\n", v)
	} else {
		fmt.Printf("v (%f) maior/igual que 4\n", v)
	}

	//SWITCHs
	num := 0
	switch num {
	case 0:
		println("zero")
	default:
		println("não zero")
	}

	switch {
	case num == 0:
		println("zero")
	default:
		println("não zero")
	}

}
