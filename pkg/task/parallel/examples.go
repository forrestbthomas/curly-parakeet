package task

import "fmt"

func Filter(i int, ch chan int) {
	fmt.Println("filtering", i)
	if i%2 != 0 {
		ch <- i
	}
}

func Doubler(i int, ch chan int) {
	fmt.Println("doubling", i)
	ch <- i * 2
}

func Tripler(i int, ch chan int) {
	fmt.Println("tripling", i)
	ch <- i * 3
}

func DoubleLen(i int, ch chan int) {
	fmt.Println("doubling length", i)
	ch <- i
	ch <- i
}
