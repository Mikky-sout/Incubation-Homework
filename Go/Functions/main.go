package main

import "fmt"

func main() {
	fmt.Println(add(2, 4))
	fmt.Println(showDetail("Mik", 24))
	fmt.Println(sumAndMinus(5, 3))
}

func add(a, b int) int {
	return a + b
}

func showDetail(name string, age int) string {
	return fmt.Sprintf("My name is %v. I'm %v years old", name, age)
}

func sumAndMinus(a, b int) (int, int) {
	return a + b, a - b
}
