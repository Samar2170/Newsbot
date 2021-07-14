package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
)

func saveNewsToi(s *goquery.Document) {
	dt := time.Now()
	dtt := dt.Format("01-02-2006")
	doc := s
	var headline, url string
	source := "TOI"
	count := 0
	doc.Find(".w_tle").Each(func(i int, s *goquery.Selection) {
		headline = s.Text()
		url, _ = s.Children().Attr("href")
		fmt.Printf("%s \n", headline)
		count += 1

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

func parseTOI() {
	url := "https://timesofindia.indiatimes.com/home/headlines"
	doc := getDoc(url)
	saveNewsToi(doc)
	Wg.Done()
}
