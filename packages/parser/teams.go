package parser

import (
	"dotaparser/packages/request"
	"dotaparser/packages/types"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseTeams() ([]types.Team, error) {
	url := request.DotabuffUrl("/esports/teams")

	resp, err := request.Request(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(resp.Body)

	var teams []types.Team

	doc.Find("table.r-tab-enabled tbody tr a").Each(func(i int, tag *goquery.Selection) {
		name := tag.Text()
		if name != "" {
			href := tag.AttrOr("href", "")
			href = href[strings.LastIndex(href, "/")+1:]
			team := types.Team{Id: href, Name: name}
			teams = append(teams, team)
		}
	})

	return teams, nil
}
