package test

type Rectangle struct {
	width, height int
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

func NewRectangle1(width, height int) Rectangle {
	return Rectangle{width, height}
}
