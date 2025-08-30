package main

import (
	"fmt"
	"time"
)

const LIM = 21
var facts [LIM]uint64

func main() {
	fmt.Println("==================FACTORIAL==================")
	start := time.Now()
	for i:=uint64(0); i < LIM; i++ {
		fmt.Printf("Factorial for %d is : %d \n", i, Factorial(uint64(i)))
	}
	end := time.Now()
	fmt.Printf("Calculation finished in %s \n", end.Sub(start)) //Calculation finished in 2.0002ms

	// fmt.Println("==================FACTORIAL CLOSURE==================")
	// start = time.Now()
	// fact := FactorialClosure()
	// for i:=uint64(0); i < LIM; i++ {
		// fmt.Printf("Factorial closure for %d is : %d \n", uint64(i), fact(uint64(i)))
	// }
	// end = time.Now()
	// fmt.Printf("Calculation finished in %s \n", end.Sub(start)) //Calculation finished in 1ms
// 
	// fmt.Println("==================FACTORIAL MEMOIZATION==================")
	// start = time.Now()
	// var result uint64 = 0
	// for i:=uint64(0); i < LIM; i++ {
		// result = FactorialMemoization(uint64(i))
		// fmt.Printf("The factorial value for %d is %d\n", uint64(i), uint64(result))
	// }
// 
	// end = time.Now()
	// fmt.Printf("Calculation finished in %s\n", end.Sub(start)) // Calculation finished in 0ms
}

func Factorial(n uint64)(result uint64) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func FactorialClosure() func(n uint64)(uint64) {
	var b uint64 = 1
	return func(n uint64)(uint64) {
		if (n > 1) {
			_, b = b, uint64(n) * uint64(b)
		} else {
			return 1
		}

		return b
	}
}

func FactorialMemoization(n uint64)(res uint64) {
	if (facts[n] != 0) {
		res = facts[n]
		return res
	}

	if (n > 0) {
		res = n * FactorialMemoization(n-1)
		return res
	}

	return 1
}
