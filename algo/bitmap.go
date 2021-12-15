package main

import (
	"reflect"
	"runtime/pprof"
	//"sort"
	"time"
)

type Bitmap struct {
	value []uint64
	lenth uint
}

func (bit *Bitmap) Add(num int) bool {
	index := num >> 6 //区
	site := num & 63  //位

	sliceLen := len(bit.value)
	for sliceLen < index+1 {
		bit.value = append(bit.value, uint64(0))
		sliceLen++
	}

	bit.value[index] |= uint64(1 << uint(site))
	bit.lenth++
	return true
}

func (bit *Bitmap) Del(num int) bool {
	index := num >> 6
	site := num & 63

	//判断是否存在
	if !(bit.isExit(num)) {
		return false
	}

	bit.value[index] ^= uint64(1 << uint(site))
	bit.lenth--
	return true
}

//顺序输出数组的数字
func (bit *Bitmap) listData() []uint64 {
	data := make([]uint64, bit.lenth)
	if len(bit.value) == 0 {
		return data
	}

	index := 0                                  //循环角标
	bits := reflect.TypeOf(bit.value[0]).Bits() //循环位数
	for i, v := range bit.value {
		if v == 0 {
			continue
		}

		for site := 0; site < bits; site++ {
			if (v & (1 << uint(site))) != 0 {
				data[index] = uint64(i*bits + site)
				index++
			}
		}

	}

	return data
}

//是否存在
func (bit *Bitmap) isExit(num int) bool {
	index := num >> 6
	site := num % 64

	if (bit.value[index] & (1 << uint(site))) == 0 {
		return false
	}

	return true
}

func normalNum() {
	//num := []int{}
	//i := int(0)
	//for i < 100000000 {
	//	num = append(num, i)
	//	i++
	//}
	//
	//sort.Ints(num)
}

func main() {
	//常规数组操作
	//normalNum()

	//bitmap数组操作
	//bitmap := new(Bitmap)
	//
	//i := int(0)
	//for i < 100000000 {
	//	bitmap.Add(int(i))
	//	i++
	//}

	//bitmap := new(Bitmap)
	//
	//bitmap.Add(4)
	//bitmap.Add(63)
	//bitmap.Add(64)
	//bitmap.Add(506)
	//bitmap.Add(120)
	//fmt.Println(bitmap)

	//fmt.Println(bitmap.listData())
	//fmt.Println(bitmap.isExit(4))

	//bitmap.Del(4)
	//fmt.Println(bitmap)
	//
	//fmt.Println(bitmap.isExit(5))
	//
	//fmt.Println(bitmap.listData())

	time.Sleep(100 * time.Second)

	pprof.StopCPUProfile()
}
