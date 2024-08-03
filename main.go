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

	heroes, err := db.GetHeroes()

	if err != nil {
		fmt.Println(err)
	}

	cachedMatches, _ := cache.LoadCachedMatches()

	errc := 0
	for _, mId := range cachedMatches {
		if !db.MatchExist(mId) {
			fmt.Println("Parse", mId)
			match, err := parser.ParseMatch(mId, &heroes)

			if err != nil {
				fmt.Println(err)
				if errc > 5 {
					break
				}
				errc += 1
			} else {
				fmt.Println("Insert", mId)
				db.InsertMatch(match, true)
			}
		} else {
			fmt.Println(mId, "already parsed")
		}
	}

	cache.CacheMatches(cachedMatches)

	db.Close()
}
