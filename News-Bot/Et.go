package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
)

func saveNewsEt(s *goquery.Document) {
	dt := time.Now()
	dtt := dt.Format("01-02-2006")
	doc := s
	var headline, url string
	source := "ET"
	doc.Find(".flt").Each(func(i int, s *goquery.Selection) {
		headline = s.Text()
		// urlS := s.Find("flr")
		urlS := s.Find("a")
		url, _ = urlS.Attr("href")
		fmt.Printf("%s ,  %s\n", headline, source)
		fmt.Println(url)

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Println(err)
		}
		defer db.Close()
		sql := `INSERT INTO news(title, url, source, date)
				VALUES($1, $2, $3, $4)`
		_, err = db.Exec(sql, headline, url, source, dtt)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Query Executed sucessfully")
	})

	doc.Find(".flr").Each(func(i int, s *goquery.Selection) {
		headline = s.Text()
		// urlS := s.Find("flr")
		urlS := s.Find("a")
		url, _ = urlS.Attr("href")
		fmt.Printf("%s ,  %s\n", headline, source)
		fmt.Println(url)

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Println(err)
		}
		defer db.Close()
		sql := `INSERT INTO news(title, url, source,date)
				VALUES($1, $2, $3, $4)`
		_, err = db.Exec(sql, headline, url, source, dtt)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Query Executed sucessfully")
	})
}

func parseEt() {
	url := "https://economictimes.indiatimes.com/headlines.cms"
	doc := getDoc(url)
	saveNewsEt(doc)
	Wg.Done()
}
