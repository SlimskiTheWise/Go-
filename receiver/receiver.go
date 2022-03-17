package receiver

type Rectangle struct {
	Width, Height int
}

func (rect *Rectangle) Area() int {
	return rect.Width * rect.Height // 리시버 변수를 이용해 현재 인스턴스에 접근가능
}
