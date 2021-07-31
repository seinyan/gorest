package lessontest

type App interface {
	Calc(x int, y int) int
	Plus(x int, y int) int
}

type app struct {
	Name string
}

func (a app) Calc(x int, y int) int {
	return 1
}

func (a app) Plus(x int, y int) int {
	return 2
}

func NewApp() App {
	return &app{
		Name: "TEST",
	}
}