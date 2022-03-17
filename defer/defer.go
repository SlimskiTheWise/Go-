package _defer

import (
	"fmt"
	"os"
)

func ReadHello() {
	file, err := os.Open("/hello.text")
	defer file.Close() //지연 호줄한 file.Close() 가 맨 마지막에 호출됨

	if err != nil {
		fmt.Println(err)
		return // defer 호출
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return // defer 호출

		fmt.Println(string(buf))
		// defer 호출
	}
}
