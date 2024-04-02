package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Person struct {
	Name string
}

func (p Person) ChangeName(newName string) {
	p.Name = newName
}

func trick1() {
	p := Person{Name: "John"}
	p.ChangeName("Jane")
	fmt.Println(p.Name)
}

func trick2() {
	nums := []int{1, 2, 3, 4, 5}
	slice := nums[1:3]
	slice = append(slice, 6)
	slice[0] = 10
	fmt.Println(nums)
}

type MyStruct struct {
	Name string
}

func trick3() {
	slice := make([]MyStruct, 2)
	slice[0].Name = "John"
	slice[1].Name = "Jane"
	fmt.Println(slice)
}

func trick4() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
}

func trick5() {
	slice := make([]int, 3, 5)
	slice = append(slice, 1, 2, 3)
	fmt.Println(len(slice), cap(slice))
}

func trick6() {
	var m map[string]int
	m["one"] = 1
	m["two"] = 2
	fmt.Println(len(m))
}

func trick7() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 5)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func trick8() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("Goroutine 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Goroutine 2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("All goroutines completed")
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
}

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func trick9() {
	counter := &Counter{}
	var wg sync.WaitGroup

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.count)
}

// func trick10() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := 1; i <= 3; i++ {
// 			ch <- i
// 			time.Sleep(time.Second)
// 		}
// 		close(ch)
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := range ch {
// 			fmt.Println(i)
// 			time.Sleep(time.Second)
// 		}
// 	}()

// 	wg.Wait()
// }

func trick11() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}

}

func trick12() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func trick13() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	// trick1()
	// trick2()
	// trick3()
	// trick4()
	// trick5()
	// trick7()
	// trick8()
	// trick9()
	// trick10()
	// trick11()
	// trick12()
	trick13()
}
