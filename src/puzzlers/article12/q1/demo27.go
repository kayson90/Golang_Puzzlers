package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

// 方案1。
func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

// 方案2。
type calculateFunc func(x int, y int) (int, error)

//其实是一个匿名函数，并且将这个函数作为结果返回
//op是一个自由变量，所以是一个闭包
func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			//新建了一个error
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	// 方案1。
	x, y := 12, 23
	//只需要函数签名是一样的就行
	op := func(x, y int) int {
		return x + y
	}
	result, err := calculate(x, y, op)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
	result, err = calculate(x, y, nil)
	fmt.Printf("The result: %d (error: %v)\n", result, err)

	// 方案2。
	x, y = 56, 78
	add := genCalculator(op)
	result, err = add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
}
