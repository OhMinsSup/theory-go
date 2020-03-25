package libs

import (
	"log"
	"runtime"
	"sync"
)

var (
	counter int
	wg sync.WaitGroup
)

func increment(num int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		// 스레드를 양보하여 큐로 돌아간다.
		runtime.Gosched()

		value++
		// 원래 변수에 증가된 값 전달
		counter = value
	}
}

func RaceListing() {
	wg.Add(2)

	go increment(1)
	go increment(2)

	wg.Wait()
	log.Println("최종 결과", counter)
}

func Listing() {
	// 스케줄러가 사용 할 논리 프로세서를 의미
	runtime.GOMAXPROCS(1)

	// wg 프로그램의 종료를 대기하기 위해 사용한다
	// 각각의 고루틴마다 하나씩, 총 두개의 카운터 생성
	// 카운팅 세마포어를 이용해서 실행 중인 고루틴의 기록을 관리
	var wg1 sync.WaitGroup
	// waitGroup의 값이 0보다 큰 값이면 wait 메소드의 실행이 블락
	wg1.Add(2)

	log.Println("고루틴을 실행합니다.")

	go func() {
		// 함수가 종료를 알라기 위해서 사용
		defer wg1.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				log.Printf("%c", char)
			}
		}
	}()

	go func() {
		// 함수가 종료를 알라기 위해서 사용
		defer wg1.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				log.Printf("%c", char)
			}
		}
	}()

	log.Println("고루틴 대기중....")
	wg1.Wait()
	log.Println("프로그램 종료")
}
