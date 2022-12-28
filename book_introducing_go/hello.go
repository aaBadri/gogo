package main

import "fmt"

func convertFahr2Cel() float64 {
	fmt.Print("enter fahrenheit:")
	var input float64
	fmt.Scanf("%f", &input)
	output := (input - 32) * 5 / 9
	return output
}

func loopExample() float64 {
	//var arr [3]float64
	//arr[0] = 1
	//arr[1] = 3
	//arr[2] = 5
	arr := [5]float64{1, 3, 5}
	total := 0.0
	for _, val := range arr {
		total += val
	}
	return total / float64(len(arr))
}

func sliceExample() {
	var slice1 []float64
	slice2 := []float64{1, 2, 3}
	slice2 = append(slice1, 7, 9)
	fmt.Println(slice1, slice2)

	src := []int{1, 2, 3}
	dst := make([]int, 2)
	var dst2 []int
	copy(dst, src)
	copy(dst2, src)
	fmt.Println(src, dst, dst2)
}
func main() {
	//fmt.Println(convertFahr2Cel())
	//fmt.Println("loopExample:", loopExample())
	sliceExample()
}
