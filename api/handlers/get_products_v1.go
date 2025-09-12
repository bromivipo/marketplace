package handlers

import (
	"context"

	generated "github.com/bromivipo/marketplace/api/definitions"
	"github.com/bromivipo/marketplace/api/pgrepo"
)

func (Server) GetProductsV1(ctx context.Context, request generated.GetProductsV1RequestObject) (generated.GetProductsV1ResponseObject, error) {
	return generated.GetProductsV1200JSONResponse{Products: pgrepo.GetProducts()}, nil
}
