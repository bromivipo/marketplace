package grpchandlers

import generated "github.com/bromivipo/marketplace/api/definitions"

type GrpcServer struct {
	generated.UnimplementedMarketplaceInternalServer
}

func NewGrpcServer() GrpcServer {
	return GrpcServer{}
}