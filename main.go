package main

import (
	"dotaparser/packages/types"
)

func main() {
	heroes, _ := types.ParseHeroes()

	/*
		if err != nil {
			fmt.Println(err)
		}

		db := dotabase.GetDB()
		dotabase.DBinit(db)
		dotabase.InsertHeroes(db, *heroes)
		db.Close()
	*/

	match, _ := types.ParseMatch(7758350202, heroes)
	match.Log()
}
