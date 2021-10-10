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
}

func swap(a int, b int) {
	/* */
	a, b = b, a
	fmt.Println(a, b)
}

func usingFmt() {
	// VERBS:
	// %d -> decimal
	// %f -> float
	// %s -> string
	// %q -> double-quoted string
	// %v -> value (any)
	// %#v -> a Go-syntax representation of the value
	// %T -> value Type
	// %t -> bool (true or false)
	// %p -> pointer (address in base 16, with leading 0x)
	// %c -> char (rune) represented by the corresponding Unicode code point

	var name, age = "tom", 20
	fmt.Println(name, "is", age, "years old.") // => tom is 20 years old.
	//Println adds /n after each print statement

	//** fmt.Printf() **//

	// fmt.Printf() prints out to stdout according to a format specifier called verb.
	// It doesn't add a newline (\n)

	a, b, c := 10, 15.5, "test"
	grades := []int{10, 20, 30}

	fmt.Printf("a is %d, b is %f, c is %s \n", a, b, c) // => a is 10, b is 15.500000, c is test
	fmt.Printf("%q\n", c)                               // => "test"
	fmt.Printf("%v\n", grades)                          // => [10 20 30]
	fmt.Printf("%#v\n", grades)                         // => b is of type float64 and grades is of type []int
	fmt.Printf("b is of type %T and grades is of type %T\n", b, grades)
	// => b is of type float64 and grades is of type []int
	fmt.Printf("The address of a: %p\n", &a) // => The address of a: 0xc000016128
	fmt.Printf("%c and %c\n", 100, 51011)    // =>  d and 읃

	const pi float64 = 3.14159265359
    fmt.Printf("pi is %.4f\n", pi) // => formatting with 4 decimal points
 
    // %b -> base 2
    // %x -> base 16
    fmt.Printf("255 in base 2 is %b\n", 255)  //  => 255 in base 2 is 11111111
    fmt.Printf("101 in base 16 is %x\n", 101) // => 101 in base 16 is 65
 
    // fmt.Sprintf() returns a string. Uses the same verbs as fmt.Printf()
    s := fmt.Sprintf("a is %d, b is %f, c is %s \n", a, b, c)
    fmt.Println(s) // => a is 10, b is 15.500000, c is test
}
