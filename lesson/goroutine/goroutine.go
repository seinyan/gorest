package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type App interface {
	Read(str int)
	Print()
}

type app struct {
	MapMutex *sync.Mutex
	Map map[int]int
}

func (a app) Read(str int) {
	a.MapMutex.Lock()
	defer a.MapMutex.Unlock()
	a.Map[1] = str
}

func (a app) Print() {
	fmt.Println("Print: ", a.Map[1])
}

func NewApp() App {
	return &app{
		MapMutex: new(sync.Mutex),
		Map: make(map[int]int),
	}
}

func main() {
	fmt.Println("goroutine.go")

	app := NewApp()

	go func() {
		for {
			app.Read(rand.Int())
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			app.Read(rand.Int())
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			app.Print()
			time.Sleep(3 * time.Millisecond)
		}
	}()

	for  {
		time.Sleep(3 * time.Second)
	}
}
