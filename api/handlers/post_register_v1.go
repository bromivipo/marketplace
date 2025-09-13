package handlers

import (
	"context"

	"github.com/bromivipo/marketplace/api/consts"
	generated "github.com/bromivipo/marketplace/api/definitions"
	"github.com/bromivipo/marketplace/api/pgrepo"
)

func (Server) PostRegisterV1(ctx context.Context, request generated.PostRegisterV1RequestObject) (generated.PostRegisterV1ResponseObject, error) {
	err := pgrepo.RegisterUser(request.Body.Username, request.Body.Password)
	if err == nil {
		return generated.PostRegisterV1200Response{}, nil
	} else {
		return generated.PostRegisterV1409JSONResponse{Code: consts.Conflict, Message: "This username is already taken"}, nil
	}
}
