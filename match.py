from bs4 import BeautifulSoup
import requests

class Team:
    def __init__(self, team_html):
        self.isDefined = True
        heros_links = team_html.select(".match-team-table .image-container-hero a")
        heros = []
        for link in heros_links:
            heros.append(link['href'].split('/')[-1])

        self.carry = None
        self.mider = None
        self.offlaner = None
        self.support = None
        self.hardSupport = None

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
                        self.carry = heros[i]
                    elif lane == "midlane-icon":
                        self.mider = heros[i]
                    elif lane == "offlane-icon":
                        self.offlaner = heros[i]
                    else:
                        self.support = heros[i]
                else:
                    if lane == "safelane-icon":
                        self.hardSupport = heros[i]
                    else:
                        self.support = heros[i]
        if self.carry is None or self.mider is None or self.offlaner is None or self.support is None or self.hardSupport is None:
            self.isDefined = False

    def consoleOutput(self):
        print("Carry:", self.carry)
        print("Mider:", self.mider)
        print("Offlaner:", self.offlaner)
        print("Support:", self.support)
        print("Hard Support:", self.hardSupport)

class Match:
    def __init__(self, id):
        self.id = id

        headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36'
        }

        url = "https://ru.dotabuff.com/matches/" + str(id)
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