package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"kontolodon1", 123},
		{"kontolodon2", 44},
		{"kontolodon3", 11},
		{"kontolodon4", 33},
		{"kontolodon5", 22},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
