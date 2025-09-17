package handlers

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/bromivipo/marketplace/partner_api/consts"
	generated "github.com/bromivipo/marketplace/partner_api/definitions"
	"github.com/bromivipo/marketplace/partner_api/pgrepo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func (Server) PostRegisterNewProductV1(ctx context.Context, request generated.PostRegisterNewProductV1RequestObject) (generated.PostRegisterNewProductV1ResponseObject, error) {
	token := strings.TrimPrefix(request.Params.Authorization , "Bearer ");
	id := pgrepo.SelectPartner(token)
	if id == nil {
		return generated.PostRegisterNewProductV1401JSONResponse{Code: consts.Unauthorized, Message: "Token is missing or unknown"}, nil
	}
	conn, err := grpc.NewClient(pgrepo.GetEnvOrDefault("API_GRPC_ADDRESS", "localhost:50051"),  grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()	
	client := generated.NewMarketplaceInternalClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.RegisterNewProduct(ctx, &generated.ProductToRegister{Name: request.Body.Name, Price: request.Body.Price, LeftInStock: int32(request.Body.LeftInStock), ProviderId: int32(*id), Category: request.Body.Category})
	if err != nil {
		return generated.PostRegisterNewProductV1404JSONResponse{Code: consts.NotFound, Message: err.Error()}, nil
	}
	return generated.PostRegisterNewProductV1200JSONResponse{ProductId: int(response.Id)}, nil
}