package basic

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() {

}
