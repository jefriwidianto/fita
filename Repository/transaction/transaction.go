package transaction

import (
	"encoding/json"
	"fita/Config"
	"fita/Controller/DTO/request"
	"fita/Controller/DTO/response"
	"github.com/graphql-go/graphql"
	"strconv"
)

func checkingOut(params graphql.ResolveParams) (interface{}, error) {
	db := Config.DATABASE_MAIN.Get()
	var items []*response.CheckingOutDTO
	var resolvedCells request.Cells
	var nameItems string
	var price float64

	cellsByte, _ := json.Marshal(params.Args["params"])
	if err := json.Unmarshal(cellsByte, &resolvedCells); err != nil {
		return nil, err
	}

	for _, value := range resolvedCells {
		rows, err := db.Query(`SELECT tc.sku, tc.name, ti.price
				FROM t_cart tc 
				INNER JOIN t_items ti ON tc.sku = ti.sku
				WHERE tc.sku = $1`, value.Sku)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			item := &response.CheckingOutDTO{}
			err = rows.Scan(&item.Sku, &nameItems, &price)
			item.TotalPrice = strconv.FormatFloat(price * float64(value.Qty), 'f', 2, 64)
			item.TotalQty = value.Qty
			item.Name = append(item.Name, nameItems)

			if value.Sku == kodePromo1 {
				item.Name = append(item.Name, productFree)
				item.TotalQty = value.Qty
			} else if value.Sku == kodePromo2 && value.Qty == 3 {
				item.TotalPrice = strconv.FormatFloat(price * 2, 'f', 2, 64)
				item.TotalQty = value.Qty
			} else if value.Sku == kodePromo3 && value.Qty >= 3 {
				totalPrice := price * float64(value.Qty)
				calculatedDiskon := totalPrice - (totalPrice / 100 * 10)
				item.TotalPrice = strconv.FormatFloat(calculatedDiskon, 'f', 2, 64)
				item.TotalQty = value.Qty
			}

			items = append(items, item)
		}
	}

	return items, nil
}

func addToChart(params graphql.ResolveParams) (interface{}, error) {
	db := Config.DATABASE_MAIN.Get()
	_, err := db.Exec(`INSERT INTO t_cart(sku, name, qty) VALUES($1, $2, $3)`, params.Args["sku"].(string),
		params.Args["name"].(string), params.Args["qty"].(int))
	if err != nil {
		return nil, err
	}

	item := &response.ItemsCartDTO{
		Sku: params.Args["sku"].(string),
		Name: params.Args["name"].(string),
		Qty: params.Args["qty"].(int),
	}
	return item, nil
}
