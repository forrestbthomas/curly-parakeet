package examples

import "fmt"

// Fan In
func Sum(i int, ch chan int) {
	j := <-ch
	fmt.Println("adding", i, j)
	ch <- i + j
}

// Fan Out
func ListMultiples(i int, ch chan int) {
	fmt.Println("listing multiples", i)
	for j := i; j > 0; {
		if j%2 == 0 {
			ch <- j
		}
		j = j / 3
	}
}

// Parallel
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
