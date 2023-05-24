import pandas as pd

def herosIndexesDataFrame():
    df = pd.read_csv("dota2_heroes_indexes.csv", index_col="Key")
    return df