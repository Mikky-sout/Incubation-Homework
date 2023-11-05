package main

import "fmt"

func main() {
	isEven()
	loop()
	forEach()
	whileLoop()
	pyramid()
}

// IF else
func isEven() {
	n := 12

	if n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

// For loop
func loop() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

// For each
func forEach() {
	array := [5]int{1, 3, 5, 7, 9}

	for _, val := range array {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

// While loop
func whileLoop() {
	count := 0
	for count < 10 {
		fmt.Print(count, " ")
		count++
	}
}

// Pyramid
func pyramid() {
	n := 5

	for i := 0; i <= n; i++ {

		for x := 0; x < n-i; x++ {
			fmt.Print(" ")
		}

		for x := 0; x < (i*2)-1; x++ {
			fmt.Print("X")
		}
		fmt.Println()
	}

}
