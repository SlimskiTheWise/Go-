package _if

import (
	"fmt"
	"io/ioutil"
)

//이코드를 if 조건문 안에서 함수를 실행한 뒤 조건을 판단하는 방식으로 바꿀수 있다.

func Read(file string) {
	var b []byte
	var err error

	b, err = ioutil.ReadFile(file)

	if err == nil {
		fmt.Printf("%s", b)
	}
}

//if 조건문 안에서 함수를 실행하고 ; 으로 구분한 뒤 조건식을 작성한다.

func Read2(file string) {
	if b, err := ioutil.ReadFile(file); err == nil {
		fmt.Printf("%s", b)
	}
}