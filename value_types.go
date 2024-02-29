package main

import "fmt"

func main() {
	// Strings: Go supports string concatenation using the `+` operator.
	// Here, "Go" and "lang" are concatenated to form the string "Golang".
	fmt.Println("Go" + "lang")

	// Integers: Go supports arithmetic operations like addition.
	// Here, 1 and 1 are added together to give the result 2.
	fmt.Println("2+1 =", 2+1)
	fmt.Println("3-1 =", 3-1)
	fmt.Println("5*6 =", 5*6)
	fmt.Println("6/2 =", 6/2)
	fmt.Println("3%2 =", 3%2)

	// Explicitly specifying float32. By default all floats would be in float64
	var num1 float32 = 1.1
	var num2 float32 = 1.5
	fmt.Println("1.1 + 1.5 (float32) = ", num1+num2)

	// Booleans: Go supports logical operations like AND.
	// Here, true AND true evaluates to true.
	fmt.Println(true && true)

	// Booleans: Go supports logical operations like OR.
	// Here, true OR false evaluates to true.
	fmt.Println(true || false)

	// Booleans: Go supports logical negation.
	// Here, negation of true results in false.
	fmt.Println(!true)
}
