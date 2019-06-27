package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HelloData struct {
	Greeting string
	Name     string
	Age      int
	Gender   string
	Happy    bool
}

func (h *HelloData) marshal() []byte {
	var m []byte
	m, err := json.Marshal(h)
	if err != nil {
		panic(err)
	}
	return m
}

var c = make(chan []byte)

func Say123Hello() {

	charlie := HelloData{
		Greeting: "Heya",
		Name:     "Charlie",
		Age:      19,
		Gender:   "Male",
		Happy:    false,
	}
	anna := HelloData{
		Greeting: "Hellooooo nurse",
		Name:     "Anna",
		Age:      25,
		Gender:   "Female",
		Happy:    true,
	}
	job := HelloData{
		Greeting: "Hi",
		Name:     "Job",
		Age:      34,
		Gender:   "Male",
		Happy:    true,
	}

	for {
		time.Sleep(1 * time.Second)
		c <- charlie.marshal()
		time.Sleep(1 * time.Second)
		c <- job.marshal()
		time.Sleep(1 * time.Second)
		c <- anna.marshal()
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("Caught sig: %v", sig)
		fmt.Println("Wait for 1 second to finish processing")
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()

	go Say123Hello()

	for {
		t0 := time.Now()
		hdata := <-c
		fmt.Printf(string(hdata) + "\n")
		t1 := time.Now()
		fmt.Println("Duration, ", t1.Sub(t0))
	}

}
