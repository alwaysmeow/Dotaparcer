import pandas as pd
import numpy as np
from datetime import date

class HeroesData:
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
    
    def getPositionValues(self, key):
        return list(self.df.loc[key])[1:]
    
    def getTeamPositionValues(self, keys):
        if len(keys) != 5:
            print("There is not five heroes")
            return None
        else:
            values = [
                self.getPositionValues(keys[0]),
                self.getPositionValues(keys[1]),
                self.getPositionValues(keys[2]),
                self.getPositionValues(keys[3]),
                self.getPositionValues(keys[4]),
            ]
            return np.array(values)
    def save(self):
        today = date.today().strftime("%m.%Y")
        self.df.to_csv(f'heroes_stats/dota2_heroes_{today}.csv', index_label='Key')
        self.df.to_csv(f'dota2_heroes_actual.csv', index_label='Key')