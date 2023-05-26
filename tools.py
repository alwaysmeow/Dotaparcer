import pandas as pd

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36'
    }

class Heroes:
    def __init__(self):
        self.df = pd.read_csv("dota2_heroes.csv", index_col="Key")
    
    def searchByName(self, name):
        name = name.replace(" ", "-").replace("'", "").lower()
        answer = self.df.index[self.df["Name"] == name]
        if len(answer) == 0:
            print(f"No matches. Uncorrect hero's name: {name}")
            return 0
        else:
            return answer[0]
    
    def save(self):
        self.df.to_csv('dota2_heroes.csv', index_label='Key')

h = Heroes()
h.searchByName("batrider")