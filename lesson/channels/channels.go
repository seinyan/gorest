package channels

import (
	"fmt"
	"math/rand"
	"time"
)

type App interface {
	WriteChan()
	PrintChan()
}

type app struct {
	OutChan chan int
	OutChan2 chan int
	OutChan3 chan int
}

func (a *app) WriteChan() {
	for  {
		a.OutChan <- rand.Int()
		a.OutChan2 <- rand.Int()
		a.OutChan3 <- rand.Int()

		time.Sleep(1 * time.Millisecond)
	}
}

func (a *app) PrintChan() {
	for  {

		select {
		case c := <-a.OutChan:
			fmt.Println("OutChan: ",c)
		case c2 := <-a.OutChan2:
			fmt.Println("OutChan2: ", c2)
		case <-a.OutChan3:
			fmt.Println("OutChan3")
		}

		time.Sleep(2 * time.Millisecond)
	}
}

func NewApp() App {
	return &app{
		OutChan: make(chan int, 10),
		OutChan2: make(chan int, 10),
		OutChan3: make(chan int, 10),
	}
}

func main() {
	fmt.Println("CHANNELS")

	app := NewApp()
	go app.WriteChan()
	app.PrintChan()
}
