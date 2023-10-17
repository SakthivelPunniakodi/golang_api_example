// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package service

import (
	"context"
	"fmt"
	"time"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-payment/internal/dto"
	daprclt "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/dapr/client"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
)

type Option func(*paymentService)

func WithLogger(logger logger.Logger) Option {
	return func(o *paymentService) {
		o.logger = logger
	}
}

func WithDaprClient(daprClient daprclt.Client) Option {
	return func(o *paymentService) {
		o.daprClient = daprClient
	}
}

type PaymentService interface {
	ProcessPayment(ctx context.Context, req dto.ProcessPaymentReq) error
}

type paymentService struct {
	logger     logger.Logger
	daprClient daprclt.Client
}

// NewPaymentService svc instance.
func NewPaymentService(opts ...Option) PaymentService {
	o := &paymentService{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (o paymentService) ProcessPayment(ctx context.Context, req dto.ProcessPaymentReq) error {
	// Simulate payment processing (in a real system, have to integrate with a payment gateway)
	time.Sleep(2 * time.Second)

	fmt.Printf("Payment processed successfully for ticket ID: %s,Amount Paid: $%s\n", req.TicketID, req.Amount)

	return nil
}
