package ex34

func SumWorker(numsCh chan []int, sumCh chan int) {
	nums := <-numsCh

	s := 0

	for _, v := range nums {
		s += v
	}

	sumCh <- s
}
