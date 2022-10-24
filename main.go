package main

import "fmt"

func main() {
	slice := []int{74, 12, 32, 35, 6, 54, 27, 8, 87, 96}
	fmt.Println(slice)
	fmt.Println(bubbleSort(slice))
	//bubbleSort(slice)
}

func bubbleSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				fmt.Println(slice)
				slice[j], slice[j+1] = slice[j+1], slice[j]
				fmt.Println(slice)
			}
		}
	}
	return slice
}
