package tour

import (
	"fmt"
	"math"
	"testing"
)

func (v Vertex) Abs() float64 {
	return math.Sqrt(float64(v.X * v.X) + float64(v.Y * v.Y))
}

func TestMethods(t *testing.T) {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

func Abs(v Vertex) float64 {
	return math.Sqrt(float64(v.X * v.X) + float64(v.Y * v.Y))
}

func TestMethodsAreFunctions(t *testing.T) {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func TestMethods2(t *testing.T) {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func (v *Vertex) Scale(f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestPointerReceivers(t *testing.T) {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

func Scale(v *Vertex, f int) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestPointersAndFunctions(t *testing.T) {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}


// TODO Methods and pointer indirection


type Abser interface {
	Abs() float64
}


func TestInterfaces(t *testing.T) {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())

	a = v
	fmt.Println(a.Abs())
}

// Ignored: Interfaces are implemented implicitly

