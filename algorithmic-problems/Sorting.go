package main

import "fmt"

// Bubble sort algorithm
// Time complexity: O(n^2)
// Space complexity: O(1)
func bubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				swap(array, i, j)
			}
		}
	}
}

// Insertion sort algorithm
// Time complexity: O(n^2)
// Space complexity: O(1)
func insertionSort(array []int) {
	var temp int
	var j int

	for i := 1; i < len(array); i++ {
		temp = array[i]

		j = i - 1
		for j >= 0 && temp < array[j] {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = temp
	}
}

// Selection sort algorithm
// Time complexity: O(n^2)
// Space complexity: O(1)
func selectionSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		swap(array, i, minIndex)
	}

}

func main() {
	array := []int{6, 1, 7, 4, 2, 9, 8, 5, 3}
	array = []int{7, 14, 11, 8, 9}

	bubbleSort(array)
	//insertionSort(array2)

	//selectionSort(array)

	fmt.Println(array)
}

// Swap two elements in an array
func swap(array []int, i int, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}
