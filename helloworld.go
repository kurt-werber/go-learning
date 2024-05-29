package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

type Vertex struct {
	X int
	Y int
}

func main() {
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
}

func printSlice(s []int) {
	fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))
	fmt.Println(s)
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
}
