package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := []string{"Hello, ", "world!"}
	fmt.Println("s1 : " + strings.Join(s1, " "))

	s2 := strings.Split("Hello, world", " ")
	fmt.Println("s2 : " + s2[1])

	data := "asset get file"
	s3 := strings.Split(data, " ")
	fmt.Println("s3 : " + s3[0])
}
