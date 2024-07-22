package main

import (
	"dotaparser/types"
	"fmt"
)

func main() {
	teams, _ := types.ParseTeamMatches(types.Team{Id: "7119388-team-spirit"})
	fmt.Print(teams)
}
