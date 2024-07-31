package main

import (
	"dotaparser/packages/dotabase"
	"dotaparser/packages/types"
	"fmt"
	"time"
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
		matches, err := team.ParseMatches()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(team.Id, matches)

		time.Sleep(time.Second)
		for _, matchId := range matches {
			if !db.MatchExist(matchId) {
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

				time.Sleep(time.Second)
			}
		}
	}

	db.Close()
}
