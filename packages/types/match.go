package types

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"dotaparser/packages/request"
)

type Match struct {
	Id      int
	Winner  Side
	Radiant Draft
	Dire    Draft
	MetaDif float64
}

func (match *Match) Log() {
	fmt.Printf("id: %d:\n\tWinner: %s\n", match.Id, match.Winner)
	fmt.Println("Radiant:")
	match.Radiant.Log()
	fmt.Println("Dire:")
	match.Dire.Log()
}

func ParseMatch(id int, heroes *Heroes) (*Match, error) {
	url := request.DotabuffUrl(fmt.Sprintf("/matches/%d", id))

	resp, err := request.Request(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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

	var matchHeroes [10]*Hero
	heroesCounter := 0

	doc.Find("table.match-team-table td.cell-fill-image a").Each(func(i int, tag *goquery.Selection) {
		name := tag.AttrOr("href", "")
		name = name[strings.LastIndex(name, "/")+1:]
		var hero *Hero
		for _, h := range *heroes {
			heroname := h.Name
			heroname = strings.ReplaceAll(heroname, " ", "-")
			heroname = strings.ReplaceAll(heroname, "'", "")
			heroname = strings.ToLower(heroname)
			if name == heroname {
				hero = &h
				break
			}
		}
		if hero != nil {
			matchHeroes[i] = hero
			heroesCounter += 1
		} else {
			fmt.Printf("Can't find hero: %s\n", name)
		}
	})

	if heroesCounter != 10 {
		return nil, errors.New("сan't parse heroes")
	}

	match := Match{
		Id:      id,
		Winner:  winner,
		Radiant: CreateDraft([5]*Hero(matchHeroes[:5])),
		Dire:    CreateDraft([5]*Hero(matchHeroes[5:])),
	}

	match.MetaDif = match.Radiant.Meta - match.Dire.Meta

	return &match, nil
}
