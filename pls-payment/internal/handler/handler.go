// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	iservice "github.com/SakthivelPunniakodi/golang_api_example/pls-payment/internal/service"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
)

type Option func(*paymentHandler)

func WithLogger(logger logger.Logger) Option {
	return func(o *paymentHandler) {
		o.logger = logger
	}
}

func WithSvc(paymentSvc iservice.PaymentService) Option {
	return func(o *paymentHandler) {
		o.paymentSvc = paymentSvc
	}
}

type PaymentHandler interface {
	serviceInvocationHandler
}

type paymentHandler struct {
	logger     logger.Logger
	paymentSvc iservice.PaymentService
}

// NewPaymentHandler handler instance.
func NewPaymentHandler(opts ...Option) PaymentHandler {
	o := &paymentHandler{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
