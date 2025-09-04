package main

import "fmt"

func Sort(arr []int, l, r int) {
	if l > r {
		return
	}

	pivot := arr[(l+r)/2]
	i := l
	j := r

	for i <= j {
		for arr[i] < pivot {
			i++
		}

		for arr[j] > pivot {
			j--
		}

		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	Sort(arr, l, j)
	Sort(arr, i, r)
}

func quickSort(arr []int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)

	Sort(sortedArr, 0, (len(sortedArr) - 1))

	return sortedArr
}

func main() {
	arr := []int{2, 5, 6, 1, 7, 3, 9, 13, 11, 15, 10}
	result := quickSort(arr)

	fmt.Println("Not Sorted:", arr)
	fmt.Println("Sorted:    ", result)
}
