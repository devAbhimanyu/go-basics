package main

import "fmt"

// Rules
// in go constant belong to compile time, so value can't be set at runtime
// const a = math.Pow(2,3) is invalid

// cannot use variable to initialize a constant
// t := 10
// const val = t is invalid

//	const lenVariable = len("hello")
//	fmt.Println("this is valid as hello is a string literal", lenVariable)

func constants() {
	const lenVariable = len("hello")
	fmt.Println("this is valid as hello is a string literal", lenVariable)

	//const values can be unused and will not throw error
	const val int = 10
	const a, b int = 1, 2

	fmt.Println(val, a, b)

	const (
		t1 = 1
		t2 = 2
		t3 = 3
	)

	fmt.Println(t1, t2, t3)

	//if only one const is defined all other variable are provided the same value
	const (
		f1 = 1
		f2
		f3
	)

	fmt.Println(f1, f2, f3) //=> 1,1,1

	//const values need to have value defined, otherwise it will throw error
	//const val2 int

	//errors in go are generally thrown during runtime
	// look at the below test case
	// a,b :=1,0

	// fmt.Printf("The output is %v", a/b); //=>panic: runtime error: integer divide by zero

	/*
		//this can be caught during compile time if const is used
		const x =1;
		const y =0;
		fmt.Printf("The output is %v", x/y); //=>invalid operation: division by zero compiler
	*/
	typedConstants()
}

func typedConstants() {
	//this is a typed constant
	const a int = 10

	//this is an untyped constant
	const strVal = "abc"
	const pi = 3.14
	const n = 1
	fmt.Println(a, strVal, pi)
	// with untyped constant type inference takes place
	// Go is a stringly typed language and float * int should not be possible
	const b float64 = 1.1 * 2
	fmt.Printf("type of b is %T \n", b)

	//untyped constants help with normal variable declarations and helps bring flexibility to Go's strongly typed rules
	var i int = n     //n changes to int
	var j float64 = n // on the backend => var j float64 = float64(n)
	_, _ = i, j
}
