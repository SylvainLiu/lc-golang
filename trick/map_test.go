package main

import (
	"fmt"
	"testing"
)

func TestMapRange(t *testing.T) {
	var counter int
	for counter != 4 {
		counter = 0

		var m = map[string]int{
			"A": 21,
			"B": 22,
			"C": 23,
		}

		for k := range m {
			if counter == 0 {
				// delete(m, "A")
				m["D"] = 24
			}

			fmt.Print(k)
			counter++
		}
		fmt.Println(" ", counter)
	}
}
