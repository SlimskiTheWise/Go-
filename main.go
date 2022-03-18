package main

import (
	h "awesomeProject/http"
	"net/http"
)

func main() {

	http.ListenAndServe(":9090", h.NewHttpHandler())
}

//map2.GuessWhatDayTodayIs("Thursday")
//
////함수를 변수에 저장해서 쓸수 있음
//i := function.Sum(1, 2, 3)
//fmt.Println(i)
//
//po.Hello(&n) //1 이 들어있는 변수 n 의 메모리 주소를 hello 함수에 넘김
//fmt.Print(n)
//
//var hello func() //함수를 담을 변수 선언
//
//fn := reflect.ValueOf(&hello).Elem() //hello의 주소를 넘긴뒤 elem 으로 값 정보를 가져옴
//
//v := reflect.MakeFunc(fn.Type(), r.H)
//
//fn.Set(v)
//
//hello()
