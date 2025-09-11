package definitions

import "github.com/shopspring/decimal"

type ProductItem struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Price decimal.Decimal `json:"price"`
	Category string `json:"category"`
}

type GetProductsV1Response struct {
	Products []ProductItem `json:"products"`
}

type ErrorResponse struct {
	Code string `json:"code"`
	Message string `json:"message"`
}