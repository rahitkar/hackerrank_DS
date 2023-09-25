package main

import "fmt"

func reverseArray(nums []int32) []int32 {
	var reverseNum []int32
	for _, num := range nums {
		reverseNum = append([]int32{num}, reverseNum...)
	}
	return reverseNum
}

func main() {
	nums := []int32{1, 4, 3, 2}
	fmt.Println(reverseArray(nums))
}
