package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("não pode trabalhar com 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "não pode trabalhar com ele."}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 falhou:", e)
		} else {
			fmt.Println("f1 funcionou:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 falhou:", e)
		} else {
			fmt.Println("f2 funcionou:", r)
		}
	}
	_, e := f2(42)
	ae, ok := e.(*argError)

	if ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
