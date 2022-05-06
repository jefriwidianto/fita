package request

type ReqBody struct {
	Mutation string `json:"mutation"`
	Query    string `json:"query"`
}

type Cell struct {
	Sku string `json:"sku"`
	Qty  int `json:"qty"`
}

type Cells []*Cell
