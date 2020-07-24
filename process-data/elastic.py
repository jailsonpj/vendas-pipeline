from elasticsearch import Elasticsearch
import logging

user = "elastic"
passwd = "YgcFDLqEuhvxafts9Ulv8KHJ"
es = Elasticsearch(["https://elastic:YgcFDLqEuhvxafts9Ulv8KHJ@c01160965fdb47918c3d06933a6df089.us-east-1.aws.found.io:9243"])

index = "vendas-analysis"

def insert_data(data):
    res=es.index(index=index, body=data,)
    
    if res == False:
        logging.info("[ELASTICSEARCH][INSERT][ERROR]")
        
    logging.info("[ELASTICSEARCH][INSERT][SUCESSUFULLY]")

        
