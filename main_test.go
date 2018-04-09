package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortOne(t *testing.T) {
	expected := []int{1}
	list := []int{1}
	MergeSort(list)
	assert.Equal(t, expected, list)
}

func TestSortTwo(t *testing.T) {
	expected := []int{1, 2}
	list := []int{2, 1}
	MergeSort(list)
	assert.Equal(t, expected, list)
}

func TestSortThree(t *testing.T) {
	expected := []int{1, 2, 3}
	list := []int{2, 1, 3}
	MergeSort(list)
	assert.Equal(t, expected, list)
}

func TestSortMatching(t *testing.T) {
	expected := []int{1, 2, 2}
	list := []int{2, 1, 2}
	MergeSort(list)
	assert.Equal(t, expected, list)
}

func TestSortLong(t *testing.T) {
	expected := []int{1, 2, 2, 4, 5, 9, 12, 18, 19, 20}
	list := []int{4, 1, 2, 20, 12, 9, 2, 5, 18, 19}
	MergeSort(list)
	assert.Equal(t, expected, list)
}
