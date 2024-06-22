/*2. Напишите функцию разделения массива чисел на массивы простых и составных чисел.
Для записи в массивы используйте два разных канала и горутины.
Важно, чтобы были использованы владельцы каналов.*/

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{7, 9, 11, 17, 57, 35, 15, 33, 68, 88, 77, 88, 87, 99, 445}
	s := make([]int, 0)
	c := make([]int, 0)
	chSimpleNum := make(chan int)
	chComposeNum := make(chan int)
	d1 := make(chan struct{})
	d2 := make(chan struct{})
	go filterNums(nums, chSimpleNum, chComposeNum)
	go addNum(&s, chSimpleNum, d1)
	go addNum(&c, chComposeNum, d2)

	<-d1
	<-d2

	defer fmt.Println(s)
	defer fmt.Println(c)

}

func addNum(s *[]int, ch <-chan int, done chan struct{}) chan struct{} {
	go func() {
		defer close(done)
		for i := range ch { // чтобы тут цикл остновился канал нужно закрыть у "владельца"
			*s = append(*s, i) // то есть там где пишет в него
		}

	}()
	return done
}

func filterNums(nums []int, chs, chc chan int) {
	defer close(chs)
	defer close(chc)
	for _, val := range nums {

		isSimple := checkNumIsSimple(val)
		if isSimple {
			chs <- val // пишем в канал для простых чисел
			continue
		}
		chc <- val // пишем в канал для составных

	}
}

func checkNumIsSimple(num int) bool {

	for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			return false
		}

	}
	return true

}
