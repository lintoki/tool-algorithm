package main

var arr = []int{9, 4, 5, 7, 8, 8, 8, 9, 9, 97, 5, 3, 2342, 4, 23, 34, 5, 4, 45, 6, 567, 65, 756, 8, 76, 867, 78, 9, 78}

func findKthLargest(nums []int, k int) int {
	nums = qSort(nums)
	return nums[k-1]

	// deapSort(&nums)
	// return nums[k-1]
}

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

//归并排序
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

//堆排序
func deapSort(arr *[]int) {
	lenth := len(*arr)

	//构建小顶堆，从最后一个非子叶节点开始
	for i := lenth / 2; i >= 0; i-- {
		sort(arr, i, lenth)
	}

	//小顶堆首尾互换，将换过的重新排列
	for j := lenth - 1; j >= 0; j-- {
		(*arr)[0], (*arr)[j] = (*arr)[j], (*arr)[0]
		sort(arr, 0, j)
	}
}

func sort(arr *[]int, cur int, lenth int) {
	//从子节点开始，判断当前节点和子节点大小，
	//如果当前节点大于左右节点，换位置，进行下个节点的判断
	for child := cur*2 + 1; child < lenth; child = child*2 + 1 {
		if child < lenth-1 && (*arr)[child] > (*arr)[child+1] {
			child++
		}

		if (*arr)[cur] < (*arr)[child] {
			break
		}

		(*arr)[cur], (*arr)[child] = (*arr)[child], (*arr)[cur]
		cur = child
	}

}

//快速排序
func qSort(arr []int) []int {
	if len(arr) <= 0 {
		return arr
	}

	n := arr[0]
	var l, r []int

	for k, v := range arr {
		if k == 0 {
			continue
		}
		if v < n {
			r = append(r, v)
		} else {
			l = append(l, v)
		}
	}

	l = qSort(l)
	r = qSort(r)

	return append(append(l, n), r...)
}

func main() {
	findKthLargest(arr, 4)
}
