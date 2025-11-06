package main

import "fmt"

func main() {
	// Decalring a Chanels
	ch1 := make(chan int)
	ch2 := make(chan int)

	sin := make(chan struct{}, 1)
	sin <- struct{}{}

	go func() {
		defer close(ch1)
		for range 20 {
			<-sin
			ch1 <- 1
			sin <- struct{}{}

		}
	}()

	go func() {
		defer close(ch2)
		for range 20 {
			<-sin
			ch2 <- 2
			sin <- struct{}{}
		}
	}()

	var nums []int

	for ch1 != nil || ch2 != nil {

		select {
		case i, ok := <-ch1:
			if !ok {
				ch1 = nil
			} else {
				nums = append(nums, i)
			}
		case s, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				nums = append(nums, s)
			}
		}
	}

	// nums := <-ch1
	fmt.Println(nums)
}
