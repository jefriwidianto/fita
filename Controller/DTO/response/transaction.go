package response

type ItemsCartDTO struct {
	Sku   string  `json:"sku"`
	Name  string  `json:"name"`
	Qty   int     `json:"qty"`
}

type CheckingOutDTO struct {
	Sku        string   `json:"sku"`
	Name       []string `json:"name"`
	TotalPrice string   `json:"total_price"`
	TotalQty   int      `json:"total_qty"`
}
