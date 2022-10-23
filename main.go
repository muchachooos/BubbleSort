package main

import "fmt"

func main() {
	slice := []int{74, 12, 32, 35, 6, 54, 27, 8, 87, 96}
	bubbleSort(slice)
	fmt.Println(bubbleSort2(slice))
	sort(slice)
}

func bubbleSort(slice []int) {
	//newSlice := make([]int, len(slice))

	fmt.Println(slice)
	for i := range slice {
		if i == len(slice)-1 {
			break
		}
		if slice[i] > slice[i+1] {
			slice[i], slice[i+1] = slice[i+1], slice[i]
			fmt.Println(slice)
			fmt.Println("1-------else-------1")
		} else {
			i++
		}
		fmt.Println(slice)
		fmt.Println("2---__++__-------2")
	}
}

func sort(slice []int) {

	fmt.Println(slice)
	for i := range slice {
		if i == len(slice)-1 {
			fmt.Println("break!!!!")
			break
		}
		for slice[i] > slice[i+1] {
			slice[i], slice[i+1] = slice[i+1], slice[i]
			fmt.Println(slice)
			fmt.Println("1-------else-------1")
			if slice[i] <=
				slice[i+1] {
				i++
			}
		}
		fmt.Println(slice)
		fmt.Println("2---__++__-------2")
	}
}


func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				fmt.Println(array)
			}
		}
	}
	return array
}

func bubbleSort2(slice []int) []int {
	n := len(slice)
	swapped := true
	//перебрать все элементы в нашем списке
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			// если текущий элемент больше следующего,поменять их местами
			if slice[i] > slice[i+1] {
				// если текущий элемент больше следующего,поменять их местами
				fmt.Println(slice)
				slice[i], slice[i+1] = slice[i+1], slice[i]
				// поменять местами значения, используя присваивание кортежа в Go
				swapped = true
				// устанавливаем swapped в true
				// если цикл закончился, а swapped по-прежнему равен
				// false, наш алгоритм будет считать, что список
				// полностью отсортирован.
			}
		}
	}
	return slice
}
