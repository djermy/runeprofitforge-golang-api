package main

import (
	"log"

	"github.com/djermy/runeprofitforge-golang-api/runescape"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	rs := runescape.New() // create new instance of runescape client

	items, err := rs.GetItems(0, "a", 0)
	if err != nil {
		panic(err)
	}
	log.Println(items)
}
