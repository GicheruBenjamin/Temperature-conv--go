package main 

import (
	"fmt"
)


func main() {
	fmt.Println("Please enter the temperature value U want to convert!!");
	var temp float64
	fmt.Scanf("%f", &temp)
	fmt.Println("Choose the unit u want to convert to!!");
	var unit string
	fmt.Scanf("%s", &unit)
	switch unit {
	case "C":
		fmt.Println("The temperature in Celsius is: ", (temp - 32) * 5/9)
	case "F":
		fmt.Println("The temperature in Fahrenheit is: ", (temp * 9/5) + 32)
	}
}