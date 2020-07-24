import pandas as pd 
from elastic import insert_data

def read_csv(filename):
    df = pd.read_csv(filename)
    return df.to_dict('records')

def read_xlsx(filename):
    df = pd.read_excel(filename, dtype={'data': str, 'escrv': int, 'material': int, 'grp.merc.': int, 'qnt.faturd': int})  
    return df.to_dict('records')

def read_html(filename):
    df = pd.read_html(filename)
    df = df[0]
    return df.to_dict('records')

if __name__ == '__main__':
    
    csv = read_csv('./raw_data/VENDAS_20190519.csv')
    for index in csv:
        insert_data(index) 

    xlsx = read_xlsx('./raw_data/VENDAS_20190520_20190522.xlsx')
    for index in xlsx:
        insert_data(index)

    html = read_html('./raw_data/VENDAS_20200524_20200525.html')
    for index in html:
        insert_data(index)
    