package item

import "fmt"

// Item represents the item
type Item struct {
	Name        string `bson: "Name"`
	Type        int    `bson: "Type"`
	ImageURL    string `bson: "ImageURL"`
	Price       int    `bson: "Price"`
	Description string `bson: "Description"`
}

// GetPrice returns string
func (i *Item) GetPrice() string {
	return fmt.Sprintf("%d円", i.Price)
}
