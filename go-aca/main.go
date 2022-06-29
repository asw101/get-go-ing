package main

import (
	"log"
	"os"
	"time"
)

func main() {
	prefix := ""
	if val := os.Getenv("HELLO"); val != "" {
		prefix = val + ": "
	}

	i := 0
	seconds := 10
	for {
		log.Printf("%s%ds\n", prefix, i)
		time.Sleep(time.Duration(seconds) * time.Second)
		i = i + seconds
	}
}
