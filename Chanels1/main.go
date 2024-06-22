/*1. Напишите 2 функции:
	Первая функция читает ввод с консоли. Ввод одного значения заканчивается нажатием клавиши enter.
	Вторая функция пишет эти данные в файл. Свяжите эти функции каналом.
Работа приложения должна завершится при нажатии клавиш ctrl+c с кодом 0. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handle(c chan os.Signal) {
	for {
		<-c // This line will block until a signal is received

		os.Exit(0)
	}
}

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go handle(c)

	fmt.Println("Пиши уже:")
	readConsChan := readConsole()

	writeToFile(readConsChan)

	defer fmt.Println("Exit")

}

func readConsole() <-chan string {
	sc := bufio.NewScanner(os.Stdin)
	ch := make(chan string)
	go func() {
		defer close(ch)
		for sc.Scan() {
			txt := sc.Text()
			ch <- txt
		}

	}()

	return ch
}

func writeToFile(ch <-chan string) {

	for txt := range ch {
		writeText(txt)
	}

}

func writeText(text string) {
	file, err := os.OpenFile("savedTextFromKeyboard.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Unable to create  or append to file:", err)
		os.Exit(1)
	}

	file.WriteString(fmt.Sprintf("%s\n", text))

	fmt.Println("Writen.")
	defer file.Close()
}
