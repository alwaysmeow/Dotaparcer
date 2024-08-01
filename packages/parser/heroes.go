package parser

import (
	"dotaparser/packages/types"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHeroes() (*types.Heroes, error) {
	var heroes types.Heroes
	heroes.Init()

	for pos := 0; pos < 5; pos++ {
		url := fmt.Sprintf("https://dota2protracker.com/_get/meta/pos-%d/html", pos+1)

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

		doc.Find("div.tbody div.grid").Each(func(index int, tag *goquery.Selection) {
			name := tag.AttrOr("data-hero", "")
			name = strings.TrimSpace(name)

			str_matches := tag.AttrOr("data-matches", "")
			matches, err := strconv.ParseInt(str_matches, 10, 64)
			if err != nil {
				matches = 0
			}

			str_winrate := tag.AttrOr("data-wr", "")
			winrate, err := strconv.ParseFloat(str_winrate, 64)
			if err != nil {
				winrate = 0
			}

			key, found := heroes.Find(name)
			if !found {
				fmt.Printf("%s not found\n", name)
			} else {
				hero := heroes[key]
				hero.Winrate[pos] = winrate / 100
				hero.Matches[pos] = int(matches)
				heroes[key] = hero
			}
		})
	}

	heroes.CalcMeta()

	// id: https://dota2protracker.com/_get/search

	return &heroes, nil
}
