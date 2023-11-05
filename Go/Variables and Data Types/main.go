package main

import (
	"fmt"
	"strconv"
)

func main() {
	adding()
	age()
	great()
	typeCheck()
	isTrue()
	showName()
	convert()
	toString()
	changeValue()
}

// Declare Int and sum each number
func adding() {
	x := 2
	y := 13

	fmt.Println(x + y)
}

// Declare with var
func age() {
	var year int = 2542

	var age int = 2542 - year
	fmt.Println("Your age is", age)
}

// Declare string
func great() {
	s1 := "Hello"
	s2 := "World"
	fmt.Println(s1, s2)
}

// Check type of variables
func typeCheck() {
	number := 1234.5
	fmt.Printf("number is type %T and value is %v\n", number, number)

	text := "Hi, I'm Mik"
	fmt.Printf("text is type %T and value is %v\n", text, text)
}

// Declare boolean
func isTrue() {
	//condition := true
	var condition bool = false
	if condition {
		fmt.Println("Yes, It's true")
	} else {
		fmt.Println("No, It's false")
	}
}

// Declare package level variables
var name string = "Paritat"

func showName() {
	fmt.Println("My name is ", name)
}

// Type conversion with strconv
func convert() {
	number := 101
	fmt.Println("Go basic " + strconv.Itoa(number))

	text := "50"
	n, _ := strconv.Atoi(text)

	fmt.Println("50 + 50 =", n+n)
}

// Convert to string with Sprint
func toString() {
	var height int = 170
	fmt.Println("My height is " + fmt.Sprint(height))
}

// Declare a constants variable
func changeValue() {
	const pi = 3.14
	const area float32 = 4.50
	fmt.Println(area, pi)
}
