package main

import (
	"container/list"
)

type Timsort struct{}

func getMinRunLength(n int) int {
	flag := 0
	for n >= 64 {
		flag |= n & 1
		n >>= 1

	}
	return n + flag
}

func (timsort Timsort) BinarySearch(array []int, element int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := (left + right) / 2
		if mid < element {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func (timsort Timsort) InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
	return array
}

func (timsort Timsort) ReverseArray(array []int) []int {
	left, right := 0, len(array)-1
	for left < right {
		if array[left] > array[right] {
			array[left], array[right] = array[right], array[left]
		}
		left++
		right--
	}
	return array
}

func (timsort Timsort) Sort(slices []int) []int {
	minRun := getMinRunLength(len(slices))
	index := 0
	runs := list.New()
	run := make([]int, 0)
	FlagLenRun := false
	if len(slices) == 1 {
		return slices
	}
	for index < len(slices) {
		index += 2
		run = slices[index-2 : index]
		if index >= len(slices) {
			if slices[index-2] > slices[index-1] {
				run = timsort.ReverseArray(slices)
			}
			runs.PushBack(run)
		}
		if slices[index-2] > slices[index-1] {
			for slices[index-1] > slices[index] {
				run = append(run, slices[index])
				index++
				if index > len(slices)-1 {
					break
				}
			}
		} else if slices[index-2] <= slices[index-1] {
			for slices[index-1] <= slices[index] {
				run = append(run, slices[index])
				index++
				if index > len(slices)-1 {
					break
				}
			}
		}
		for len(run) < minRun && index <= len(slices)-1 {
			FlagLenRun = true
			run = append(run, slices[index])
			index++
		}
		if FlagLenRun == true {
			run = timsort.InsertionSort(run)
		} else {
			run = timsort.ReverseArray(run)
		}
		runs.PushBack(run)
	}
	return slices
}
