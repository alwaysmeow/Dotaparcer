from bs4 import BeautifulSoup
import requests

class Team:
    def __init__(self, team_html):
        heros_links = team_html.select(".match-team-table .image-container-hero a")
        heros = []
        for link in heros_links:
            heros.append(link['href'].split('/')[-1])
        print(heros)

class Match:
    def __init__(self, id):
        self.id = id

        headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36'
        }

        url = "https://ru.dotabuff.com/matches/" + str(id)
        resp = requests.get(url, headers=headers)
        soup = BeautifulSoup(resp.text, "lxml")

        workspace = soup.select_one(".team-results")
        self.radiant = Team(workspace.select_one(".radiant"))
        self.dire = Team(workspace.select_one(".dire"))

# test
Match(7122507163)