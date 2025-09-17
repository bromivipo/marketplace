package handlers

import (
	"context"
	"log"
	"strings"

	"github.com/bromivipo/marketplace/partner_api/consts"
	generated "github.com/bromivipo/marketplace/partner_api/definitions"
	"github.com/bromivipo/marketplace/partner_api/pgrepo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func (Server) PostAddProductsInStockV1(ctx context.Context, request generated.PostAddProductsInStockV1RequestObject) (generated.PostAddProductsInStockV1ResponseObject, error) {
	token := strings.TrimPrefix(request.Params.Authorization , "Bearer ");
	id := pgrepo.SelectPartner(token)
	if id == nil {
		return generated.PostAddProductsInStockV1401JSONResponse{Code: consts.Unauthorized, Message: "Token is missing or unknown"}, nil
	}
	conn, err := grpc.NewClient(pgrepo.GetEnvOrDefault("API_GRPC_ADDRESS", "localhost:50051"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	client := generated.NewMarketplaceInternalClient(conn)
	grpcReq := generated.AddProductsRequest{}
	for _, product := range request.Body.Products {
		grpcReq.Products = append(grpcReq.Products, &generated.ProductToAdd{Id: int32(product.ProductId), Amount: int32(product.NumberToAdd)})
	}
	_, err = client.AddProductsInStock(ctx, &grpcReq)
	if err != nil {
		return generated.PostAddProductsInStockV1404JSONResponse{Code: consts.NotFound, Message: err.Error()}, nil
	}
	return generated.PostAddProductsInStockV1200Response{}, nil
}