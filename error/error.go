package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	err := div(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
}

func div(a, b int) error {
	if b == 0 {
		return errors.New("Divident cant be zero :) ")
	}
	return nil
}

type LoginError struct {
}
