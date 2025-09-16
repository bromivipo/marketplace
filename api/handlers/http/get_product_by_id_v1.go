package httphandlers

import (
	"context"
	"fmt"

	"github.com/bromivipo/marketplace/api/consts"
	generated "github.com/bromivipo/marketplace/api/definitions"
	"github.com/bromivipo/marketplace/api/pgrepo"
)


func (Server) GetProductByIdV1(ctx context.Context, request generated.GetProductByIdV1RequestObject) (generated.GetProductByIdV1ResponseObject, error) {
	product := pgrepo.GetProductById(request.Params.Id)

	if product == nil {
		return generated.GetProductByIdV1404JSONResponse{Code: consts.NotFound, Message: fmt.Sprintf("product with id %v not found", request.Params.Id)}, nil
	}

	return generated.GetProductByIdV1200JSONResponse{Product: *product}, nil
}
