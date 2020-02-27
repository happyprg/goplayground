package main

import "fmt"

func permute(nums []int) []int {
	uniqueItems := make(map[int]bool, len(nums))

	for _, value := range nums {
		uniqueItems[value] = true
	}

	var result []int
	for key, value := range uniqueItems {
		if value == true {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	items := permute([]int{1, 2, 3, 3})
	fmt.Println(items)
	permuteItems := make(map[int][]int, len(items)*len(items)-1)
	for i := 0; i < len(items); i++ {
		result := make([]int, len(items))
		for j := 0; j < len(items); j++ {
			result[j] = items[j]
		}
		permuteItems[i] = result
	}
	fmt.Println(permuteItems)
}
