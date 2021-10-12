package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type dataChannel2 struct {
	values []float64
	err    error
}

func main() {
	ch1 := make(chan []int)
	ch2 := make(chan dataChannel2)
	ch3 := make(chan []string)

	var l1 []int
	var l2 dataChannel2
	var l3 []string

	var wg sync.WaitGroup
	wg.Add(3)
	go getList1(ch1)
	go getList2(ch2)
	go getList3(ch3)

	go func(chan []int, sync.WaitGroup, *[]int) {
		l1 = <-ch1
		wg.Done()
	}(ch1, wg, &l1)

	go func(chan dataChannel2, sync.WaitGroup, *dataChannel2) {
		v := <-ch2
		l2.err = v.err
		l2.values = v.values

		wg.Done()
	}(ch2, wg, &l2)

	go func(chan []string, sync.WaitGroup, *[]string) {
		l3 = <-ch3
		wg.Done()
	}(ch3, wg, &l3)

	wg.Wait()

	if l2.err != nil {
		fmt.Println("Error on get data on channel 2")
		return
	}

	fmt.Printf("total list 1: %d \n", len(l1))
	fmt.Printf("total list 2: %d \n", len(l2.values))
	fmt.Printf("total list 3: %d \n", len(l3))

}

func getList1(ch chan []int) {
	defer close(ch)
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(10000)
	var list []int
	fmt.Printf("The first list will have %d itens \n ", r)
	time.Sleep(5 * time.Second)
	for x := 0; x < r; x++ {
		list = append(list, x)
	}

	ch <- list
}

func getList2(ch2 chan dataChannel2) {
	defer close(ch2)
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(1000)
	var list []float64
	fmt.Printf("The second list will have %d itens \n ", r)
	for x := 0; x < r; x++ {
		list = append(list, float64(x))
	}
	dt := dataChannel2{
		err:    nil,
		values: list,
	}

	ch2 <- dt
}

func getList3(ch3 chan []string) {
	defer close(ch3)
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(15000)
	var list []string
	fmt.Printf("The third list will have %d itens \n", r)

	for x := 0; x < r; x++ {
		list = append(list, fmt.Sprintf("line: %d", x))
	}

	ch3 <- list
}
