package main

import "sync"

var Wg sync.WaitGroup

func main() {
	Wg.Add(1)
	go parseEt()
	Wg.Add(1)
	go parseTOI()
	Wg.Add(1)
	go parseRtv()
	Wg.Wait()
}
