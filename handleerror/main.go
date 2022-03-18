package handleerror

import (
	"fmt"
	"log"
	"time"
)

type HelloOneError struct {
	time  time.Time // 시간
	value int       //에러를 발생시킨 값
}

func main() {
	s, err := helloOne(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s) // Hello 1

	s, err = helloOne(2)
	if err != nil {
		log.Fatal(err) // 에러 문자열이 출력되고 프로그램 종료 log.Panic() 을 쓰면 종료하지 않고 런타임 에러 발생
	}

	// 에러가 발생하여 프로그램이 종료되었으므로 아래는 실행되지 않음
	fmt.Println(s)

	fmt.Println("Hello, world")
}
