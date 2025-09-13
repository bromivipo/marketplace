package handlers

import (
	"context"

	"github.com/bromivipo/marketplace/api/consts"
	generated "github.com/bromivipo/marketplace/api/definitions"
	"github.com/bromivipo/marketplace/api/pgrepo"
)

func (Server) PostLoginV1(ctx context.Context, request generated.PostLoginV1RequestObject) (generated.PostLoginV1ResponseObject, error) {
	password := pgrepo.GetUserPassword(request.Body.Username)
	if password == nil {
		return generated.PostLoginV1404JSONResponse{Code: consts.NotFound, Message: "User wasn`t found"}, nil
	}
	if *password == request.Body.Password {
		return generated.PostLoginV1200JSONResponse{Token: "token"}, nil
	} else {
		return generated.PostLoginV1401JSONResponse{Code: consts.Unauthorized, Message: "Wrong password"}, nil
	}
}
