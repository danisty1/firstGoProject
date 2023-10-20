package main

import "fmt"

func multiply(x int, y int) int {
	return x * y
}

func sum(x int, y int) int {
	return x + y
}

func diff(x int, y int) int {
	return x - y
}

func divide(x int, y int) int {
	if y == 0 {
		panic("На ноль делить нельзя!")
	} else {
		return x / y
	}
}

func action(x int, y int, operation func(int, int) int) int {
	result := operation(x, y)
	return result
}

func main() {
	x, y := 0, 0
	act := ""
	fmt.Println("Введите первое число: ")
	fmt.Scan(&x)
	fmt.Println("Введите второе число: ")
	fmt.Scan(&y)
	fmt.Println("Выберите операцию (+,-,/,*): ")
	fmt.Scan(&act)
	if string(act) == "+" {
		fmt.Println("Ответ:", action(x, y, sum))
	} else if string(act) == "-" {
		fmt.Println("Ответ:", action(x, y, diff))
	} else if string(act) == "*" {
		fmt.Println("Ответ:", action(x, y, multiply))
	} else if string(act) == "/" {
		fmt.Println("Ответ:", action(x, y, divide))
	} else {
		fmt.Println("Неверная операция.")
	}
}
