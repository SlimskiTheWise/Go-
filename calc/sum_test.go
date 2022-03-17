package calc

import "testing"

type TestData struct { //테스트 데이터 구조체
	argument1 int
	argument2 int
	result    int
}

var testdata = []TestData{ //테스트 데이터 슬라이스
	{2, 6, 8},
	{-8, 3, -5},
	{6, -6, 0},
	{0, 0, 0},
}

//func TestSum(t *testing.T) {
//	r := Sum(1, 2)
//	if r != 3 {
//		t.Error("결과값이 3이 아닙니다 r=", r)
//	}
//}

func TestSum(t *testing.T) {
	for _, d := range testdata {
		r := Sum(d.argument1, d.argument2)
		if r != d.result {
			t.Errorf(
				"%d + %d의 결괏값이 %d 이 아닙니다. r=%d",
				d.argument1,
				d.argument2,
				d.result,
				r,
			)
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}
