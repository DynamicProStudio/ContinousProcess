package main

import "fmt"

var version = "dev"

func main() {
	fmt.Printf("Version: %s\n", version)
	fmt.Println(hello("Ada"))
}

func hello(msg string) string {
	return fmt.Sprintf("Hello, %s", msg)
}
