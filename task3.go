package main

import "fmt"

func main() {
	var number int

	fmt.Println("Введите число")
	fmt.Scan(&number)

	fmt.Println(factorial(number))
}

func factorial(num int) int{
	var counter int = 1

	for i := 1; i <= num; i++ {
		counter *= i
	}
	return counter
}