package main

import (
	"fmt"
)

func main() {
	// nums1 := []int{1,2,2,1}
	// nums2 := []int{2,2}
	nums1 := []int{4,9,5}
	 nums2 := []int{9,4,9,8,4}

fmt.Println(intersection(nums1, nums2))
}

func intersection(nums1 []int, nums2 []int) []int {
    myMap := make(map[int]struct{})
    var result []int

    for _, num := range nums1 {
        myMap[num] = struct{}{}
    }

    for _, num := range nums2 {
        if _, exists := myMap[num]; exists {
            result = append(result, num)
        delete(myMap, num)
        }
    }
    return result    
}