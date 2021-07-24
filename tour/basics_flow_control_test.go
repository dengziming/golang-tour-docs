package tour

import (
	"fmt"
	"math"
	"runtime"
	"testing"
	"time"
)

func TestFor(t *testing.T) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1

	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func TestIf(t *testing.T) {
	fmt.Println(sqrt(2), sqrt(-4))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func TestIfWithStat(t *testing.T) {
	fmt.Println(
		pow(3,2, 10),
		pow(3,3, 20),
		)
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func TestIfElse(t *testing.T) {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z * z - x) / (2 * z)
	}
	return z
}

func TestLoopsAndFunctionsExercise(t *testing.T) {
	fmt.Println(Sqrt(2))
}

func TestSwitch(t *testing.T) {
	fmt.Println("Go runs on")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func TestSiwtchEvaluationOrder(t *testing.T) {
	fmt.Println("When's Saturday")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

func TestSwitchWithNoCondition(t1 *testing.T) {

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func TestDefer(t *testing.T) {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func TestStackintDefer(t *testing.T) {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
