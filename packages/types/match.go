package types

import (
	"fmt"
)

type Match struct {
	Id      int
	Winner  Side
	Radiant Draft
	Dire    Draft
	MetaDif float64
}

func (match *Match) Log() {
	fmt.Printf("id: %d:\n\tWinner: %s\n", match.Id, match.Winner)
	fmt.Println("Radiant:")
	match.Radiant.Log()
	fmt.Println("Dire:")
	match.Dire.Log()
}
