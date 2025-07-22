package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10}
	n := len(arr)
	chunkSize := 3
	numChunks := (n + chunkSize - 1) / chunkSize
	wg := &sync.WaitGroup{}
	resultChan := make(chan int, numChunks)
	for i := 0; i < n; i = i + chunkSize {
		start := i
		end := start + chunkSize
		if end > n {
			end = n
		}
		wg.Add(1)
		go sumChunk(arr[start:end], resultChan, wg)

	}

	wg.Wait()
	close(resultChan)

	totalSum := 0
	for val := range resultChan {
		totalSum += val
	}
	fmt.Println("total sum:", totalSum)

}

func sumChunk(chunkArr []int, resultChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, num := range chunkArr {
		sum += num
	}
	fmt.Println("chunk sum : ", sum)
	resultChan <- sum
}
