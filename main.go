package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a 
	}
	return b
}

func TestUntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMInTableDrive(t *testing.T) {
	var tests = [] struct {
		a, b int
		want int
	}{ 
		{0,1,0},

