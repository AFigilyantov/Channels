/*3. Реализуйте функцию слияния двух каналов в один.*/
// https://medium.com/justforfunc/two-ways-of-merging-n-channels-in-go-43c0b57cd1de
package main

import "sync"

func main() {

}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)  // Создаем канал
	var wg sync.WaitGroup  // создаем waitgroupe
	wg.Add(len(cs))        // добавляем количество каналов на входе
	for _, c := range cs { // слушаем каналы
		go func(c <-chan int) {
			for v := range c {
				out <- v // записываем в выходной канал
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
