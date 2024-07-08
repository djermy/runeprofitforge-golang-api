package main

import (
	"log"

	"github.com/djermy/runeprofitforge-golang-api/runescape"
	"github.com/djermy/runeprofitforge-golang-api/store"
	"github.com/djermy/runeprofitforge-golang-api/store/psqlstore"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	rs := runescape.New() // create new instance of runescape client
	var err error
	var s store.Store

	s, err = psqlstore.New()
	if err != nil {
		panic(err)
	}

	items, err := rs.GetItems(0, "a", 0)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		err = s.CreateItem(&item)
		if err != nil {
			log.Println(err)
			log.Printf("failed to get item with id: %d", item.ID)
		}
	}

	s.GetItems()
}
