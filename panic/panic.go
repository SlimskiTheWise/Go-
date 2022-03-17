package panic

import "fmt"

//프로그램이 잘못되어 에러가 발생한뒤 종료되는 상황을 패닉이라함
//문법적인 에러는 아니지만 로직에 따라 에러로 처리하고 싶을때 사용
//recover 함수와 함께 사용하면 다른 언어의 try catch 같이 사용 가능

func F() {

	defer func() {
		s := recover()

		fmt.Println(s)
	}()

	panic("Error")
}
