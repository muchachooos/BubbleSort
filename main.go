package main

import "fmt"

func main() {
	slice := []int{74, 12, 32, 35, 96, 87, 27, 8, 54, 6}

	bubbleSort(slice)

}

func bubbleSort(slice []int) []int {

	//var j int
	//newSlice := make([]int, len(slice))

	for i := range slice {
		if slice[i] > slice[i+1] {
			swap(slice[i], slice[i+1])
		} else {
			i++
		}
	}
}

func swap(a, b int) {

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("--------")
	//var bKeeper int

	var aKeeper int
	aKeeper = a
	a = b
	b = aKeeper

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
