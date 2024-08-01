package parser

import (
	"dotaparser/packages/request"
	"dotaparser/packages/types"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseTeamMatches(team *types.Team, page int) ([]int, error) {
	url := request.DotabuffUrl(fmt.Sprintf("/esports/teams/%s/matches?pages=%d", team.Id, page))

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
