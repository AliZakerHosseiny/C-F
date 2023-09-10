package main

import "fmt"

func main() {

	var fahrenheit float64

	fmt.Print("Enter the temperature in Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 30) / 2

	fmt.Println("The temperature in Celsius is:", celsius)
}
