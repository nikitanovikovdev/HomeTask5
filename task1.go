package main

import "fmt"

func main() {
	var number int

	fmt.Println("Введите число")
	fmt.Scan(&number)

	fmt.Println(fib(number))
}

func fib(num int) int{
	var num1 int = 1
	var num2 int = 1
	var result int

	for i := 3; i <= num; i++ {
		result = num1 + num2
		num1 = num2
		num2 = result
	}
	return  result
}
