package main

import "fmt"

func main() {
	length := 0
    fmt.Println("Enter the number of inputs")
    fmt.Scanln(&length)
    fmt.Println("Enter the inputs")
    numbers := make([]int, length)
    for i := 0; i < length; i++ {
        fmt.Scanln(&numbers[i])
    }
	fmt.Println(numbers)
	fmt.Println(getOddNumber(numbers)) 
}	

func getOddNumber(numbers []int) int {
	var i, result int
	result = 0
	for i = 0; i < len(numbers); i++ {
		result ^= numbers[i]
	}
	return result
}