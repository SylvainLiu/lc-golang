package main

import (
	"testing"
)

func TestMultiDimSlice(t *testing.T) {
	slice := make([][]string, 1)
	for i := range slice {
		slice[i] = make([]string, 1)
	}
}
