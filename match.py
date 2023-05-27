from bs4 import BeautifulSoup
import requests
from tools import headers

class Team:
    def __init__(self, team_html):
        self.isDefined = True
        heroes_links = team_html.select(".match-team-table .image-container-hero a")
        heroes = []
        for link in heroes_links:
            heroes.append(link['href'].split('/')[-1])

        self.heroes = [None, None, None, None, None]

        role_tags = team_html.select(".tf-fa .role-icon")
        lane_tags = team_html.select(".tf-fa .lane-icon")
        if len(role_tags) != 5 or len(lane_tags) != 5:
            self.isDefined = False
        else:
            for i in range(5):
                role = role_tags[i]['class'][-1]
                lane = lane_tags[i]['class'][-1]
                if role == "core-icon":
                    if lane == "safelane-icon":
                        self.heroes[0] = heroes[i]
                    elif lane == "midlane-icon":
                        self.heroes[1] = heroes[i]
                    elif lane == "offlane-icon":
                        self.heroes[2] = heroes[i]
                    else:
                        self.heroes[3] = heroes[i]
                else:
                    if lane == "safelane-icon":
                        self.heroes[4] = heroes[i]
                    else:
                        self.heroes[3] = heroes[i]
        if self.heroes[0] is None or self.heroes[1] is None or self.heroes[2] is None or self.heroes[3] is None or self.heroes[4] is None:
            self.isDefined = False

    def consoleOutput(self):
        print("Carry:", self.heroes[0])
        print("Mider:", self.heroes[1])
        print("Offlaner:", self.heroes[2])
        print("Support:", self.heroes[3])
        print("Hard Support:", self.heroes[4])

class Match:
    def __init__(self, id):
        self.id = id

        url = "https://dotabuff.com/matches/" + str(id)
        resp = requests.get(url, headers=headers)
        soup = BeautifulSoup(resp.text, "lxml")

        self.winner = soup.select_one(".match-result")['class'][-1].capitalize()

        workspace = soup.select_one(".team-results")
        self.radiant = Team(workspace.select_one(".radiant"))
        self.dire = Team(workspace.select_one(".dire"))
        self.isDefined = self.radiant.isDefined and self.dire.isDefined
        self.consoleOutput()

    def consoleOutput(self):
        print("====================")
        print("Match ID:", self.id)
        print("\n====Team Radiant====")
        self.radiant.consoleOutput()
        print("\n====Team Dire====")
        self.dire.consoleOutput()
        print("\nWinner:", self.winner)
        print("====================")

# test
Match(7154793285)