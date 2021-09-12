package tour

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sync"
	"testing"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func TestGoroutines(t *testing.T) {
	go say("world")
	say("hello")
}


func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func TestChannels(t *testing.T) {
	s := []int{7, 2, 8 -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/ 2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <- c

	fmt.Println(x, y, x + y)
}

func TestBufferedChannels(t *testing.T) {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println( <- ch)
	fmt.Println( <- ch)
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x ,y = y , x + y
	}
	close(c)
}

func TestRangeAndClose(t *testing.T) {
	c := make(chan int, 10)

	go fibonacci2(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci3(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit :
			fmt.Println("quit")
			return
		}
	}
}

func TestSelect(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<- c)
		}
		quit <- 0
	}()

	fibonacci3(c, quit)
}

func TestDefaultSelection(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("       .")
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// exercise: equivalent binary trees

func doWalk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	doWalk(t.Left, ch)
	ch <- t.Value
	doWalk(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int)  {
	doWalk(t, ch)
	close(ch)
}

func TestWalk(t *testing.T) {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if j, ok := <- ch2; !ok || j != i {
			return false
		}
	}

	_, ok := <-ch2
	if ok {
		return false
	}
	return true
}

func TestEquivalentBinaryTrees(t *testing.T) {

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(3), tree.New(3)))
	fmt.Println(Same(tree.New(4), tree.New(4)))
	fmt.Println(Same(tree.New(5), tree.New(5)))

	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(3)))
	fmt.Println(Same(tree.New(3), tree.New(1)))
}


type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func TestMutex(t *testing.T) {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("someKey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("someKey"))
}

// TODO Exercise: Web Crawler
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

