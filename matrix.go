package main

import ()

type Matrix struct {
	size uint
	data []int
}

func NewMatrix(n uint) Matrix {
	data := make([]int, n * n)
	for i,_ := range data {
		data[i] = 0
	}
	return Matrix {n, data}
}