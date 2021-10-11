package main

import "fmt"

func iotaVals() {
	//iota successive untyped constants
	//its value is the index of the respective ConstSpec in that constant declaration
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c) //=> 0,1,2

	const (
		c1 = iota
		c2
		c3
	)
	fmt.Println(c1, c2, c3) //=> 0,1,2

	const t = 2
	const (
		a1 = (iota * t)
		a2
		a3
	)
	fmt.Println(a1, a2, a3) //=> 0,1,2

}
