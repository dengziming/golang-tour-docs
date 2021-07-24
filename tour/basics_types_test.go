package tour

import (
	"fmt"
	"golang.org/x/tour/wc"
	"math"

	//"golang.org/x/tour/pic"
	"strings"
	"testing"
)


// TestPointers test pointer
func TestPointers(t *testing.T) {
	i, j := 42, 2701

	// The type *T is a pointer to a T value. Its zero value is nil.
	p := &i
	fmt.Println(*p)

	// The * operator denotes the pointer's underlying value.
	*p = 21
	fmt.Println(i)

	p = &j

	*p = *p / 37
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

func TestVertex(t *testing.T) {
	v := Vertex{1, 2}
	fmt.Println(v)

	v.Y = 4
	fmt.Println(v)

	p := &v

	p.X = 1e9
	(*p).Y = 2e3

	fmt.Println(*p)
}

func TestStructLiterals(t *testing.T) {

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p = &Vertex{1, 2}
	)

	fmt.Println(v1, p, v2, v3)
}

func TestArrays(t *testing.T) {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int {2,3,5,7,11,13}
	fmt.Println(primes)
}

func TestSlices(t *testing.T) {
	primes := [6]int {2,3,5,7,11,13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func TestSlicesPointers(t *testing.T) {

	names := [4]string {
		"nolan",
		"batman",
		"black widow",
		"wolve",
	}
	fmt.Println(names)

	a := names[0: 2]

	b := names[1: 3]

	fmt.Println(a, b)

	b[0] = "twins"

	fmt.Println(a, b)
	fmt.Println(names)
}

func TestSlicksLiterals(t *testing.T) {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(s)
}

func TestSlickDefaults(t *testing.T) {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func TestSliceLengthAndCapacity(t *testing.T) {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func TestNilSlice(t *testing.T) {
	var s []int
	fmt.Println(s ,len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func TestMakingSlice(t *testing.T) {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func TestSliceOfSlice(t *testing.T) {

	board := [][]string{
		[]string{"", "", ""},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][2] = "X"
	board[2][2] = "0"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func TestAppendSlice(t *testing.T) {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2,3,4)
	printSlice(s)
}

func TestRange(t *testing.T) {
	var pow = []int{1, 2,4, 8, 16, 32, 64, 128}
	for i, v := range  pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow = make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func Pic(dx, dy int) [][] uint8 {
	var pic = make([][]uint8, dx)
	for _, dys := range pic {
		dys[0] = 1
	}
	return pic
}

// TODO
func TestSliceExercise(t *testing.T) {
	//pic.Show(Pic)
}

func TestMaps(t *testing.T) {
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40, 1}
	fmt.Println(m["Bell Labs"])
}

func TestMapLiterals(t *testing.T) {
	var m = map[string]Vertex{
		"Bell Labs": {1, 2},
		"Google": Vertex{3, 4},
	}

	fmt.Println(m)
}

func TestMutatingMapps(t *testing.T) {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println(m)

	m["Answer"] = 48
	fmt.Println(m)

	delete(m, "Answer")
	fmt.Println(m)

	v, ok := m["Answer"]
	fmt.Println(v, ok)
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _,token := range strings.Fields(s) {
		v, ok := m[token]
		if !ok {
			m[token] = 1
		} else {
			m[token] = v + 1
		}
	}
	return m
}

func TestMapsExercise(t *testing.T) {
	wc.Test(WordCount)
}

func compute(fn func(float64, float64) float64) float64 {
return fn(3, 4)
}

func TestFuncValues(t *testing.T) {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x * x + y * y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestFunctionClosures(t *testing.T) {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(- 2 * i),
			)
	}
}

func fibonacci() func() int {
	value1, value2 := 0, 1

	return func() int {
		sum := value1 + value2
		value1 = value2
		value2 = sum
		return sum
	}
}

func TestFibonacciClosure(t *testing.T) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

