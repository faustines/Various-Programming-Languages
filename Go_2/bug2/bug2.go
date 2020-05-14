/*
ECS 140A Programming Hw5
Faustine Yiu 913062973
Source: https://golang.org, ECS 140A Go Lecture, StackOverflow
*/

package bug2

import (
	"sync"
)

func bug2(n int, foo func(int) int, out chan int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(a int) {
			out <- foo(a)
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(out)
}
