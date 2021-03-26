package main

import "fmt"

func InsertionSort(arr []int) []int {
	for out := 1; out < len(arr); out++ {
		temp := arr[out]
		in := out
		for ; in > 0 && arr[in-1] >= temp; in-- {
			arr[in] = arr[in-1]
		}
		arr[in] = temp
	}
	return arr
}

func main() {
	unsorted := []int{12, 5, 8, 3, 2, 5, 8, 3, 2, 5, 8, 3, 2, 5, 8, 3, 2, 5, 8, 3, 2, 5, 8, 3, 2, 5, 8, 3, 2}
	fmt.Println("Welcome to the InsertionsSorting!")
	fmt.Println(InsertionSort(unsorted))
	//int aree := 2
	//fmt.Println(aree)
	//fmt.Println(InsertionSort())
}
