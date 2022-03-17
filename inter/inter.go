package inter

import "fmt"

type MyInt int //int 형을 myInt 로 정의

func (i MyInt) Print() { //먼저 (type 새_자료형 자료형) 형식으로 기존 자료형을 새 자료형으로 정의 할수있다. 여기서는 기본 자료형인 int에 메서드를 연결하기 위해 MyInt 를 새로 정의 했다.
	fmt.Println(i)
}

type Rectangle struct { // 사각형 구조체 정의
	width, height int
}

func (r Rectangle) Print() {
	fmt.Println(r.width, r.height)
}

type Printer interface {
	Print()
}
