import datetime 
from datetime import datetime as dt
import pandas as pd
import psycopg2
import matplotlib as mpl
from wordcloud import WordCloud, STOPWORDS
import matplotlib.pyplot as plt

conn = psycopg2.connect(user="admin",
                        password="admin@123",
                        host="localhost",
                        port=5432,
                        database="newsbot")





def getData():
    dtt=dt.now()
    dtt=dtt.date()
    # dtt = dtt-datetime.timedelta(days=30)
    dtt=pd.Timestamp(dtt)
    sql=("SELECT * FROM news")
    df=pd.read_sql(sql, conn,parse_dates=['date'])
    df.fillna("NA",inplace=True)
    mask = (df['source'] == "ET")
    df= df.loc[mask]
    print(df)

    return df,dtt

def get_wordCloud():
    df,dtt=getData()
    mpl.rcParams['figure.figsize']=(12.0,12.0)
    mpl.rcParams['font.size']=12
    mpl.rcParams['savefig.dpi']=100
    mpl.rcParams['figure.subplot.bottom']=.1
    stopwords=set(STOPWORDS)
    wordcloud=WordCloud(background_color='white',stopwords=stopwords,
                        max_words=500, max_font_size=42, random_state=42,
                        ).generate(str(df['title']))
    print(wordcloud)
    fig=plt.figure(1)
    plt.imshow(wordcloud)
    plt.axis('off')
    plt.savefig(f"Plots/Date/wc_NC_{dtt}.png")

get_wordCloud()