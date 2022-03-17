package embedding

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello")
}

type Student struct { //학생 구조체 안에 사람 구조체를 필드로 가지고 있음 has-a 관계
	p      Person
	school string
	grade  int
}

//main 메소드에서 var s Studnet 이 있을때 s.p.greeting() 으로 사용가능

//type Student struct { //필드명 없이 타입만 선언할 경우 is-a 관계
//	Person
//	school string
//	grade  int
//}
