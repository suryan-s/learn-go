// Package main is the entry point for the Go program.
package main

import "fmt" // Importing the fmt package to use its Println function.

func main() {
	// Println is a function from the fmt package used to print a line to standard output.
	// Here, it prints the string "Hello, World!" followed by a newline character.
	fmt.Println("Hello, World!")

	// Another way of using Println, here it prints the string "Hello," followed by the string "World!"
	// Both strings are separated by a space, and the function automatically appends a newline character.
	fmt.Println("Hello,", "World!")

	// Another way is by using Print which prints without followed by a newline character.
	// A newline could be specified by using \n at the end. Print also does not add any
	// space between two strings
	fmt.Print("Hello,", "World!\n")
}
