package main

import (
	"fmt"
)

func main() {

	var a, b float64

	fmt.Println("enter two numbers to perform operations a, b:")

	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)

	fmt.Println("numbers given by user:", a, b, "are processing now")
	defer last(a, b)
	first(a, b)
	second(a, b)
	third(a, b)
}

func last(a float64, b float64) {
	var result float64
	result = a + b
	fmt.Println("addition of two given numbers:", result)
}

func first(a float64, b float64) {
	var result1 float64
	result1 = a - b
	fmt.Println("substraction of two given numbers:", result1)
}

func second(a float64, b float64) {
	var result2 float64
	result2 = a * b
	fmt.Println("multiplication of two given numbers:", result2)
}

func third(a float64, b float64) {
	var result3 float64
	result3 = (a / b)
	fmt.Println("division of two given numbers:", result3)
}
