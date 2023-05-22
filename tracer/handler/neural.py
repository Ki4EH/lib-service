import pandas as pd
from sklearn.linear_model import LinearRegression


def predict(learns):
    file = 'dataset.csv'
    data = pd.read_csv(file)
    print(data.head(10))
    data = data.astype('float32')
    print(data.head(10))
    reg = LinearRegression()
    reg.fit(data[['par1', 'par2', 'par3', 'par4', 'par5']],data.fin)
    all = reg.predict([learns])
    return all[0]