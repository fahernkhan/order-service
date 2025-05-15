package models

import "time"

/*

	order_id
	products: [
		product_id: 1 qty: 2
		product_id: 2 qty: 1
	]
	event_time

*/

type ProductStockUpdateEvent struct {
	OrderID   int64         `json:"order_id"`
	Products  []ProductItem `json:"products"`
	EventTime time.Time     `json:"event_time"`
}

type ProductItem struct {
	ProductID int64 `json:"product_id"`
	Qty       int   `json:"qty"`
}
