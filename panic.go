package main

import (
	"os"
	"time"
)

func main() {

	go func() {
		panic("a problem")
	}()

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second)
}
