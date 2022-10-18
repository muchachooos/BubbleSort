package main

import "fmt"

func main() {
	slice := []int{74, 12, 32, 35, 96, 87, 27, 8, 54, 6}
	fmt.Println(bubbleSort(slice))
}

func bubbleSort(slice []int) []int {
	//var j int
	//newSlice := make([]int, len(slice))

	//var iKeeper int

	for i := range slice {
		if i == len(slice)-1 {
			break
		}
		if slice[i] > slice[i+1] {
			slice[i], slice[i+1] = slice[i+1], slice[i]
		} else {
			i++
		}
	}
	return slice
}

//if slice[i] > slice[i+1] {
//slice[i] = iKeeper
//slice[i] = slice[i+1]
//slice[i+1] = iKeeper
//}
//var aKeeper int
//aKeeper = a
//a = b
//b = aKeeper
//--------------------------------
//for i := range slice {
//	if slice[i] > slice[i+1] {
//		swap(slice[i], slice[i+1])
//		i++
//	} else {
//		i++
//	}
//	if i == len(slice)-1 {
//		break
//	}
//}
//return slice
