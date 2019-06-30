package main

import (
	"fmt"
	"time"
)

func main() {
	//Insertion sort

	b := []int{34, 19, 72, 52, 41, 34, 8}

	fmt.Printf("Before insertion sort: %v\n", b)

	/* We must prove these things about the loop invariant
	1. Initialization - It is true before the first iteration of the loop
	2. Maintenance - It is true before an iteration of the loop and remains true before the next iteration
	3. Termination - When the loop terminates, the invariant gives us a useful property to show the algorithm is correct
	*/

	t0 := time.Now()
	for j := 1; j < len(b); j++ {
		key := b[j]
		i := j - 1
		for i >= 0 && b[i] < key {
			b[i+1] = b[i]
			i -= 1

		}
		b[i+1] = key

	}
	duration := time.Since(t0).Nanoseconds()

	fmt.Printf("After insertion sort: %v\n", b)
	fmt.Println(duration)

	// Selection sort: An array of size n, find the smallest element of A and exchange it with element in A[i]

	A := []int{34, 19, 72, 52, 41, 34, 8}
	fmt.Printf("Before insertion sort: %v\n", A)

	t0 = time.Now()

	for i := 0; i < len(A); i++ {
		var minIndx int = i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[minIndx] {
				minIndx = j
			}
		}
		A[i], A[minIndx] = A[minIndx], A[i]
	}

	duration = time.Since(t0).Nanoseconds()

	fmt.Printf("After insertion sort: %v\n", A)
	fmt.Printf("Duration: %v\n", duration)
}
