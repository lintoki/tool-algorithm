package main

import "fmt"

var arr = []int{9, 4, 5, 7, 8, 8, 8, 9, 9, 97, 5, 3, 2342, 4, 23, 34, 5, 4, 45, 6, 567, 65, 756, 8, 76, 867, 78, 9, 78}

//快排
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	item := arr[0]

	var left, right []int

	for k, v := range arr {
		if k == 0 {
			continue
		}

		if v < item {
			left = append(left, v)
			continue
		}
		right = append(right, v)
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, item), right...)
}

//归并
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	i := len(arr) / 2

	left := mergeSort(arr[0:i])
	right := mergeSort(arr[i:])

	result := msort(left, right)

	return result
}

func msort(left []int, right []int) []int {
	i, j := 0, 0
	ll, lr := len(left), len(right)

	result := []int{}

	for i < ll && j < lr {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	fmt.Println(mergeSort(arr))
	fmt.Println(quickSort(arr))
}
