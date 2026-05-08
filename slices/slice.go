package main

import (
	"fmt"
	"slices"
)

func main() {
	name := []string{"rayhan", "marcello", "ananda", "prunomo"}
	angka := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(slices.Min(angka))
	fmt.Println(slices.Contains(name, "rayhan"))
	fmt.Println(slices.Index(name, "marcello"))

}
