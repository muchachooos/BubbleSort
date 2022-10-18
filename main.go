package main

import "fmt"

func main() {
	slice := []int{74, 12, 32, 35, 96, 87, 27, 8, 54}

	bubbleSort(slice)
}

func bubbleSort(slice []int) {

	//swap := func(a, b int) {
	//
	//}
	//
	//newSlice := make([]int, len(slice))

	for i := range slice {
		if slice[i] > slice[1] {
			slice[i], slice[1] = slice[1], slice[i]
		}
	}

	fmt.Println(slice)
}
