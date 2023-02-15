package ex3

import (
	"fmt"
	"math/rand"
	"sync"
)

func theory2Main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	sum := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)

		// ставим перед любой функцией слово «go», и она выполняется конкурентно в горутине
		go func() {
			// делаем долгий вызов к стороннему API. Так как каждый вызов происходит в своей горутине, мы делаем 10 вызовов одновременно
			n := externalHTTPNum()

			mu.Lock()
			sum += n
			mu.Unlock()

			wg.Done()
		}()
	}

	// ждем, пока все 10 горутин вернут ответ
	wg.Wait()

	fmt.Println(sum) // 55
}

func externalHTTPNum() int {
	return rand.Intn(13)
}
