package grpchandlers

import (
	"context"
	"log"

	generated "github.com/bromivipo/marketplace/api/definitions"
	"github.com/bromivipo/marketplace/api/pgrepo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*GrpcServer) RegisterNewProduct(_ context.Context, request *generated.ProductToRegister) (*generated.ProductId, error) {
	log.Println("INFO: Got register new product request")
	id, err := pgrepo.InsertNewProduct(request)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	return &generated.ProductId{Id: id}, nil
}