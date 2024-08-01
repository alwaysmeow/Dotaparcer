package parser

import (
	"dotaparser/packages/request"
	"dotaparser/packages/types"
	"errors"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseMatch(id int, heroes *types.Heroes) (*types.Match, error) {
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
	winner := types.None
	if strings.Contains(winnerClass, "radiant") {
		winner = types.Radiant
	} else if strings.Contains(winnerClass, "dire") {
		winner = types.Dire
	}

	var matchHeroes [10]*types.Hero
	heroesCounter := 0

	doc.Find("table.match-team-table td.cell-fill-image a").Each(func(i int, tag *goquery.Selection) {
		name := tag.AttrOr("href", "")
		name = name[strings.LastIndex(name, "/")+1:]
		var hero *types.Hero
		for _, h := range *heroes {
			if name == h.FormatName {
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

	match := types.Match{
		Id:      id,
		Winner:  winner,
		Radiant: types.CreateDraft([5]*types.Hero(matchHeroes[:5])),
		Dire:    types.CreateDraft([5]*types.Hero(matchHeroes[5:])),
	}

	match.MetaDif = match.Radiant.Meta - match.Dire.Meta

	return &match, nil
}
