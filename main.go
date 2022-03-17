package main

import (
	"awesomeProject/function"
	file2 "awesomeProject/if"
	map2 "awesomeProject/map"
	string2 "awesomeProject/string"
	"fmt"
)

var str string = "Hello World!\n"
var koreanStr = "안녕 세상!"

func main() {

	string2.PrintStringLen(str)
	string2.PrintUTFStringLen(koreanStr)

	file2.Read2("./hello.txt")

	map2.GuessWhatDayTodayIs("Thursday")

	//함수를 변수에 저장해서 쓸수 있음
	i := function.Sum(1, 2, 3)
	fmt.Println(i)

}
