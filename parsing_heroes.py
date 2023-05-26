from bs4 import BeautifulSoup
import requests
import pandas as pd

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36'
    }

url = "https://dotabuff.com/heroes"
resp = requests.get(url, headers=headers)
soup = BeautifulSoup(resp.text, "lxml")

heroes_a = soup.select(".hero-grid a")

heroes = {}

for i in range(len(heroes_a)):
    heroes[i + 1] = heroes_a[i]['href'].split('/')[-1]

df = pd.DataFrame.from_dict(heroes, orient='index', columns=['Name'])

df.to_csv('dota2_heroes_indexes.csv', index_label='Key')