package models

import "gorm.io/gorm"

type StockDecreaseLog struct {
	// Id int64 `json:"id" gorm:"primaryKey"`
	gorm.Model
	OrderId      int64 `json:"order_id"`
	ProductRefer int64 `json:"product_id"`
}
