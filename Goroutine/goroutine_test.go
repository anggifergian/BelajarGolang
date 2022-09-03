package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestRunHelloWorld(t *testing.T) {
	go RunHelloWorld()

	fmt.Println("[Finished]")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display...", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100_000; i++ {
		// Running concurrent
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}