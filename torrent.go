package main

import (
	"log"
	"os"
)

type Torrent struct {
	infoHash []byte
	quit     chan struct{}
}

func NewTorrent(filename string, quit chan struct{}) (*Torrent, error) {
	torrent := &Torrent{quit: quit}

	log.Printf("Reading Torrent File")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer log.Printf("Finished Reading Torrent File")
	defer file.Close()

	return torrent, err
}

func init() {
}
