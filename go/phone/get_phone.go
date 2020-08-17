package main

import (
	"fmt"
	"regexp"
)

func main() {
	length := 0
    fmt.Println("Enter the number of inputs")
    fmt.Scanln(&length)
    fmt.Println("Enter the inputs")
    numbers := make([]string, length)
    for i := 0; i < length; i++ {
        fmt.Scanln(&numbers[i])
    }
	fmt.Println(getPhone(numbers))
}

func getPhone(numbers []string) (list []string) {
	for i := 0; i < len(numbers); i++ {
		regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
		reg := regexp.MustCompile(regular)
		if reg.MatchString(numbers[i]) {
			phone := numbers[i]
			list = append(list, phone)
		}
	}
	return
}
