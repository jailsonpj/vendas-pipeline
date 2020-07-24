from elasticsearch import Elasticsearch
import logging
import time


index = "vendas-analysis"


class ElasticService:
    def __init__(self):
        self.client = Elasticsearch("https://elastic:YgcFDLqEuhvxafts9Ulv8KHJ@c01160965fdb47918c3d06933a6df089.us-east-1.aws.found.io:9243")
        self.index = index
        self.doc_type = '_doc'

    def on_check_index(self):
        '''
        Ckeck index exists
        '''
        if not self.client.indices.exists(index=self.index):
            logging.info("[ELASTICSEARCH][INDEX][NO-EXIST]")
            self.client.indices.create(index=self.index, ignore=400)
            time.sleep(2)

            logging.info("[ELASTICSEARCH][INDEX][CREATE][SUCESSUFULLY]")
        else:
            logging.info("[ELASTICSEARCH][INDEX][EXISTS]")

   
    def on_insert_data(self, data):
        '''
        Insert document in elasticsearch
        '''
        self.on_check_index()
        
        res =  self.client.index(index=self.index, doc_type='_doc', body=data,)
        if res == False:
            logging.info("[ELASTICSEARCH][INSERT][ERROR]")
            return False

        logging.info("[ELASTICSEARCH][INSERT][SUCESSUFULLY]")
        return True


        
