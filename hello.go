package main

import "fmt"

func convertFahr2Cel() float64 {
	fmt.Print("enter fahrenheit:")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input - 32) * 5 / 9
	return output
}
func main() {
	fmt.Println("Hello World", 10)
	fmt.Println(convertFahr2Cel())
}
