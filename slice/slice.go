package slice

import "fmt"

func AddString(a []string, s string) {
	a = append(a, s)
	fmt.Println(a)
}

func AddSlice(a []string, b []string) {
	a = append(a, b...)
	fmt.Println(a)
}

//슬라이스는 레퍼런스 타입이다. 내장된 배열에 대한 포인터이므로 슬라이스끼리 대입하면 값이 복사되지 않고 참조만 된다.
//슬라이스를 복사하기 위해서는 copy 함수를 쓴다.
//슬라이스를 복사함으로 원본 슬라이스는 바뀌지 않는다.

func CopySlice(a []string) {
	b := make([]string, len(a)) // 복사될 슬라이스의 공간을 할당.
	copy(b, a)
	fmt.Println(b)
}

//슬라이스는 기존 슬라이스의 일정 위치를 지정하여 부분 슬라이스를 만들수 있다.

func SliceSlice(a []int) {
	b := a[0:len(a)]
	fmt.Println(b)
}
