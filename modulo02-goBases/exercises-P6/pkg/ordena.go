package main

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

const (
	A = 100
	B = 1000
	C = 10000
)

func ordenaUm(num int) time.Duration {
	start := time.Now()
	variavel := rand.Perm(num)
	sort.Ints(variavel)
	time.Sleep(10 * time.Millisecond)
	elapsed := time.Since(start)
	log.Printf("\nBubble Sort : %d", elapsed)
	return elapsed
}

func ordenaDois(num int) time.Duration {
	start := time.Now()
	variavel := rand.Perm(num)

	sort.Slice(variavel, func(i, j int) bool {
		return variavel[i] > variavel[j]
	})

	elapsed := time.Since(start)
	log.Printf("\nBubble Sort decrescent:  %d", elapsed)
	return elapsed

}
