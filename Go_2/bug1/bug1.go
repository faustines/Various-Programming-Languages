/*
ECS 140A Programming Hw5
Faustine Yiu 913062973
Source: https://golang.org
*/

package bug1
import (
	"sync"
)

// Counter stores a count.
type Counter struct {
	n int64
	mux sync.Mutex
}

// Inc increments the count in the Counter.
func (c *Counter) Inc() {
	 c.mux.Lock()
	 c.n++
	 c.mux.Unlock()
}
