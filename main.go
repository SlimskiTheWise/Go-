package main

import (
	string2 "awesomeProject/string"
)

var str string = "Hello World!\n"
var koreanStr = "안녕 세상!"

func main() {

	string2.PrintStringLen(str)
	string2.PrintUTFStringLen(koreanStr)

}
