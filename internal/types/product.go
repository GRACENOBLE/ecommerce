package types

type Product struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         int64  `json:"price"`
	StockQuantity string `json:"stock"`
	ImageURL      string `json:"image"`
	CreatedAt     string `json:"created-at"`
	UpdatedAT     string `json:"updated-at"`
}
