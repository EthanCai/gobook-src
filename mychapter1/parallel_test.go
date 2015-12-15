package mychapter1

import (
	"testing"
	"fmt"
	"time"
)

var gNum uint

func TestCase1(t *testing.T) {
	time.Sleep(2 * 1000 * 1000 * 1000)
	for i := 1; i <= 10000; i++ {
		gNum++;
		fmt.Println("TestCase1 gNum =", gNum)
	}
}

func TestCase2(t *testing.T) {
	time.Sleep(2 * 1000 * 1000 * 1000)
	gNum++;
	fmt.Println("TestCase2 gNum =", gNum)
}
