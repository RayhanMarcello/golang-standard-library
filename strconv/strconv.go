package main

import (
	"fmt"
	"strconv"
)

func main() {
	x, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("err : ", err.Error())
	}
	fmt.Println(x)

	id := "12"
	resultInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("err", err.Error())
	}

	fmt.Println(resultInt)

	stringInt := strconv.Itoa(123)
	fmt.Println(stringInt)
}
