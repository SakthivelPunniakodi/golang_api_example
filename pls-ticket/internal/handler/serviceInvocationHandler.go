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

	ierr "sampleapi/pls-shared/errors"
	"sampleapi/pls-ticket/internal/dto"
)

type serviceInvocationHandler interface {
	GenerateTicket(ctx context.Context, data []byte) (any, error)
	GetTicketInfo(ctx context.Context, data []byte) (any, error)
	DiscardTicket(ctx context.Context, data []byte) (any, error)
}

func (o ticketHandler) GenerateTicket(ctx context.Context, data []byte) (any, error) {
	req := dto.GenerateTicketReq{}
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, ierr.WrapErrorf(err, ierr.Unknown, "json.Unmarshal")
	}

	o.logger.Infof("Req data to ticket GenerateTicket method: %v", req)

	res, err := o.ticketSvc.GenerateTicket(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("ticketSvc.GenerateTicket: %w", err)
	}

	return res, nil
}

func (o ticketHandler) GetTicketInfo(ctx context.Context, data []byte) (any, error) {
	req := dto.GetTicketInfoReq{}
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, ierr.WrapErrorf(err, ierr.Unknown, "json.Unmarshal")
	}

	o.logger.Infof("Req data to ticket GetTicketInfo method: %v", req)

	res, err := o.ticketSvc.GetTicketInfo(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("ticketSvc.GetTicketInfo: %w", err)
	}

	return res, nil
}

func (o ticketHandler) DiscardTicket(ctx context.Context, data []byte) (any, error) {
	req := dto.DiscardTicketReq{}
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, ierr.WrapErrorf(err, ierr.Unknown, "json.Unmarshal")
	}

	o.logger.Infof("Req data to ticket DiscardTicket method: %v", req)

	if err := o.ticketSvc.DiscardTicket(ctx, req); err != nil {
		return nil, fmt.Errorf("ticketSvc.DiscardTicket: %w", err)
	}

	return nil, nil
}
