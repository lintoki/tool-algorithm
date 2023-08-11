package main

import "fmt"

var arr = []int{9, 4, 5, 7, 8, 8, 8, 9, 9, 97, 5, 3, 2342, 4, 23, 34, 5, 4, 45, 6, 567, 65, 756, 8, 76, 867, 78, 9, 78}

func findKthLargest(nums []int, k int) int {
	nums = quickSort(nums)
	fmt.Println(nums)
	return nums[k-1]

	// deapSort(&nums)
	// return nums[k-1]
}

//冒泡
func mpSort(arr []int) []int {
	for i := 0; i < len(arr); i++ { //第几次排序
		for j := 0; j < len(arr)-1-i; j++ { //排序的数字
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
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
// mergeSort函数接受一个[]int类型的切片作为参数，返回一个有序的切片
func mergeSort(slice []int) []int {
	// 如果切片的长度小于2，说明已经有序，直接返回
	if len(slice) < 2 {
		return slice
	}
	// 找到切片的中间位置，将切片分成左右两半
	mid := len(slice) / 2
	left := slice[:mid]
	right := slice[mid:]
	// 对左右两半分别进行归并排序，然后将结果合并
	return merge(mergeSort(left), mergeSort(right))
}

// merge函数接受两个有序的[]int类型的切片作为参数，返回一个合并后的有序切片
func merge(left, right []int) []int {
	// 创建一个新的切片，用于存放合并后的结果
	result := make([]int, 0, len(left)+len(right))
	// 定义两个指针，分别指向左右两个切片的第一个元素
	i := 0
	j := 0
	// 循环比较左右两个切片的元素，将较小的元素追加到结果切片中，同时移动对应的指针
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// 如果左边的切片还有剩余的元素，将它们追加到结果切片中
	for ; i < len(left); i++ {
		result = append(result, left[i])
	}
	// 如果右边的切片还有剩余的元素，将它们追加到结果切片中
	for ; j < len(right); j++ {
		result = append(result, right[j])
	}
	// 返回结果切片
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

//猴子大王
// 定义一个结构体类型monkey，表示循环链表的节点
type monkey struct {
	m    int     // 猴子的编号
	next *monkey // 指向下一个节点的指针
}

// 定义一个函数newMonkey，接受一个整数参数v，返回一个新创建的monkey节点
func newMonkey(v int) *monkey {
	return &monkey{
		m:    v,
		next: nil,
	}
}

// 定义一个函数monkeyKing，接受一个整数参数num，输出最后的大王编号
func monkeyKing(num int) {
	// 创建一个头节点head，并将其指向自己，形成一个只有一个节点的循环链表
	head := newMonkey(1)
	head.next = head
	// 定义一个当前节点cur，并将其初始化为head
	cur := head
	// 用一个循环创建剩余的9个节点，并将它们依次插入到循环链表中
	for i := 2; i <= 10; i++ {
		// 创建一个新节点item，并将其编号设为i
		item := &monkey{i, head}
		// 将item插入到cur和cur.next之间，并更新cur为item
		cur.next = item
		cur = item
	}
	// 定义一个遍历节点i，并将其初始化为head
	i := head
	// 定义一个计数器n，并将其初始化为1
	n := 1
	// 用一个循环模拟报数和出圈的过程，直到循环链表中只剩下一个节点
	for i.next != i {
		// 如果n等于num-1，说明i.next是要出圈的猴子，则将其从循环链表中删除，并重置n为1
		if n >= num-1 {
			i.next = i.next.next
			n = 1
		} else {
			// 否则，将n加1，表示报数加1
			n++
		}
		// 将i更新为i.next，表示继续向后报数
		i = i.next

	}
	// 输出i.m，即为最后剩下的猴子编号，也就是大王编号
	fmt.Println(i.m)
	return
}

//是否是环链表
//定义一个函数hasCycle，接受一个头节点head，返回一个布尔值表示链表是否有环
func hasCycle(head *ListNode) bool {
	if head == nil { // 如果头节点为空，则返回false
		return false
	}
	slow := head                          // 定义一个慢指针，并初始化为头节点
	fast := head                          // 定义一个快指针，并初始化为头节点
	for fast != nil && fast.Next != nil { // 用一个循环来移动快慢指针，直到快指针到达尾节点或者快慢指针相遇
		slow = slow.Next      // 将慢指针向后移动一个节点
		fast = fast.Next.Next // 将快指针向后移动两个节点
		if slow == fast {     // 如果快慢指针相等，则说明有环，返回true
			return true
		}
	}
	return false // 如果循环结束，说明没有环，返回false
}

//链表反转
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode // 前一个节点
	cur := head       // 当前节点
	for cur != nil {
		temp := cur.Next // 保存下一个节点
		cur.Next = pre   // 将当前节点的Next指针指向前一个节点
		pre = cur        // 将前一个节点更新为当前节点
		cur = temp       // 将当前节点更新为下一个节点
	}
	return pre // 返回新的头节点
}

func main() {
	//findKthLargest(arr, 4)
	monkeyKing(4)
}
