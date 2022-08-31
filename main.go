package main

import (
	"fmt"
	"sync"
)

var n = 1
var m = []string{"0"}

func testConcurrent() {
	wg := &sync.WaitGroup{}
	wg.Add(2000)
	for i := 0; i < 2000; i++ {
		go func(wg *sync.WaitGroup, i int) {
			n += 1
			// m = append(m, fmt.Sprintf("%d", i))
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
}

func main() {
	testConcurrent()
	// time.Sleep(time.Second * 10)
	fmt.Println("n: ", n)
	// fmt.Println("len(m): ", len(m))
}
