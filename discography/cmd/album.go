package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/mkfsn/mizukinana/discography"
)

func main() {
	index := flag.Int("i", 3, "index of albums: 1 ~")
	flag.Parse()

	albums := discography.Albums
	if *index <= 0 || len(albums) < *index {
		fmt.Println("Error: no such album")
		return
	}

	bytes, err := json.MarshalIndent(albums[*index-1], "", "\t")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf(string(bytes))
}
