from django.http.response import JsonResponse
from django.shortcuts import render
from django.views import View
import requests

class searchNews(View):
    def get(self,request):
        return render(request,'searchNews.html')

    # def post(self,request):
    #     st = request.POST.get('searchTerm')
    #     url="http://localhost:9200/_search?pretty"
    #     myObj= {
    #             "query": { 
    #                 "bool": { 
    #                 "must": [
    #                     { "match": { "title": st}}
    #                 ]
    #                 }
    #             }
    #         }
    #     res= requests.post(url,data=myObj)
    #     resp=res.json()
    #     print(resp)
    #     return render(request,"searchNews.html",resp)
