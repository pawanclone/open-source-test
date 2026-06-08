package main

import "fmt"

func main() {
	fmt.Println("hello world")
	sum(1,2,4,5,6)
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

