from django.urls import path
import core.views as v
urlpatterns = [
    path('search/', v.searchNews.as_view(), name='search' ),
]
