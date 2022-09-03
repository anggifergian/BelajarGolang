package goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func GiveResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Anggi Fergian"
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Anggi Fergian"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)

	channel <- "Anggi"
	channel <- "Fergian"
	channel <- "Pratama"

	fmt.Println(<-channel)

	fmt.Println("Selesai...")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulanangan ke " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println("menerima data", data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveResponse(channel1)
	go GiveResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2: ", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
		}

		if counter == 2 {
			break
		}
	}
}
