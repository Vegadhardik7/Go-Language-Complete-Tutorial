package main

import "fmt"

func main() {

	defer printMessage("SKY")
	defer printMessage("EARTH")
	defer printMessage("AIR")

	values := accessElement([]int{2, 4, 5, 6}, 2)
	fmt.Println(values)
}

func accessElement(nums []int, ele int) int {
	defer printMessage("WATER")
	val := nums[ele]
	return val
}

func printMessage(m string) {
	fmt.Println(m)
}
