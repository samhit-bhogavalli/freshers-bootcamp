package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//function to ask students and get their replies
func askStudent(waitGroup *sync.WaitGroup, ratings chan int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	ratings <- rand.Intn(1000)
	waitGroup.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	ratings := make(chan int, 5)
	for i := 0; i < 200; i++ {
		waitGroup.Add(1)
		go askStudent(&waitGroup, ratings)
	}

	totalRating := 0
	count := 0
	done := make(chan bool)
	go func() {
		for rating := range ratings {
			totalRating += rating
			count++
		}
		done <- true
	}()

	waitGroup.Wait()
	close(ratings)
	<-done
	fmt.Println(totalRating / count)
}
