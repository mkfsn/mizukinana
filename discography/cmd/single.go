//+build tools

package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/mkfsn/mizukinana/discography"
)

func main() {
	index := flag.Int("i", 32-1, "index of singles: 1 ~")
	flag.Parse()

	singles := discography.Singles
	if *index <= 0 || len(singles) < *index {
		fmt.Println("Error: no such album")
		return
	}

	bytes, err := json.MarshalIndent(singles[*index], "", "\t")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf(string(bytes))
}
