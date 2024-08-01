package main

import (
	"dotaparser/packages/cache"
	"dotaparser/packages/dotabase"
	"dotaparser/packages/parser"
	"fmt"
)

func main() {

	db := dotabase.GetDB()
	db.Init()

	_, err := db.GetHeroes()

	if err != nil {
		fmt.Println(err)
	}

	teams, err := parser.ParseTeams()

	if err != nil {
		fmt.Println(err)
	}

	cachedMatches, _ := cache.LoadCachedMatches()

	for _, team := range teams {
		fmt.Println(team.Name)
		matches, err := parser.ParseTeamMatches(&team, 4)

		if err != nil {
			fmt.Println(err)
		}

		cachedMatches = append(cachedMatches, matches...)
	}

	cache.CacheMatches(cachedMatches)

	db.Close()
}
