package main

import (
	"awesomeProject/function"
	file2 "awesomeProject/if"
	map2 "awesomeProject/map"
	panic2 "awesomeProject/panic"
	po "awesomeProject/pointer"
	re "awesomeProject/receiver"
	r "awesomeProject/reflect"
	string2 "awesomeProject/string"
	s "awesomeProject/struct"
	"awesomeProject/test"
	"fmt"
	"reflect"
)

var str string = "Hello World!\n"
var koreanStr = "안녕 세상!"
var n = 1

func main() {

	panic2.F()

	string2.PrintStringLen(str)
	string2.PrintUTFStringLen(koreanStr)

	file2.Read2("./hello.txt")

	map2.GuessWhatDayTodayIs("Thursday")

	//함수를 변수에 저장해서 쓸수 있음
	i := function.Sum(1, 2, 3)
	fmt.Println(i)

	po.Hello(&n) //1 이 들어있는 변수 n 의 메모리 주소를 hello 함수에 넘김
	fmt.Print(n)

	rect := new(s.Rectangle)
	rect.Width = 20
	rect.Height = 20
	s.RectangleArea(rect)
	fmt.Println(rect)

	rect3 := new(re.Rectangle)
	rect3.Width = 20
	rect3.Height = 20

	fmt.Println(rect3.Area())

	result := test.NewRectangle(12, 12)
	result1 := test.NewRectangle1(12, 12)

	fmt.Println(result, "  ", result1)

	var hello func() //함수를 담을 변수 선언

	fn := reflect.ValueOf(&hello).Elem() //hello의 주소를 넘긴뒤 elem 으로 값 정보를 가져옴

	v := reflect.MakeFunc(fn.Type(), r.H)

	fn.Set(v)

	hello()

}
