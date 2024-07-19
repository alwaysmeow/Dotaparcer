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
	Radiant [5]Hero
	Dire    [5]Hero
}

func ParseMatch(id int) (*Match, error) {
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

	return &Match{
		Id:     id,
		Winner: winner,
	}, nil
}
