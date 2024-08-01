package types

import (
	"encoding/json"
	"io"
	"math"
	"os"
)

type Heroes map[int]Hero

func (heroes *Heroes) Init() {
	file, _ := os.Open("dotaconstants/build/heroes.json")
	byteValue, _ := io.ReadAll(file)
	file.Close()

	_ = json.Unmarshal(byteValue, &heroes)

	for key, hero := range *heroes {
		hero.AddFormatName()
		(*heroes)[key] = hero
	}
}

func (heroes *Heroes) Find(name string) (int, bool) {
	for key, hero := range *heroes {
		if hero.Name == name {
			return key, true
		}
	}
	return 0, false
}

func (heroes *Heroes) CalcMeta() {
	maxMatches := [5]int{0, 0, 0, 0, 0}

	for _, hero := range *heroes {
		for pos := 0; pos < 5; pos++ {
			if hero.Matches[pos] > maxMatches[pos] {
				maxMatches[pos] = hero.Matches[pos]
			}
		}
	}

	for key, hero := range *heroes {
		for pos := 0; pos < 5; pos++ {
			hero.Meta[pos] = hero.Winrate[pos] * math.Sqrt(float64(hero.Matches[pos])/float64(maxMatches[pos]))
		}
		(*heroes)[key] = hero
	}
}
