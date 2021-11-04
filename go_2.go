package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go executeTask(i, &wg)
	}
	wg.Wait()
	fmt.Println("Total execution time:", time.Since(start))
}

func executeTask(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Executing task", i)
	time.Sleep(250 * time.Millisecond)
}
