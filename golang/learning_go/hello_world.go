package main

import "fmt"

func helloWorld(name string) {
	fmt.Printf("Hello,%s\n", name)
}

func main() {
	helloWorld("Ken Thompson")
}
