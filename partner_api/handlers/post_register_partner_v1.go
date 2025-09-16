package handlers

import (
	"context"

	"github.com/bromivipo/marketplace/partner_api/pgrepo"
	"github.com/bromivipo/marketplace/partner_api/consts"
	generated "github.com/bromivipo/marketplace/partner_api/definitions"
	"github.com/google/uuid"
)

func (Server) PostRegisterPartnerV1(ctx context.Context, request generated.PostRegisterPartnerV1RequestObject) (generated.PostRegisterPartnerV1ResponseObject, error) {
	token := uuid.New()
	err := pgrepo.RegisterPartner(request.Body.PartnerName, token.String())
	if err == nil {
		return generated.PostRegisterPartnerV1200JSONResponse{PermanentToken: token.String()}, nil
	} else {
		return generated.PostRegisterPartnerV1409JSONResponse{Code: consts.Conflict, Message: "This partner is already registered"}, nil
	}
}
