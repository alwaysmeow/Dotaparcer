import requests
from bs4 import BeautifulSoup
from tools import headers
from datetime import datetime, timedelta, timezone

url = "https://dotabuff.com/players"
resp = requests.get(url, headers=headers)
soup = BeautifulSoup(resp.text, "lxml")

items = soup.select(".content-inner .sortable tbody tr .cell-large")

current_time = datetime.now(timezone.utc)
# Время по UTC (Coordinated Universal Time) - международный стандарт времени, используемый для согласования времени по всему миру

for item in items:
    time = datetime.fromisoformat(item.select_one("time")["datetime"]).replace(tzinfo=timezone.utc) # Так же преобразование к нужному формату 
    time_difference = current_time - time
    if time_difference <= timedelta(days=7):
        a_tag = item.select_one(".link-type-player")
        name = a_tag.get_text()
        id = a_tag["href"].split("/")[-1]
        print(name, id, time_difference)