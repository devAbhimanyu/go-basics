package main

import "fmt"

func main() {
	//using var keyword
	//var identifier type
	var variable int = 1
	fmt.Println(variable)

	//example of type inference
	var name = "tom"
	fmt.Println("the name is", name)

	//go throws compile time error when it finds unused variables
	var age = 10
	// _ works as a blank identifier and can be used to bypass the above error
	_ = age

	//using short declaration operator
	msg := "this is shorthand"
	fmt.Println(msg)

	//below statement will throw an error
	//age := 40
	//as age is already defined an line 16,
	//it threw error => no new variables on left side of :=compiler

	//----multiple declarations using short declaration
	car, cost := "Tata", 5000000
	fmt.Println(car, cost)

	// error at line 27 will not be throw is one of the declared variable is new
	car, cost2 := "Hyundai", 1000000
	fmt.Println(car, cost2)

	//----multiple declarations using var declaration
	// diff types
	var (
		fName string
		cAge  int
	)
	fmt.Println(fName, cAge)

	// same types
	var a, b, c int
	fmt.Println(a, b, c)
	//assignment
	a, b, c = 1, 2, 3
	fmt.Println(a, b, c)
	swap(a, b)

	usingFmt()

	constants()
}

func swap(a int, b int) {
	/* */
	a, b = b, a
	fmt.Println(a, b)
}
