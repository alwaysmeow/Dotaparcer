import pandas as pd
from datetime import date

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36'
    }

class Heroes:
    def __init__(self, data = None):
        if data is None:
            self.df = pd.read_csv("dota2_heroes_actual.csv", index_col="Key")
        else:
            self.df = pd.DataFrame(data)
            self.df.set_index("Key", inplace = True)
    
    def searchByName(self, name):
        name = name.replace(" ", "-").replace("'", "").lower()
        answer = self.df.index[self.df["Name"] == name]
        if len(answer) == 0:
            print(f"No matches. Uncorrect hero's name: {name}")
            return 0
        else:
            return answer[0]
    
    def save(self):
        today = date.today().strftime("%d.%m.%Y")
        self.df.to_csv(f'heros_stats/dota2_heroes_{today}.csv', index_label='Key')
        self.df.to_csv(f'dota2_heroes_actual.csv', index_label='Key')