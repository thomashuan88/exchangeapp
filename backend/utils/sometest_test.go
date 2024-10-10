package utils

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTaskControl(t *testing.T) {
	tasknum := 5

	wg := sync.WaitGroup{}
	for i := 0; i < tasknum; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

func TestXX1(t *testing.T) {
	test := make(chan int, 10)

	go func() {
		for {
			select {
			case val, ok := <-test:
				if !ok {
					t.Logf("chan closed")
					return
				}
				t.Logf("val: %d", val)
			}
		}

		// for val := range test {
		// 	t.Logf("val: %d", val)
		// }
	}()

	go func() {
		test <- 1
		time.Sleep(1 * time.Second)
		test <- 2
		close(test)
	}()

	time.Sleep(5 * time.Second)
}

func TestXX2(t *testing.T) {
	test := make(chan int, 5)
	exit := make(chan struct{})

	go func() {
		for {
			select {
			case val := <-test:
				t.Logf("val: %d", val)
			case <-exit:
				t.Logf("exit")
				return
			}
		}
	}()

	go func() {
		test <- 1
		time.Sleep(1 * time.Second)
		test <- 2
		close(exit)
	}()

	time.Sleep(5 * time.Second)
}

func TestXX3(t *testing.T) {
	test := make(chan int, 5)

	go func() {
		for {
			select {
			case val := <-test:
				t.Logf("val: %d", val)
			case <-time.After(2 * time.Second):
				t.Logf("timeout")
				return
			}

		}
	}()

	go func() {
		test <- 1
		time.Sleep(2 * time.Second)
		test <- 2
	}()

	time.Sleep(5 * time.Second)
}

func TestXX4(t *testing.T) {
	a := context.Background()
	b, cancel := context.WithCancel(a)
	b = context.WithValue(b, "k1", "v1")
	c := context.WithValue(b, "k2", "v2")
	d := context.WithValue(c, "k3", "v3")
	e := context.WithValue(d, "k3", "v4")
	cancel()
	f := context.WithValue(e, "k5", "v5")

	fmt.Printf("d: %s\n", d.Value("k3"))
	fmt.Printf("e: %s\n", e.Value("k3"))
	fmt.Printf("f: %s\n", f.Value("k2"))
}

func TestXX5(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				t.Logf("Context cancelled!")
				return
			}
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	time.Sleep(2 * time.Second)
}

func task(name string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("task %s start\n", name)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("task %s exit\n", name)
			return
		default:
			fmt.Printf("task %s running\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func TestXX6(t *testing.T) {
	ctxA, cancelA := context.WithCancel(context.Background())

	ctxB, cancelB := context.WithCancel(ctxA)
	ctxC, cancelC := context.WithCancel(ctxA)
	ctxD, _ := context.WithCancel(ctxA)

	ctxE, _ := context.WithCancel(ctxB)
	ctxF, _ := context.WithCancel(ctxB)

	ctxG, _ := context.WithCancel(ctxC)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go task("A", ctxA, &wg)

	wg.Add(1)
	go task("B", ctxB, &wg)

	wg.Add(1)
	go task("C", ctxC, &wg)

	wg.Add(1)
	go task("D", ctxD, &wg)

	wg.Add(1)
	go task("E", ctxE, &wg)

	wg.Add(1)
	go task("F", ctxF, &wg)

	wg.Add(1)
	go task("G", ctxG, &wg)

	time.Sleep(2 * time.Second)

	cancelB()
	time.Sleep(1 * time.Second)

	cancelC()
	time.Sleep(1 * time.Second)

	cancelA()
	time.Sleep(1 * time.Second)

	wg.Wait()
	t.Log("all task stopped")
}
