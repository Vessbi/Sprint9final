package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле

func TestMaximum(t *testing.T) {

	var flagtest = []struct {
		in  []int
		out int
	}{
		{nil, 0},
		{[]int{5, 10, 67, 350}, 350},
		{[]int{10}, 10},
		{[]int{0}, 0},
	}
	for _, tt := range flagtest {
		assert.Equal(t, maximum(tt.in), tt.out)
	}
}

func TestGenerateRandomElements(t *testing.T) {
	var flagtest = []struct {
		in  int
		out []int
	}{
		{0, nil},
		{-10, nil},
	}

	for _, tt := range flagtest {
		assert.Equal(t, generateRandomElements(tt.in), tt.out)
		assert.Equal(t, cap(generateRandomElements(tt.in)), cap(tt.out))
	}

}
