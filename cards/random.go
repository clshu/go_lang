package main

import (
	"math/rand"
	"time"
)

// Create random number between [0, n)
// Generate a new seed every time

func random(n int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(n)
}
