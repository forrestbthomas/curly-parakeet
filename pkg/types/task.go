package types

type TaskWork func(int, chan int)
type TaskDefinition func(ch chan int) chan int
