package main

import (
	"dotaparser/packages/dotabase"
	"dotaparser/packages/types"
	"fmt"
)

func main() {
	heroes, _ := types.ParseHeroes()

	db := dotabase.GetDB()
	db.Init()

	teams, err := types.ParseTeams()

	if err != nil {
		fmt.Println(err)
	}

	for _, team := range teams {
		fmt.Printf("Parsing %s\n", team.Name)
		matches, _ := team.ParseMatches()
		for _, matchId := range matches {
			fmt.Printf("Parsing match %d\n", matchId)
			match, err := types.ParseMatch(matchId, heroes)

			if err != nil {
				fmt.Println(err)
			} else {
				err = db.InsertMatch(match, true)

				fmt.Printf("Insert match %d\n", matchId)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	db.Close()
}
