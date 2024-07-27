package types

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Team struct {
	Id   string
	Name string
}

func ParseTeams() ([]Team, error) {
	url := "https://ru.dotabuff.com/esports/teams"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

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

func ParseTeamMatches(team Team) ([]int, error) {
	url := fmt.Sprintf("https://ru.dotabuff.com/esports/teams/%s/matches", team.Id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
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
