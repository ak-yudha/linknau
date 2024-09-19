// process.go
package main

import (
	"fmt"
	"sync"
)

type Processor interface {
	Process(number int) int
}

type SquareProcessor struct{}

func (p *SquareProcessor) Process(number int) int {
	return number * number
}

// ProcessNumbers memproses daftar angka secara asinkron.
func ProcessNumbers(numbers []int, processor Processor) []int {
	var wg sync.WaitGroup
	results := make([]int, len(numbers))
	resultsChan := make(chan int, len(numbers))

	for _, number := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			result := processor.Process(num)
			resultsChan <- result
		}(number)
	}

	wg.Wait()
	close(resultsChan)

	i := 0
	for result := range resultsChan {
		results[i] = result
		i++
	}

	return results
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	processor := &SquareProcessor{}
	results := ProcessNumbers(numbers, processor)
	fmt.Println("Squared numbers:", results)
}
