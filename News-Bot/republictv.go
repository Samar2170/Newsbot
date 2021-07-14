package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
)

func saveNews(link string) {
	dt := time.Now()
	dtt := dt.Format("01-02-2006")
	doc := getDoc(link)
	source := "RepublicTv"
	doc.Find("loc").Each(func(i int, s *goquery.Selection) {
		link = s.Text()
		n, err := url.Parse(link)
		if err != nil {
			log.Println(err)
		}
		m := strings.SplitAfter(n.Path, "/")
		o := m[len(m)-1]
		news := strings.Replace(o, ".html", "", 1)

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Println(err)
		}
		defer db.Close()
		sql := `INSERT INTO news(title, url, source, date)
				VALUES($1, $2, $3, $4)`
		_, err = db.Exec(sql, news, link, source, dtt)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Query Executed sucessfully")

	})
}

func parseRtv() {
	link := "https://www.republicworld.com/sitemap.xml"
	doc := getDoc(link)
	var newS []string
	doc.Find("loc").Each(func(i int, s *goquery.Selection) {
		str := s.Text()
		newS = append(newS, str)

	})
	fmt.Println(len(newS))
	link2 := newS[0]
	saveNews(link2)
	Wg.Done()
}
