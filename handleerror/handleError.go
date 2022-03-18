package handleerror

import (
	"fmt"
	"time"
)

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}

	return "", HelloOneError{time.Now(), n}
}

func (e HelloOneError) Error() string { //에러 타입이 아니더라도 에러 함수를 구현하면 에러로 사용가능
	return fmt.Sprintf("%v: %d는 1 이 아닙니다", e.time, e.value)
}

