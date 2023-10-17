// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"sampleapi/pls-payment/internal/dto"
	ierr "sampleapi/pls-shared/errors"
)

type serviceInvocationHandler interface {
	ProcessPayment(ctx context.Context, data []byte) (any, error)
}

func (o paymentHandler) ProcessPayment(ctx context.Context, data []byte) (any, error) {
	req := dto.ProcessPaymentReq{}
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, ierr.WrapErrorf(err, ierr.Unknown, "json.Unmarshal")
	}

	o.logger.Infof("Req data to payment ProcessPayment method: %v", req)

	if err := o.paymentSvc.ProcessPayment(ctx, req); err != nil {
		return nil, fmt.Errorf("paymentSvc.ProcessPayment: %w", err)
	}

	return nil, nil
}
