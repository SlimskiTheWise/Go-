package function

func sum(a int, b int) int {
	return a + b
}

//go 언어는 리턴값에 이름을 지정해줄수있다
// 이름을 지정해줄시에 return 값에 변수를 지정하지 않음
func sum2(a int, b int) (r int) {
	r = a + b
	return
}

//리턴값 여러개 사용하기

func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b

	//ex) sum, dif := SumAndDiff(6,2) // 8, 4

}

//함수의 매개변수 개수가 정해져 있지 않고 유동적으로 변하는 형태를 가변인자라고 한다.
//매개변수의 자료형 앞에 ... 를 붙여 가변인자로 지정하고 ...int 로 지정 했으므로 int 여러개를 받을수 있다

func Sum(i ...int) int {
	total := 0

	for _, value := range i {
		total += value
	}

	return total
}

//익명함수
var r = func(a int, b int) int {
	return a + b
}(1, 2)
