package practice

import (
	"fmt"
	"testing"
)

func TestCh(t *testing.T) {
	t.Run("ch에 값 넣고빼기", func(t *testing.T) {
		t.Run("case1", func(t *testing.T) {
			ch := make(chan int, 3)
			ch <- 1
			ch <- 2
			ch <- 3
			ch <- 4
			fmt.Println("finish")
		})

		t.Run("case2", func(t *testing.T) {
			ch := make(chan int, 3)
			close(ch)
			ch <- 1
			fmt.Println("finish")
		})

		t.Run("case3", func(t *testing.T) {
			ch := make(chan int, 3)
			close(ch)
			fmt.Println(<-ch)
			fmt.Println(<-ch)
			fmt.Println(<-ch)
			fmt.Println("finish")
		})

		t.Run("case4", func(t *testing.T) {
			ch := make(chan int, 3)
			ch <- 1
			ch <- 2
			<-ch
			ch <- 3
			ch <- 4
			fmt.Println("finish")
		})
	})

	t.Run("range", func(t *testing.T) {
		// TODO 1단계: ch에서 값을 가져와 출력하기
		// TODO 2단계: 에러 없애기
		ch := make(chan int, 3)
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)

		for value := range ch {
			fmt.Println(value)
		}
	})

	t.Run("for-select", func(t *testing.T) {
		// TODO 1단계: ch에서 값을 가져와 출력하기
		// TODO 2단계: 에러 없애기
		ch := make(chan int, 3)
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)

		for {
			select {
			case value, ok := <-ch:
				// 채널이 클로즈 되더라도 zero value를 반환하기 때문에
				// ok를 확인해야 한다.
				if !ok {
					return
				}
				fmt.Println(value)
			}
		}
	})
}