package task

import "fmt"

func Sum(i int, ch chan int) {
	fmt.Println("adding", i)
	j := <-ch
	ch <- i + j
}
