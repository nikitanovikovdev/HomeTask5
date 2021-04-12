package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Println("Введите число")
	fmt.Scan(&number)

	if number < 0 {
		fmt.Println("Проверьте Ваше число")
		return
	}

	fmt.Println(check(number))
}

func check(num int) string {
	var counter int
	var result, noResult string = "Верно", "Не верно"

	for i := 0; i <= num; i++ {
		counter += i
	}

	if counter == num*(num+1)/2 {
		return result
	} else {
		return noResult
	}
}
