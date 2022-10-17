package practice

import (
	"fmt"
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T) {
	t.Run("goroutine으로 값 출력하기", func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			i := i
			go func() {
				defer wg.Done()
				fmt.Println(i)
			}()
		}
		wg.Wait()
	})

	t.Run("goroutine 끝날때까지 기다리기", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				fmt.Println(i)
			}
		}()
		wg.Wait()
	})
}
