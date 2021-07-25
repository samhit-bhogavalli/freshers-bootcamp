package main

import (
	"fmt"
	"sync"
)

type counter struct {
	frequency map[string]int
	mutex     sync.Mutex
	workers   int
	waitGroup sync.WaitGroup
	input     []string
}

//function for counting alphabets in a job
func (cnt *counter) worker(jobs chan string) {
	defer cnt.waitGroup.Done()
	for job := range jobs {
		for _, rune := range job {
			cnt.mutex.Lock()
			cnt.frequency[string(rune)]++
			cnt.mutex.Unlock()
		}
	}
}

//a function which takes slice of strings and calculates frequency
func (cnt *counter) getCount() {
	jobs := make(chan string, 3)
	for i := 0; i < cnt.workers; i++ {
		cnt.waitGroup.Add(1)
		go cnt.worker(jobs)
	}
	for _, val := range cnt.input {
		jobs <- val
	}
	close(jobs)
	cnt.waitGroup.Wait()
}

func main() {
	//testing
	input := []string{"quick", "quick", "quick", "quick", "quick", "quick"}
	frequencyCounter := counter{make(map[string]int), sync.Mutex{}, 3, sync.WaitGroup{}, input}
	frequencyCounter.getCount()
	fmt.Println(frequencyCounter.frequency)
}
