package main

import (
	"fmt"
	"math"
	"service"
	"time"
)

const (
	Big   = 1 << 3
	Small = Big >> 1
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v,%s", e.When, e.What)
}
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type IPAddr [4]byte

func (ip *IPAddr) Stringer() string {
	var ips = ""
	for i := 0; i < 4; i++ {
		var one string
		one = string(ip[i])
		ips = ips + one
		if i < 3 {
			ips = ips + "."
		}
	}
	return ips
}
func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Println(a.Stringer())
		fmt.Printf("%v: %v\n", n, a)
	}
	return

	if err := run(); err != nil {
		fmt.Println(err)
	}
	now := time.Now()
	fmt.Println(now)
	return

	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	//delete(m, "Answer")
	//fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 5))

	fmt.Println(Big * 0.1)
	fmt.Println(Small*10 + 1)
	fmt.Println("Hello World")
	var sum int = 0
	for i := 1; i < 10; i++ {
		if i > 1 {
			fmt.Print("+")
		}
		fmt.Print(i)
		sum = sum + i
	}
	fmt.Print("=")
	fmt.Println(sum)
	var p []int
	p = []int{2, 3, 4, 7, 11}
	fmt.Println("p==", p)
	p = append(p, 0)
	fmt.Println("p append==", p)
	var a []int
	a = append(a, 1)
	fmt.Println("a append==", a)
	fmt.Println("Rang:")
	var pow = []int{1, 2, 4, 8, 16}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	fmt.Println("Fibonacci:")
	var f = fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i, f(i))
	}

	aPerson := Person{"Arthur Dent", 42}
	fmt.Println(aPerson)

	fmt.Println("http get ")
	//resp, err := http.Get("http://www.baidu.com/")
	service.Test()
	service.Format(10, 90)
}
func fibonacci() func(int) int {
	preOne := 0
	thisOne := 0
	return func(x int) int {
		if x < 2 {
			thisOne = x
		} else {
			var tmp int
			tmp = thisOne
			thisOne = thisOne + preOne
			preOne = tmp
		}
		return thisOne
	}
}

type Person struct {
	Name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.age)
}

func makeChinese() {

	fmt.Println()
}
