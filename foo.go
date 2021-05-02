package main

type Foo interface {
	Bar(x int) int
}

type Bar interface {
	Sub(y string) int
}
