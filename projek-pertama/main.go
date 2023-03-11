package main

import "fmt"

func main() {
	var i = 21
	var j = true
	var russia = "Я (ya)"

	fmt.Println(i)
	fmt.Printf("%T \n", i)
	fmt.Println("%")
	fmt.Println(j)
	fmt.Printf("%v \n", j)
	fmt.Printf("%x \n", []byte(russia))
}
