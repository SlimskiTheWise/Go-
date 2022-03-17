package structure

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

type Rectangle1 struct { //자료형이 같을 경우 한줄로 나열 가능
	width, height int
}

//Rectangle 타입으로 인스턴스 생성
var rect Rectangle

func decl() {
	var rect1 *Rectangle   //구조체 포인터 선언
	rect1 = new(Rectangle) //구조체 포인터에 메모리 할당

	rect2 := new(Rectangle) // 구조체 포인터 선언과 동시에 메모리 할당

	fmt.Println(rect1, rect2)
}

//구조체 생성자 패턴 활용하기
func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

//rect := NewRectangle(20, 10) // &{20 10}

//사각형의 넓이 구하기
func RectangleArea(rect *Rectangle) int {
	return rect.Width * rect.Height
}
