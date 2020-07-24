import pandas as pd 
from elastic import ElasticService

# read file csv
def read_csv(filename):
    df = pd.read_csv(filename)
    return df.to_dict('records')

#read file xlsx
def read_xlsx(filename):
    df = pd.read_excel(filename, dtype={'data': str, 'escrv': int, 'material': int, 'grp.merc.': int, 'qnt.faturd': int})  
    return df.to_dict('records')

#read file html
def read_html(filename):
    df = pd.read_html(filename)
    df = df[0]
    return df.to_dict('records')

if __name__ == '__main__':

    es = ElasticService()

    print("Insert Documents CSV")
    csv = read_csv('./raw_data/VENDAS_20190519.csv')
    for index in csv:
        es.on_insert_data(index)

    print("Insert Documents XLSX")
    xlsx = read_xlsx('./raw_data/VENDAS_20190520_20190522.xlsx')
    for index in xlsx:
        es.on_insert_data(index)

    print("Insert Documents HTML")
    html = read_html('./raw_data/VENDAS_20200524_20200525.html')
    for index in html:
        es.on_insert_data(index)