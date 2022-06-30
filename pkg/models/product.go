package models

import ( 
	"github.com/lib/pq"
)


type Product struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Img string `json:"img"`
	Categories pq.StringArray `json:"categories" gorm:"type:text[]"`
	Size string `json:"size"`
	Color string `json:"color"`
	Price int64 `json:"price"`
	Stock int64 `json:"stock"`
	// Name string `json:"name"`
	// Stock int64 `json:"stock"`
	// Price int64 `json:"price"`
	// Image string `json:"image"`
	// Description string `json:"description"`
	StockDecreaseLog StockDecreaseLog `gorm:"foreignKey:ProductRefer"`
}