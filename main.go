package main

import "fmt"

func main() {	
	data := NewStore()

	data.Set("a", "nafas")
	data.Set("b", "blah blah blah")

	data.Delete("b")

	fmt.Println(data.Get("a"))

	fmt.Println(data.Get("b"))
}
