package main

import (
	"fmt"
	"time"
)

// Concurrency is the ability of a program to handle multiple tasks seemingly at the same time. Go has strong built-in support for concurrency, primarily through Goroutines and Channels.

// Goroutines (go keyword):
// A goroutine is a lightweight thread managed by the Go runtime. It's much cheaper than an OS thread.
// You create a goroutine by simply putting the go keyword before any function call.
// When you call go function(), the function starts running concurrently (in the background), and the main program execution continues without waiting for that function to finish.

// A simple function to simulate work
func printNumbers(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(500 * time.Millisecond) // Sleep for 500 milliseconds
	}
}
