package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func sleep1(d time.Duration) {
	now := time.Now()
	future := now.Add(d)

	for {
		if time.Now().After(future) {
			break
		}
	}
}

func sleep2(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Program is working...")

	sleep1(3 * time.Second)
	fmt.Println("Message 3 seconds later...")

	sleep2(3 * time.Second)
	fmt.Println("Message 3 seconds later...")
}
