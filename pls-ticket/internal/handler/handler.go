// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
	iservice "github.com/SakthivelPunniakodi/golang_api_example/pls-ticket/internal/service"
)

type Option func(*ticketHandler)

func WithLogger(logger logger.Logger) Option {
	return func(o *ticketHandler) {
		o.logger = logger
	}
}

func WithSvc(ticketSvc iservice.TicketService) Option {
	return func(o *ticketHandler) {
		o.ticketSvc = ticketSvc
	}
}

type TicketHandler interface {
	serviceInvocationHandler
}

type ticketHandler struct {
	logger    logger.Logger
	ticketSvc iservice.TicketService
}

// NewTicketHandler handler instance.
func NewTicketHandler(opts ...Option) TicketHandler {
	o := &ticketHandler{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
