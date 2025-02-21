package models

type Product struct {
	ProductID   string  `json:"product_id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockQty    int     `json:"stock_qty"`
	Category    string  `json:"category"`
	SKU         string  `json:"sku"`
	TaxRate     float64 `json:"tax_rate"`
}
