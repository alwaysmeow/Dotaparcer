package main

import (
	"dotaparser/packages/dotabase"
	"dotaparser/packages/types"
	"fmt"
)

func main() {
	heroes, err := types.ParseHeroes()

	if err != nil {
		fmt.Println(err)
	}

	db := dotabase.GetDB()
	dotabase.DBinit(db)
	dotabase.InsertHeroes(db, *heroes)
	db.Close()
}
