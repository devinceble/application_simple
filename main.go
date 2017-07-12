package main

import "fmt"

type Details struct {
	Name string
	Exp  int
}

func main() {
	var detail Details
	detail.Name = "Leivince John Marte"
	detail.Exp = 6
	fmt.Printf("My name is %q, I have %d years experience as a Senior Developer", detail.Name, detail.Exp)
}
