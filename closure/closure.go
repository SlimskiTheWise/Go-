package closure

//클로저란 함수안에 함수를 선언 밎 정의 할수있고 바깥족 함수에 선언된 변수에도 접근할 수있는 함수를 말함

func calc() func(x int) int {
	a, b := 3, 5
	return func(x int) int {
		return a*x + b //클로저이므로 함수를 호출할때마다 변수 a 와 b 의 값을 사용할 수 있음
	}
}
