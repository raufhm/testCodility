package main

import "fmt"

func main() {
	result := Solution([]int{2, 2, 2}, 2) //6
	fmt.Println(result)
}

func Solution(A []int, S int) int {

	/*
		- find S inside array A, then count.
		- loop A && sum firstElement with N Element,
			if total / len idx from firstElement to N element is equal to S, then count
	*/

	var counter int
	var sumElement int
	for idx1, dataA := range A {
		sumElement = dataA
		for idx2, dataB := range A {
			if idx2 < idx1 {
				continue
			} else if idx2 == idx1 {
				if dataA == S {
					counter++
				}
			} else {
				lenIdx := len(A[idx1 : idx2+1])
				sumElement = sumElement + dataB
				result := float64(sumElement) / float64(lenIdx)
				if result == float64(S) {
					counter++
				}
			}

		}
	}
	return counter
}
