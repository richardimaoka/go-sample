package main

import "fmt"

type (
	Me struct {
		V1 int
		V2 string
		V3 int
		V4 string
	}
)

func New(v1 int, v2 string, opts ...func(me *Me)) *Me {
	me := new(Me)

	me.V1 = v1
	me.V2 = v2
	me.V3 = -1
	me.V4 = "unknown"

	for _, opt := range opts {
		opt(me)
	}

	return me
}

func main() {
	m1 := New(1, "hello")
	fmt.Printf("%v\n", m1)

	m2 := New(2, "world", func(me *Me) { me.V4 = "golang" })
	fmt.Printf("%v\n", m2)

	m3 := New(3, "naruhodo", func(me *Me) { me.V2 = "golang" })
	fmt.Printf("%v\n", m3)
}
