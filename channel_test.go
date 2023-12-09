package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	//pakek anonimous func
	go func() {
		time.Sleep(2 * time.Second)
		//ini cara mengirim data ke channel
		channel <- "Eko Kurniawan"
		fmt.Println("Mengirim data ke channel")
	}()

	//channel <- "eko"

	// ini cara menerima nilai dari channel
	data := <-channel

	fmt.Println(data)

	//fmt.Println(<-channel)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	//ini cara mengirim data ke channel
	channel <- "Eko Kurniawan"
	fmt.Println("Mengirim data ke channel")
}

func TestChannnelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)

	//data := <-channel
	//fmt.Println(data)
	fmt.Println(<-channel)
	time.Sleep(3 * time.Second)
}
