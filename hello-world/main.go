// a package clause starts every source file
// main is a special name declaring an executable rather than
package main

// import declaration declares packages used in this file
import "fmt"

// package scoped variables and constants
const secondsInHour = 3600

// a function declaration. main is a special function name
// it is the entry point for the executable program
// Go uses brace brackets to delimit code blocks
func main() {
	// Println() function prints out a line to stdout
	fmt.Println("Hello Go")

	// Local Scoped Variables and Constants Declarations, calling functions etc
	distance := 60.8

	fmt.Printf("the distance in miles is %f \n", distance*0.62137)

}
