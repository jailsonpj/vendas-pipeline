from elasticsearch import Elasticsearch
from elasticsearch.helpers import scan

es = Elasticsearch([{'host': 'localhost', 'port':9200}])
index = "vendas-analysis"


def get_all():
    data = []

    match_all = {
        "query":{
            "match_all":{}
        }
    }
    resp = scan(es, index=index,query=match_all)
    for item in resp:
        #print(item['_source'])
        data.append(item['_source'])

    return data

def get_material(material):
    data = []

    match = {
        "query":{
            "match":{
                "Material": material
            }
        }
    }

    resp = scan(es, index=index,query=match)
    for item in resp:
        #print(item['_source'])
        data.append(item['_source'])

    return data

def get_escrv(escrv):
    data = []

    match = {
        "query":{
            "match":{
                "Escrv": escrv
            }
        }
    }

    resp = scan(es, index=index,query=match)
    for item in resp:
        #print(item['_source'])
        data.append(item['_source'])

    return data

def get_data(data):
    data = []

    match = {
        "query":{
            "match":{
                "Data": data
            }
        }
    }

    resp = scan(es, index=index,query=match)
    for item in resp:
        #print(item['_source'])
        data.append(item['_source'])

    return data

def get_grp_merc(grpmerc):
    data = []

    match = {
        "query":{
            "match":{
                "GrpMerc": data
            }
        }
    }

    resp = scan(es, index=index,query=match)
    for item in resp:
        #print(item['_source'])
        data.append(item['_source'])

    return data