package response_models

type Receipt struct {
	Id            string `json:"id"`
	Retailer      string `json:"retailer"`
	PurchaseDate  string `json:"purchaseDate"`
	PurchaseTime  string `json:"purchaseTime"`
	Items         []Item `json:"items"`
	Total         string `json:"total"`
	Points        int    `json:"points"`
}

type Item struct {
	Id               string `json:"id"`
	ReceiptId        string `json:"receiptId"`
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}