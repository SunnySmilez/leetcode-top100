package interview

import (
	"sync"
)

// 多线程顺序输出abc
func Concurrent() {
	var wg sync.WaitGroup
	wg.Add(3)

	a := make(chan bool)
	b := make(chan bool)
	c := make(chan bool, 1)

	c <- true

	go printA(&wg, c, a)
	go printB(&wg, a, b)
	go printC(&wg, b, c)

	wg.Wait()

}

func printA(wg *sync.WaitGroup, c chan bool, a chan bool) {
	for i := 0; i < 3; i++ {
		if <-c == true {
			print("A")
			a <- true
		}
	}

	wg.Done()
}

func printB(wg *sync.WaitGroup, a chan bool, b chan bool) {
	for i := 0; i < 3; i++ {
		if <-a == true {
			print("B")
			b <- true
		}
	}

	wg.Done()
}

func printC(wg *sync.WaitGroup, b chan bool, c chan bool) {
	for i := 0; i < 3; i++ {
		if <-b == true {
			print("C")
			c <- true
		}
	}

	wg.Done()
}
