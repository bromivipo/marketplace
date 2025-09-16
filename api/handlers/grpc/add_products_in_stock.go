package grpchandlers

import (
	"context"

	generated "github.com/bromivipo/marketplace/api/definitions"
	"github.com/bromivipo/marketplace/api/pgrepo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*GrpcServer) AddProductsInStock(_ context.Context, request *generated.AddProductsRequest) (*generated.Empty, error) {
	updateErr := pgrepo.UpdateProductsAmount(request)
	if updateErr != nil {
		return nil, status.Errorf(codes.NotFound, "Product with id %v was not found", updateErr.Id)
	}
	return &generated.Empty{}, nil
}