package main

import (
	"dotaparser/types"
	"os"
)

func main() {
	heroes, err := types.ParseHeroes()

	if err != nil {
		os.Exit(1)
	}

	match, _ := types.ParseMatch(7853986291, heroes)
	match.Log()
}
