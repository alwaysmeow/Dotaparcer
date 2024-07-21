package types

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Match struct {
	Id      int
	Winner  Side
	Radiant [5]*Hero
	Dire    [5]*Hero
}

func (match *Match) log() {
	fmt.Printf("id: %d:\n\tWinner: %s", match.Id, match.Winner)
	fmt.Println("Radiant:")
	for i := 0; i < 5; i++ {
		match.Radiant[i].log()
	}
	fmt.Println("Dire:")
	for i := 0; i < 5; i++ {
		match.Dire[i].log()
	}
}

func ParseMatch(id int, heroes *Heroes) (*Match, error) {
	url := fmt.Sprintf("https://dotabuff.com/matches/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	winnerClass := doc.Find(".match-result").AttrOr("class", "")
	winner := None
	if strings.Contains(winnerClass, "radiant") {
		winner = Radiant
	} else if strings.Contains(winnerClass, "dire") {
		winner = Dire
	}

	var teams [10]*Hero

	doc.Find("table.match-team-table td.cell-fill-image a").Each(func(i int, s *goquery.Selection) {
		name := s.AttrOr("href", "")
		name = name[strings.LastIndex(name, "/")+1:]
		found := false
		index := 0
		for !found && index < len(*heroes) {
			heroname := (*heroes)[index].Name
			heroname = strings.ReplaceAll(heroname, " ", "-")
			heroname = strings.ToLower(heroname)
			if name == heroname {
				found = true
			} else {
				index += 1
			}
		}
		if found {
			teams[i] = &(*heroes)[index]
		}
		if !found {
			fmt.Printf("Can't find hero: %s\n", name)
		}
	})

	match := Match{
		Id:      id,
		Winner:  winner,
		Radiant: [5]*Hero(teams[:5]),
		Dire:    [5]*Hero(teams[5:]),
	}

	match.log()

	return &match, nil
}
