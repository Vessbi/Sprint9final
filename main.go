package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return nil
	}
	sliceElem := make([]int, 0, size)

	for i := 0; i < size; i++ {
		sliceElem = append(sliceElem, rand.Int())
	}

	return sliceElem
}

func maximum(data []int) int {

	if len(data) <= 0 {
		return 0
	}
	maxNumber := 0
	for _, v := range data {
		if maxNumber < v {
			maxNumber = v
		}
	}

	return maxNumber
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	//result := make(chan int, CHUNKS)
	result := make([]int, CHUNKS)

	if len(data) < CHUNKS {
		return maximum(data)
	}

	for i := 0; i < CHUNKS; i++ {

		idx1 := len(data) / CHUNKS * i
		idx2 := len(data) / CHUNKS * (i + 1)
		sl1 := data[idx1:idx2]
		wg.Add(1)
		go func(sl1 []int, i int) {
			defer wg.Done()

			result[i] = maximum(sl1)

		}(sl1, i)

	}
	wg.Wait()

	return maximum(result)

}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	RE := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(RE)
	end := time.Now()
	elapsed := time.Duration.Milliseconds(end.Sub(start))
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(RE)
	end = time.Now()
	elapsed = time.Duration.Milliseconds(end.Sub(start))
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
