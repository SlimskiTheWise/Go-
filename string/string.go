package string

import (
	"fmt"
	"unicode/utf8"
)

//문자의 길이를 구할때는 len 을 쓴다
func PrintStringLen(s string) {
	fmt.Println(len(s))
}

//한글 한자 일본어등의 문자열 길이를 구하려면 utf8.RuneCountInString 함수를 쓴다.
func PrintUTFStringLen(s string) {
	fmt.Println(utf8.RuneCountInString(s))
}
