package task

import "fmt"

func ListMultiples(i int, ch chan int) {
	fmt.Println("listing multiples", i)
	for j := i; j > 0; {
		if j%2 == 0 {
			ch <- j
		}
		j = j / 3
	}
}
