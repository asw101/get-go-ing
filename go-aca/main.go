package main

import (
	"log"
	"time"
)

func main() {
	i := 0
	seconds := 10
	for {
		log.Printf("%ds\n", i)
		time.Sleep(time.Duration(seconds) * time.Second)
		i = i + seconds
	}
}
