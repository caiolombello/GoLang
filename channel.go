package main

import (
	"fmt"
	"math/rand"
	"time"
)

func webServerWorker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("WorkerId:", workerId, "Mensagem processada:", res)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {
	msg := make(chan int)

	for i := 1; i <= 3; i++ {
		go webServerWorker(i, msg) // Roda em outra thread
	}

	for i := 0; i < 10; i++ {
		msg <- i
	}
}
