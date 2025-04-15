package data

type Product struct {
	Id          string  `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  string  `json:"categoryId"`
	UnitePrice  float64 `json:"unitePrice"`
}

type Category struct {
	Id           string `json:"id"`
	CategoryName string `json:"categoryName"`
}
