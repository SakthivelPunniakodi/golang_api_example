// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package ticket

import (
	"context"
	"encoding/json"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
	daprclt "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/dapr/client"
	ierr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/errors"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
)

type Ticket interface {
	GenerateTicket(ctx context.Context, req GenerateTicketReq) (GenerateTicketRes, error)
	GetTicketInfo(ctx context.Context, req GetTicketInfoReq) (GetTicketInfoRes, error)
	DiscardTicket(ctx context.Context, req DiscardTicketReq) error
}

type ticket struct {
	logger logger.Logger
	client daprclt.Client
}

func NewTicket(logger logger.Logger, client daprclt.Client) Ticket {
	return ticket{
		logger: logger,
		client: client,
	}
}

func (o ticket) GenerateTicket(ctx context.Context, req GenerateTicketReq) (GenerateTicketRes, error) {
	o.logger.Infof("Req data to ticket GenerateTicket method: %v", req)

	data, err := o.client.InvokeMethodWithContent(ctx, constant.Ticket, constant.GenerateTicket, req)
	if err != nil {
		return GenerateTicketRes{}, ierr.WrapErrorf(err, ierr.Unknown, "client.InvokeMethodWithContent")
	}

	res := GenerateTicketRes{}
	if err := json.Unmarshal(data, &res); err != nil {
		return GenerateTicketRes{}, ierr.WrapErrorf(err, ierr.Unknown, "json.Unmarshal")
	}

	o.logger.Infof("Res data from ticket GenerateTicket method: %v", res)

	return res, nil
}

func (o ticket) GetTicketInfo(ctx context.Context, req GetTicketInfoReq) (GetTicketInfoRes, error) {
	o.logger.Infof("Req data to ticket GetTicketInfo method: %v", req)

	data, err := o.client.InvokeMethodWithContent(ctx, constant.Ticket, constant.CalculateFees, req)
	if err != nil {
		return GetTicketInfoRes{}, ierr.WrapErrorf(err, ierr.Unknown, "client.InvokeMethodWithContent")
	}

	res := GetTicketInfoRes{}
	if err := json.Unmarshal(data, &res); err != nil {
		return GetTicketInfoRes{}, ierr.WrapErrorf(err, ierr.Unknown, "json.Unmarshal")
	}

	o.logger.Infof("Res data from ticket GetTicketInfo method: %v", res)

	return res, nil
}

func (o ticket) DiscardTicket(ctx context.Context, req DiscardTicketReq) error {
	o.logger.Infof("Req data to ticket DiscardTicket method: %v", req)

	_, err := o.client.InvokeMethodWithContent(ctx, constant.Ticket, constant.DiscardTicket, req)
	if err != nil {
		return ierr.WrapErrorf(err, ierr.Unknown, "client.InvokeMethodWithContent")
	}

	return nil
}
