package main

import "fmt"

// Return types can have names
func fun() (r int) {
	r = 1
	return
}

// Variadic Functions
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func twoIncrement() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println("twoIncrement:")
	fmt.Println("	", increment())
	fmt.Println("	", increment())
	return
}

func makeEvenGenerator() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}
func evenGenerator() {
	nextEven := makeEvenGenerator()
	fmt.Println("evenGenerator:")
	fmt.Println("	", nextEven())
	fmt.Println("	", nextEven())
	fmt.Println("	", nextEven())
}

//func first() int {
//	fmt.Println("1st")
//	return 1
//}
//func second() {
//	fmt.Println("2nd")
//}
//func main() {
//	defer second()
//	fmt.Println(first())
//}

//func defer_fun() {
//	defer func() {
//		str := recover()
//		fmt.Println(str)
//	}()
//	panic("PANIC")
//}

func zero(xPointer *int) {
	*xPointer = 0
}

func main() {
	fmt.Println("fun:", fun())
	fmt.Println("add:", add(1, 2, 3))
	numbers := []int{1, 2, 3}
	fmt.Println("add:", add(numbers...))
	twoIncrement()
	evenGenerator()

	x := 5
	zero(&x)
	fmt.Println("x:", x)
	xPointer := new(int)
	zero(xPointer)
	fmt.Println("x:", *xPointer)
}
