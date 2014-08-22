package main

import (
	"log"
	"os"
)

func main() {
	log.Printf("Main Started")
	defer log.Printf("Main Ended")
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s: <torrent file>\n", os.Args[0])
	}
	quit := make(chan struct{})

	_, err := NewTorrent(os.Args[1], quit)
	if err != nil {
		log.Fatal(err)
	}

}
