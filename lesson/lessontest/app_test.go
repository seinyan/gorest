package lessontest

import (
	"fmt"
	"testing"
)

var testApp App

//func TestMain(m *testing.M) {
//	testApp = NewApp()
//	fmt.Println("TestMain")
//}

func TestApp_Calc(t *testing.T) {
	testApp := NewApp()

	fmt.Println(testApp)

	t.Error(testApp.Calc(1,2))
	t.Error(testApp.Plus(1,2))
}



/*
func TestPlus(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		var result, x, y = 4, 2, 2

		realTestResult := Plus(x, y)
		if result != realTestResult {
			t.Error("incorrect func")
		}
	})

	t.Run("negative", func(t *testing.T) {
		var result, x, y = 8, 4, 1

		realTestResult := Plus(x, y)
		if result == realTestResult {
			t.Error("incorrect func")
		}
	})


}
 */