package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	// var utc = time.Date(2005, time.January, time.Now().Day(), 1, 2, 3, 4, time.UTC)
	// fmt.Println("utc : ", utc)
	// fmt.Println("utc (local): ", utc.Local())

	formater := "2005-01-06 15:04:03"
	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formater, value)
	if err != nil {
		fmt.Println("err : ", err.Error())
	}
	fmt.Println(valueTime)

	var duration1 time.Duration = time.Second * 100
	fmt.Println(duration1)

	var duration2 time.Duration = time.Millisecond * 10
	fmt.Println(duration2)

	var duration3 time.Duration = duration1 - duration2
	fmt.Println(duration3)
}
