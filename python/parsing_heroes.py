from bs4 import BeautifulSoup
import requests
import pandas as pd
from tools import headers
from heroes_data import HeroesData

# Parsing heroes from Dotabuff
url = "https://dotabuff.com/heroes"
resp = requests.get(url, headers=headers)
soup = BeautifulSoup(resp.text, "lxml")

heroes_a = soup.select("table tbody tr div.tw-flex a")

heroes = []

for i in range(len(heroes_a)):
    heroes.append(heroes_a[i]['href'].split('/')[-1])

number_of_heroes = len(heroes)
keys = range(1, number_of_heroes + 1)

data = {
    "Key": keys,
    "Name": heroes,
    "pos1_value": [0] * number_of_heroes,
    "pos2_value": [0] * number_of_heroes,
    "pos3_value": [0] * number_of_heroes,
    "pos4_value": [0] * number_of_heroes,
    "pos5_value": [0] * number_of_heroes,
}

HeroesData = HeroesData(data)

# Parcing stats from Protracker (doesn't work)
# https://dota2protracker.com/_get/meta/pos-4/html

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

print(tables)

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

for name, stats in heroes.items():
    key = HeroesData.searchByName(name)
    summ = sum(stats)
    for i in range(5):
        stats[i] = stats[i] / summ
        HeroesData.df.at[key, f"pos{i + 1}_value"] = stats[i]

HeroesData.save()