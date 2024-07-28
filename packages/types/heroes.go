package types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Heroes map[int]Hero

func (heroes *Heroes) Init() {
	file, _ := os.Open("dotaconstants/build/heroes.json")
	byteValue, _ := io.ReadAll(file)
	file.Close()

	_ = json.Unmarshal(byteValue, &heroes)
}

func (heroes *Heroes) find(name string) (int, bool) {
	for key, hero := range *heroes {
		if hero.Name == name {
			return key, true
		}
	}
	return 0, false
}

func ParseHeroes() (*Heroes, error) {
	var heroes Heroes
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

			key, found := heroes.find(name)
			if !found {
				fmt.Printf("%s not found\n", name)
			} else {
				hero := heroes[key]
				hero.Matches[pos] = int(matches)
				hero.Winrate[pos] = winrate
				heroes[key] = hero
			}
		})
	}

	// id: https://dota2protracker.com/_get/search

	return &heroes, nil
}
