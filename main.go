package main

import (
	file2 "awesomeProject/if"
	string2 "awesomeProject/string"
)

var str string = "Hello World!\n"
var koreanStr = "안녕 세상!"

func main() {

	string2.PrintStringLen(str)
	string2.PrintUTFStringLen(koreanStr)

	file2.Read2("./hello.txt")
}
