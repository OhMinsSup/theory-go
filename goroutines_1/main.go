package main

import "fmt"

// Go는 매우 간단하지만 가볍고 빠른 동시성 개념인 고루틴(goroutine)을 지원한다
// 스레드에 비해 메모리 소비, 설치와 철거 비용, context switching 비용이 현저히 적다
// 일반적으로 스레드는 guard page라는 메모리 영역과 함께 1Mb 정도로 시작하는데, 고루틴은 2kb의 스택 공간만 필요하다
// 추후에 힙 저장 공간을 확보하여 사용하는데, 이 특징 덕분에 고루틴 수천 개를 만들어 사용해도 부담이 적다
// 고루틴은 정보를 공유하는 방식이 아니라 서로 메세지를 주고 받는 방식으로 동작한다.
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
