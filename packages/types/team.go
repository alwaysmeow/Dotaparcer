package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"dotaparser/packages/request"
)

type Team struct {
	Id   string
	Name string
}

func ParseTeams() ([]Team, error) {
	url := request.DotabuffUrl("/esports/teams")

	resp, err := request.Request(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	var teams []Team

	doc.Find("table.r-tab-enabled tbody tr a").Each(func(i int, tag *goquery.Selection) {
		name := tag.Text()
		if name != "" {
			href := tag.AttrOr("href", "")
			href = href[strings.LastIndex(href, "/")+1:]
			team := Team{Id: href, Name: name}
			teams = append(teams, team)
		}
	})

	return teams, nil
}

func (team *Team) ParseMatches(page int) ([]int, error) {
	url := request.DotabuffUrl(fmt.Sprintf("/esports/teams/%s/matches", team.Id))

	resp, err := request.Request(url)
	if err != nil {
		return nil, err
	}

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	var matches []int

	doc.Find("table.recent-esports-matches tbody tr div a").Each(func(i int, tag *goquery.Selection) {
		class := tag.AttrOr("class", "")
		if class == "lost" || class == "won" {
			matchStr := tag.AttrOr("href", "")
			matchStr = matchStr[strings.LastIndex(matchStr, "/")+1:]
			matchId, _ := strconv.Atoi(matchStr)
			matches = append(matches, matchId)
		}
	})

	return matches, nil
}
