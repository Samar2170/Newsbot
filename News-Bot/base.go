package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin@123"
	dbname   = "newsbot"
)

func getDoc(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("%s", err)
	} else {
		log.Println("Request sucessful")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Println("request unsucessful")
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		print("error while reading to goquery")
	}
	return doc
}
