package main

import "testing"

// Пишите тесты в этом файле

func TestMaximum(t *testing.T) {
	randSlice := [][]int{
		{5, 10, 67, 350, -1},
		{10},
		{},
	}
	// ожидаемые значения для каждой последовательности
	max := []int{350, 10, 0}

	for i, list := range randSlice {
		if maximum(list) != max[i] {
			t.Error(i, ":", maximum(list), "!=", max[i])
		}
	}
}

func TestGenerateRandomElements(t *testing.T) {
	if generateRandomElements(0) != nil {
		t.Error("slide size cannot be less than 0")
	}
}
