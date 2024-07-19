package main

import (
	"fmt"
	"log"

	"dotaparser/types"
)

func main() {
	match, err := types.ParseMatch(7154793285)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Match %d: %s won!", match.Id, match.Winner)
}
