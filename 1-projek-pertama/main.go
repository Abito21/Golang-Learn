package main

import "fmt"

func main() {
	// Variabel
	var i = 21
	var j = true

	fmt.Println(i)
	fmt.Printf("%T \n", i)
	fmt.Println("%")
	fmt.Println(j)

	// Variabel unicode
	j = false
	// ja (Я)
	var russia = "\u042f"

	fmt.Printf("%v \n", j)
	// fmt.Printf("%q \n", russia)
	fmt.Printf(russia)

	// Variabel base
	var base10 = 21
	var base8 = 25
	var base16f = 'f'
	var base16F = 'F'

	fmt.Printf("\n%d \n", base10)
	fmt.Printf("%o \n", base8)
	fmt.Printf("%x \n", base16f)
	fmt.Printf("%X \n", base16F)

	// Variabel menampilkan nilai unicode
	var urussia = 'Я'

	fmt.Printf("%+q \n", urussia)

	// Variabel float
	var k float64 = 123.456

	fmt.Printf("%.3f \n", k)
	fmt.Printf("%f \n", k)
	fmt.Printf("%e \n", k)
}
