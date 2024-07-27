package main

import (
	"dotaparser/packages/types"
	"fmt"
)

func main() {
	/*
		//heroes, _ := types.ParseHeroes()
		db := dotabase.GetDB()
		dotabase.DBinit(db)
		//insertHero(db, (*heroes)[0])
		hero, _ := dotabase.GetHero(db, 1)
		hero.Log()
		db.Close()
	*/
	h, _ := types.ParseHeroes()
	fmt.Println(h)
}
