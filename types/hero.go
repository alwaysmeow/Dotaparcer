package types

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Hero struct {
	Name    string
	Id      int
	Matches [5]int
	Winrate [5]float64
}

func (hero *Hero) log() {
	fmt.Printf("%s:\n", hero.Name)
	for i := 0; i < 5; i++ {
		fmt.Printf("\tPos %d: %d matches, %.2f winrate\n", i+1, hero.Matches[i], hero.Winrate[i])
	}
}

type Heroes []Hero

func (heroes *Heroes) find(name string) (*Hero, bool) {
	for i := range *heroes {
		if (*heroes)[i].Name == name {
			return &(*heroes)[i], true
		}
	}
	return nil, false
}

func (heroes *Heroes) append(hero Hero) *Hero {
	*heroes = append(*heroes, hero)
	return &(*heroes)[len(*heroes)-1]
}

func ParseHeroes() (*Heroes, error) {
	var heroes Heroes
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
			if pos == 0 {
				heroes = append(heroes, Hero{})
			}
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

			if pos == 0 {
				heroes[index].Name = name
				heroes[index].Matches[pos] = int(matches)
				heroes[index].Winrate[pos] = winrate
			} else {
				hero, found := heroes.find(name)
				if !found {
					hero = heroes.append(Hero{Name: name})
				}
				hero.Matches[pos] = int(matches)
				hero.Winrate[pos] = winrate
			}
		})
	}

	// id: https://dota2protracker.com/_get/search

	return &heroes, nil
}