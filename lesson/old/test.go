package old

import (
	"fmt"
	"strconv"
)

func main() {
	//MyArr()
	//MySlice()
	//MyMap()
}

func MyMap()  {

	var a = map[string]int{}

	for i := 1; i < 10; i++ {
		a[strconv.Itoa(i)] = i+100
	}
	fmt.Println(a)

	delete(a, "2")

	fmt.Println(a["2"])

	if _, ok := a["2"]; ok {
		fmt.Println("v, ok := a[\"2\"]; ok ")
	}

	fmt.Println(a)

}

func MySlice() {

	var a []int

	a = make([]int, 5)

	for i := 0; i < 10; i++ {
		a = append(a, i+100)
	}

	fmt.Println("a", a, "len: ", len(a), "cap: ", cap(a))

	b := append(a[:1], a[6:10]...)

	fmt.Println(b)

}

func MyArr()  {

	var a [5]int

	a1 := [2]int{1,2}

	a[0] = 1
	a[4] = 4

	fmt.Println(a)
	fmt.Println("len: ", len(a), "cap: ", cap(a))
	fmt.Println("a1: ", a1)
}