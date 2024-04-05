package main

import "fmt"

func main() {

	fmt.Println("##############POINTERS##############")

	var ptr *int

	var j int = 98

	fmt.Printf("ptr: %d\n", ptr)
	fmt.Printf("j: %d\n", j)

	ptr = &j

	fmt.Printf("ptr: %d\n", *ptr)

}
