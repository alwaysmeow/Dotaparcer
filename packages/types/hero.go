package types

import (
	"fmt"
)

type Hero struct {
	Name    string `json:"localized_name"`
	Id      int    `json:"id"`
	Matches [5]int
	Winrate [5]float64
	Meta    [5]float64
}

func (hero *Hero) Log() {
	fmt.Printf("%s:\n", hero.Name)
	for i := 0; i < 5; i++ {
		fmt.Printf("\tPos %d: %d matches, %.2f winrate, %.4f meta\n", i+1, hero.Matches[i], hero.Winrate[i], hero.Meta[i])
	}
}
