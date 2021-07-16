from time import time
import matplotlib.pyplot as plt
import pandas as pd
import psycopg2
from sklearn.feature_extraction.text import TfidfVectorizer, CountVectorizer
from sklearn.decomposition import NMF, LatentDirichletAllocation
from datetime import datetime as dt
import datetime

conn = psycopg2.connect(user="admin",
                        password="admin@123",
                        host="localhost",
                        port=5432,
                        database="newsbot")



n_features=50
n_components=10
n_top_words=5

dtt=dt.now()
dtt=dtt.strftime("%d-%m-%Y")

def plot_top_words(model,feature_names,n_top_words, title):
    fig,axes=plt.subplots(2,5,figsize=(90,45), sharex=True)
    axes=axes.flatten()
    for topic_idx, topic in enumerate(model.components_):
        top_features_ind=topic.argsort()[:-n_top_words -1 :1]
        top_features=[feature_names[i] for i in top_features_ind]
        weights=topic[top_features_ind]

        ax=axes[topic_idx]
        ax.barh(top_features,weights,height=0.7)
        ax.set_title(f"Topic {topic_idx +1}", fontdict={"fontsize":30})
        ax.invert_yaxis()
        ax.tick_params(axis="both", which="major", labelsize=20)
        for i in "top right left".split():
            ax.spines[i].set_visible(False)
        fig.suptitle(title, fontsize=40)
    plt.subplots_adjust(top=0.9,bottom=0.05, wspace=0.9, hspace=0.3)
    plt.savefig(f"Plots/Date/{dtt}_topics.png")

def getData():
    dtt=dt.now()
    dtt=dtt.date()
    dtt = dtt-datetime.timedelta(days=30)
    dtt=pd.Timestamp(dtt)
    sql=("SELECT * FROM news")
    df=pd.read_sql(sql, conn,parse_dates=['date'])
    df.fillna("NA",inplace=True)
    # mask = (df['date'] == dtt)
    # df= df.loc[mask]
    return df


def bar_chart():
    dtt=dt.now()
    dtt=dtt.date()
    dtt=dtt.strftime("%d-%m-%Y")
    df=getData()
    data=[ d for d in df['title']]
    # print(data)
    tfidf_vectorizer=TfidfVectorizer(max_df=0.95, min_df=2,
                                    max_features=n_features,
                                    stop_words='english')
    tfidf=tfidf_vectorizer.fit_transform(data)
    nmf=NMF(n_components=n_components, random_state=1,alpha=.1,l1_ratio=.5).fit(tfidf)
    tfidf_feature_names=tfidf_vectorizer.get_feature_names()
    plot_top_words(nmf, tfidf_feature_names,n_top_words, f'Topics for {dtt}')                                    


bar_chart()