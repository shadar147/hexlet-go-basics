package ex34

import "fmt"

func Test() {
	numsCh := make(chan []int)
	sumCh := make(chan int)

	go SumWorker(numsCh, sumCh)
	numsCh <- nil
	res := <-sumCh
	fmt.Println(res)

	go SumWorker(numsCh, sumCh)
	numsCh <- []int{}
	res = <-sumCh
	fmt.Println(res)

	go SumWorker(numsCh, sumCh)
	numsCh <- []int{10, 10, 10}
	res = <-sumCh
	fmt.Println(res)

	go SumWorker(numsCh, sumCh)
	numsCh <- []int{500, 5, 10, 25}
	res = <-sumCh
	fmt.Println(res)

	close(numsCh)
	close(sumCh)
}
