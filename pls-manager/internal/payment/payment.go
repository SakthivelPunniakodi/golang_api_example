// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package payment

import (
	"context"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
	daprclt "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/dapr/client"
	ierr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/errors"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
)

type Payment interface {
	ProcessPayment(ctx context.Context, req ProcessPaymentReq) error
}

type payment struct {
	logger logger.Logger
	client daprclt.Client
}

func NewPayment(logger logger.Logger, client daprclt.Client) Payment {
	return payment{
		logger: logger,
		client: client,
	}
}

func (o payment) ProcessPayment(ctx context.Context, req ProcessPaymentReq) error {
	o.logger.Infof("Req data to payment ProcessPayment method: %v", req)

	_, err := o.client.InvokeMethodWithContent(ctx, constant.Payment, constant.ProcessPayment, req)
	if err != nil {
		return ierr.WrapErrorf(err, ierr.Unknown, "client.InvokeMethodWithContent")
	}

	return nil
}
