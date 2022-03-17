package pointer

import "fmt"

var numPtr *int //포인터형 변수를 선언하면 nil로 초기화됨 * 를 자료형 앞에 붙힌다

//빈 포인터형 변수는 바로 사용할수 없으므로 new 함수로 메모리를 할당해야 한다

func po() {
	var numPtr1 *int = new(int) //new 함수로 공간할당

	*numPtr1 = 1 //역참조

	fmt.Println(*numPtr1)
}

func Hello(n *int) {
	*n = 2 //포인터 변수 n 을 역참조하여 메모리에 2를 대입
}
