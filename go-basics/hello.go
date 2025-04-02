package main

import (
	"errors"
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func addFuc(x, y int) int {
	return x + y
}

func fact(x int) int {
	if x == 0 {
		return 1
	}
	return fact(x-1) * x
}

func factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Factorial if undefined for negative number")
	}

	if n == 0 {
		return 1, nil
	}

	fact, err := factorial(n - 1)
	if err != nil {
		return 0, err
	}

	return n * fact, nil
}

func divide(dividend, divisor int) (quotient int, remainder int) {
	if divisor == 0 {
		return
	}

	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func main() {
	fmt.Println("hello")
	fmt.Println(add(1, 2))
	fmt.Println(fact(5))
	fmt.Println(addFuc(1, 2))

	q, r := divide(10, 3)
	fmt.Println("quotient", q)
	fmt.Println("Remainder", r)

	fmt.Println(sum(10, 203, 40))
}
