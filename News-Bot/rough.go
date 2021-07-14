package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

func sampleDate() {
	dt := time.Now()
	dtt := dt.Format("01-02-2006")
	fmt.Println(dt.Format("01-02-2006"))
	fmt.Printf("%v %T", dtt, dtt)
}

func test() {
	link := "https://www.republicworld.com/india-news/law-and-order/as-soon-as-they-heard-my-name-i-was-attacked-advocate-on-being-beaten-up-by-police.html"
	u, err := url.Parse(link)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	// fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
	res1 := strings.Split(u.Path, "/")
	res2 := strings.SplitAfter(u.Path, "/")
	res3 := res2[len(res2)-1]
	res3 = strings.Replace(res3, ".html", "", 1)
	fmt.Println(res1[0])
	fmt.Printf("%T\n", res2)
	fmt.Println(res3)
}
