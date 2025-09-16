package httphandlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/bromivipo/marketplace/api/consts"
	generated "github.com/bromivipo/marketplace/api/definitions"
	"github.com/bromivipo/marketplace/api/pgrepo"
)

func (Server) PostMakeOrderV1(ctx context.Context, request generated.PostMakeOrderV1RequestObject) (generated.PostMakeOrderV1ResponseObject, error) {
	token := strings.TrimPrefix(request.Params.Authorization , "Bearer ");
	updateErr := pgrepo.UpdateProducts(request.Body.ProductIds)
	if updateErr == nil {
		if err := pgrepo.InsertOrder(request.Body.ProductIds, token); err != nil {
			fmt.Print("ERROR:", err)
		}
		return generated.PostMakeOrderV1200Response{}, nil
	}
	switch updateErr.Reason {
		case pgrepo.OutOfStock:
			return generated.PostMakeOrderV1409JSONResponse{Code: consts.Conflict, Message: fmt.Sprintf("Product with id %v is out of stock", updateErr.Id)}, nil
		case pgrepo.NotFound:
			return generated.PostMakeOrderV1404JSONResponse{Code: consts.NotFound, Message: fmt.Sprintf("Product with id %v not found", updateErr.Id)}, nil
	}
	return nil, nil
}