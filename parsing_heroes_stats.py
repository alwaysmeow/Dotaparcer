import requests
from bs4 import BeautifulSoup
from tools import *

url = "https://www.dota2protracker.com/"

resp = requests.get(url, headers=headers)
soup = BeautifulSoup(resp.text, "lxml")

tables = {
    1: soup.select(".top-hero-table .tabs-2 tbody tr"),
    2: soup.select(".top-hero-table .tabs-3 tbody tr"),
    3: soup.select(".top-hero-table .tabs-4 tbody tr"),
    4: soup.select(".top-hero-table .tabs-5 tbody tr"),
    5: soup.select(".top-hero-table .tabs-6 tbody tr"),
}

heroes = {}

for i in [1, 2, 3, 4, 5]:
    for item in tables[i]:
        name = item.select_one("a")["title"]
        matches = int(item.select_one(".td-matches .perc-wr").text)
        if name in heroes:
            heroes[name][i - 1] = matches
        else:
            heroes[name] = [0, 0, 0, 0, 0]
            heroes[name][i - 1] = matches

HeroesData = Heroes()
HeroesData.df["pos1_value"] = [0] * 124
HeroesData.df["pos2_value"] = [0] * 124
HeroesData.df["pos3_value"] = [0] * 124
HeroesData.df["pos4_value"] = [0] * 124
HeroesData.df["pos5_value"] = [0] * 124

for name, stats in heroes.items():
    key = HeroesData.searchByName(name)
    summ = sum(stats)
    for i in range(5):
        stats[i] = stats[i] / summ
        HeroesData.df.at[key, f"pos{i + 1}_value"] = stats[i]

HeroesData.save()