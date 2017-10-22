package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(i int) string {
	result := ""
	if i%15 == 0 {
		result = "fizzbuzz"
	} else if i%5 == 0 {
		result = "buzz"
	} else if i%3 == 0 {
		result = "fizz"
	} else {
		result = strconv.Itoa(i)
	}
	return result
}

func main() {
	fmt.Println(fizzbuzz(15))
}
