package main

type Subscriber[T any] interface {
	Subscribe(addr string) chan T
}

type mySubscriber[T any] struct{}

func (sub mySubscriber[int]) Subscribe(addr string) chan int {
	ch := make(chan int)

	return ch
}

func main() {
	var subscriber Subscriber[int] = mySubscriber[int]{}

	subscriber.Subscribe("test")
}
