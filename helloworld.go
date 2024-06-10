package main

import (
	"fmt"
	"io"
	"math"
	"runtime"
	"strings"
	"time"
)

type Vertex struct {
	X float64
	Y float64
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string { //error is a built in interface containing Error()
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type IPAddr [4]byte

func (ipaddr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
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
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func doSomething(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%d is an int\n", v)
	case string:
		fmt.Printf("%q is a string\n", v)
	default:
		fmt.Println("something else")
	}
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

	ver := Vertex{3, 4}
	fmt.Println(v.Abs())
	myF := MyFloat(-math.Sqrt2)
	fmt.Println(myF.Abs())
	ver.Scale(3)
	fmt.Println(v)
	//type assertion provides access to interface value's underlying concrete value
	var inter interface{} = "hello"
	s := inter.(string)
	fmt.Println(s)
	u, ok := inter.(float64) //would fail without ok variable because type assertion fails
	fmt.Println(u, ok)
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip) //by adding String() method to IPAddr, fmt can convert to string (it now implements Stringer interface)
	}
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
