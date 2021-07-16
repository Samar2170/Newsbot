import psycopg2
import pandas as pd
import json

from datetime import datetime
from elasticsearch import Elasticsearch

es=Elasticsearch(['localhost'],port=9200)

conn = psycopg2.connect(user="admin",
                        password="admin@123",
                        host="localhost",
                        port=5432,
                        database="newsbot")



def indexFirst():
    es.indices.create(index='indian-news',ignore=400)

def add_data():
    df = pd.read_sql("SELECT * FROM news", conn)
    df.fillna("NA",inplace=True)
    for i,r in df.iterrows():
        news={} 
        news['title']=r['title']
        news['url']=r['url']
        news['source']=r['source']
        #news=json.load(news)
        es.index(index='indian-news',id=i, doc_type='Text', body=news)


def add_data_by_date():
    dt = datetime.now()
    dt=dt.date()
    dt=pd.Timestamp(dt)
    sql=("SELECT * FROM news")

    df = pd.read_sql(sql, conn, parse_dates=["date"])
    df.fillna("NA",inplace=True)
    mask = (df['date'] == dt)
    df= df.loc[mask]

    for i,r in df.iterrows():
        news={} 
        news['title']=r['title']
        news['url']=r['url']
        news['source']=r['source']
        news['date']=r['date']
        es.index(index='indian-news',id=i, doc_type='Text', body=news)
    


# indexFirst()

# add_data()

add_data_by_date()
