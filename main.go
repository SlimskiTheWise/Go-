package main

import (
	"awesomeProject/tcpserver"
	"fmt"
	"net"
)

func main() {

	In, err := net.Listen("tcp", ":8000") //tcp 프로토콜에 8000 포트로 연결을 받음
	if err != nil {
		fmt.Println(err)
		return
	}

	defer In.Close()

	for {
		conn, err := In.Accept() //클라이안트가 연결되면 tcp 연결을 리턴
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close() //main 함수가 끝나기 직전에 tcp 연결을 닫음

		go tcpserver.RequestHandler(conn)
	}
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
