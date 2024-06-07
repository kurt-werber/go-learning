package main

import (
	"fmt"
	"math"
	"runtime"
	"strings"
	"time"
)

type Vertex struct {
	X float64
	Y float64
}

// has a receiver arg, therefore is a method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// nonstruct types can also have methods
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// pointer receivers can be used to modify the value the pointer points to
// pointer receivers better most of the time
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

const (
	google    = "google"
	bell_labs = "bell labs"
)

var m map[string]Vertex

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	v.Scale(3)
	fmt.Println(v)
}

func fibonacci() func() int {
	twoAgo, oneAgo := 0, 1
	offset := 0
	return func() int {
		switch offset {
		case 0:
			offset += 1
			return 0
		case 1:
			offset += 1
			return 1
		default:
			sum := twoAgo + oneAgo
			twoAgo = oneAgo
			oneAgo = sum
			return sum
		}
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func printSlice(s []int) {
	fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))
	fmt.Println(s)
}

func WordCount(s string) map[string]int {
	splitString := strings.Fields(s)
	wordCountMap := make(map[string]int)
	for _, v := range splitString {
		if elem, present := wordCountMap[v]; present {
			wordCountMap[v] = elem + 1
		} else {
			wordCountMap[v] = 1
		}
	}
	return wordCountMap
}

func oldStuff() {
	fmt.Println("Hello world!")
	Sqrt(2)
	fmt.Println(abs(-5))
	fmt.Println(math.Sqrt(2))
	fmt.Println(runtime.GOOS)
	fmt.Println(time.Now())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Saturday)
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("Yay")
	case today + 1:
		fmt.Println("Almost")
	default:
		fmt.Println("dang")
	}
	defer fmt.Println("sorry I'm late")
	defer fmt.Println("I'm later but I cheated")
	fmt.Println("be there soon")
	i := 5
	iPointer := &i
	fmt.Println(iPointer)
	fmt.Println(*iPointer)
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	a := make([]int, 5, 10)
	printSlice(a)
	a = a[:8]
	printSlice(a)
	for i := 0; i < 100; i++ {
		a = append(a, i)
		printSlice(a)
	}
	for i, v := range a {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
	for _, v := range a {
		fmt.Printf("Value: %d\n", v)
	}
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40, -74}
	fmt.Println(m["Bell Labs"])

	mapLiteral := map[string]Vertex{
		bell_labs: {
			40, -74,
		},
		google: {
			37, -122,
		},
	}
	fmt.Println(mapLiteral)

	delete(mapLiteral, google)
	elem, ok := mapLiteral[google]
	fmt.Printf("The key %s is in the map: %t. Therefore its value is %v\n", google, ok, elem)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
