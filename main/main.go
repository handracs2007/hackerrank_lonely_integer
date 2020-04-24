// https://www.hackerrank.com/challenges/lonely-integer/problem

package main

import "fmt"

func lonelyinteger(a []int32) int32 {
	var counts [101]int

	for _, idx := range a {
		counts[idx] ++
	}

	for idx, count := range counts {
		if count == 1 {
			return int32(idx)
		}
	}

	return -1
}

func main() {
	fmt.Println(lonelyinteger([]int32{1}) == 1)
	fmt.Println(lonelyinteger([]int32{1, 1, 2}) == 2)
	fmt.Println(lonelyinteger([]int32{0, 0, 1, 2, 1}) == 2)
}
