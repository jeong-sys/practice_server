package main

import (
	"fmt"
	"module/greeting"
)

func main() {
	message := greeting.Hello("John")
	fmt.Println(message)
}
