package store

import "github.com/djermy/runeprofitforge-golang-api/runescape"

type Store interface {
	GetItems() ([]runescape.Item, error)
	GetItem(string) (runescape.Item, error)
	CreateItem(*runescape.Item) error
	UpdateItem(string, *runescape.Item) error
	DeleteItem(string) error
}
