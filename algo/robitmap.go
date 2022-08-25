package main

import (
	"bytes"
	"fmt"
	"github.com/RoaringBitmap/roaring"
)

func main() {
	// example inspired by https://github.com/fzandona/goroar
	fmt.Println("==roaring==")
	rb1 := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000)
	fmt.Println(rb1.String())

	rb2 := roaring.BitmapOf(3, 4, 7, 1000)
	fmt.Println(rb2.String())

	rb3 := roaring.New()
	fmt.Println(rb1.ToArray())

	fmt.Println("Cardinality: ", rb1.GetCardinality())

	fmt.Println("Contains 3? ", rb1.Contains(3))

	rb1.Or(rb2)
	fmt.Println(rb1.String())

	rb3.Add(1)
	rb3.Add(5)

	rb3.Or(rb1)

	// computes union of the three bitmaps in parallel using 4 workers
	roaring.ParOr(4, rb1, rb2, rb3)
	// computes intersection of the three bitmaps in parallel using 4 workers
	roaring.ParAnd(4, rb1, rb2, rb3)

	// prints 1, 3, 4, 5, 1000
	i := rb3.Iterator()
	for i.HasNext() {
		fmt.Println(i.Next())
	}
	fmt.Println()

	// next we include an example of serialization
	buf := new(bytes.Buffer)
	fmt.Println(2222, string(buf.Bytes()))
	rb1.WriteTo(buf) // we omit error handling
	fmt.Println(2222, string(buf.Bytes()))
	newrb := roaring.New()
	newrb.ReadFrom(buf)
	if rb1.Equals(newrb) {
		fmt.Println("I wrote the content to a byte stream and read it back.")
	}
	// you can iterate over bitmaps using ReverseIterator(), Iterator, ManyIterator()
}

//示例：
////rb1 := roaring.BitmapOf(1, 2, 3, 4, 5, 100, 1000, 1000000)
//
//for i := uint32(1); i < 300000000; i++ {
//rb1.Add(i)
//}
//
//buf := new(bytes.Buffer)
//rb1.WriteTo(buf) // we omit error handling
//p.data.Redis.Set(ctx, "book_test_111", buf.String())
//fmt.Println(buf.String(), "=====")
//
//rb2 := roaring.NewBitmap()
//s, _ := p.data.Redis.GetString(ctx, "book_test_111")
////fmt.Println(s)
//buf2 := bytes.NewBufferString(s)
//rb2.ReadFrom(buf2)
//
//fmt.Println(rb2.Contains(100), rb2.Contains(100000000), rb2.Contains(100000000), rb2.Contains(299999999))
//return
