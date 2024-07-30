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

	fmt.Println(teams)

	for _, team := range teams {
		fmt.Printf("Parsing %s", team.Name)
		matches, _ := team.ParseMatches()
		for _, matchId := range matches {
			fmt.Printf("Parsing match %d", matchId)
			match, _ := types.ParseMatch(matchId, heroes)
			db.InsertMatch(match, true)
		}
	}

	db.Close()
}
