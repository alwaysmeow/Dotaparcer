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

heroes = []

for i in range(len(heroes_a)):
    heroes.append(heroes_a[i]['href'].split('/')[-1])

keys = range(1, len(heroes) + 1)

data = {
    "Key": keys,
    "Name": heroes,
}

df = pd.DataFrame(data)
df.set_index("Key", inplace = True)

df.to_csv('dota2_heroes_indexes.csv', index_label='Key')