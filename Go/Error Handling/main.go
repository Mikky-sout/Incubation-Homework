package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(isEmpty(""))
	fmt.Println(divide(5, 0))
}

func isEmpty(s string) error {
	if s == "" {
		return errors.New("string is empty")
	} else {
		fmt.Println(s)
		return nil
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("error : can't divide by 0")
	} else {
		return a / b, nil
	}
}
