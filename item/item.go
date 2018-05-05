package item

import "fmt"

// Item represents the item
type Item struct {
	Name        string
	ImageURL    string
	Price       int
	Description string
}

// GetPrice returns string
func (i *Item) GetPrice() string {
	return fmt.Sprintf("%d円", i.Price)
}
