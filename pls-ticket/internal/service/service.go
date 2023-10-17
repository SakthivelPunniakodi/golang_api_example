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
	"sync/atomic"
	"time"

	"sampleapi/pls-shared/constant"
	daprclt "sampleapi/pls-shared/dapr/client"
	"sampleapi/pls-shared/logger"
	"sampleapi/pls-shared/utils"
	"sampleapi/pls-ticket/internal/dto"

	"github.com/shopspring/decimal"
)

type Option func(*ticketService)

func WithLogger(logger logger.Logger) Option {
	return func(o *ticketService) {
		o.logger = logger
	}
}

func WithDaprClient(daprClient daprclt.Client) Option {
	return func(o *ticketService) {
		o.daprClient = daprClient
	}
}

type TicketService interface {
	GenerateTicket(ctx context.Context, req dto.GenerateTicketReq) (dto.GenerateTicketRes, error)
	GetTicketInfo(ctx context.Context, req dto.GetTicketInfoReq) (dto.GetTicketInfoRes, error)
	DiscardTicket(ctx context.Context, req dto.DiscardTicketReq) error
}

type ticketService struct {
	logger          logger.Logger
	daprClient      daprclt.Client
	tickets         map[string]dto.GenerateTicketRes
	incrementNumber uint64
}

// NewTicketService svc instance.
func NewTicketService(opts ...Option) TicketService {
	o := &ticketService{
		tickets: make(map[string]dto.GenerateTicketRes),
	}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (o *ticketService) GenerateTicket(ctx context.Context, req dto.GenerateTicketReq) (dto.GenerateTicketRes, error) {
	ticketId := atomic.AddUint64(&o.incrementNumber, 1)

	validationKey, err := utils.HashUint64ToUUID(ticketId)
	if err != nil {
		return dto.GenerateTicketRes{}, err
	}

	ticket := dto.GenerateTicketRes{
		TicketID: fmt.Sprintf("TICKET-%d", ticketId),
		// TicketID:      "TICKET-1",
		VehicleNumber: req.VehicleNumber,
		VehicleType:   req.VehicleType,
		EntryTime:     time.Now(),
		ValidationKey: validationKey.String(),
		// ValidationKey: "6b86b273-ff34-fce1-9d6b-804eff5a3f57",
		FloorNumber: req.FloorNumber,
		SpotNumber:  req.SpotNumber,
		SpotType:    req.SpotType,
	}

	o.tickets[ticket.TicketID] = ticket

	return ticket, nil
}

func (o *ticketService) GetTicketInfo(ctx context.Context, req dto.GetTicketInfoReq) (dto.GetTicketInfoRes, error) {
	ticket, err := o.validateTicket(req.TicketID, req.ValidationKey)
	if err != nil {
		fmt.Println(err)
		return dto.GetTicketInfoRes{}, err
	}

	exitTime := time.Now()
	duration := exitTime.Sub(ticket.EntryTime)
	fees := calculateParkingFees(ticket.SpotType, duration)

	return dto.GetTicketInfoRes{
		TicketID:      ticket.TicketID,
		VehicleNumber: ticket.VehicleNumber,
		VehicleType:   ticket.VehicleType,
		EntryTime:     ticket.EntryTime,
		FloorNumber:   ticket.FloorNumber,
		SpotNumber:    ticket.SpotNumber,
		SpotType:      ticket.SpotType,
		Fees:          fees,
	}, nil
}

func (o *ticketService) DiscardTicket(ctx context.Context, req dto.DiscardTicketReq) error {
	_, err := o.validateTicket(req.TicketID, req.ValidationKey)
	if err != nil {
		return err
	}

	delete(o.tickets, req.TicketID)

	return nil
}

func (o *ticketService) validateTicket(ticketID, validationKey string) (dto.GenerateTicketRes, error) {
	tID, err := utils.ExtractNumber(ticketID)
	if err != nil {
		return dto.GenerateTicketRes{}, err
	}

	vKey, err := utils.HashUint64ToUUID(tID)
	if err != nil {
		return dto.GenerateTicketRes{}, err
	}

	if validationKey != vKey.String() {
		return dto.GenerateTicketRes{}, fmt.Errorf("Invalid validation key")
	}

	ticket, ok := o.tickets[ticketID]
	if !ok {
		return dto.GenerateTicketRes{}, fmt.Errorf("Ticket not found")
	}

	if ticket.ValidationKey != validationKey {
		return dto.GenerateTicketRes{}, fmt.Errorf("Invalid validation key")
	}

	return ticket, nil
}

func calculateParkingFees(spotType constant.ParkingSpotType, duration time.Duration) decimal.Decimal {
	return getFeeAccordingToSpotType(spotType).Add(getFeeAccordingToDuration(duration))
}

func getFeeAccordingToSpotType(spotType constant.ParkingSpotType) decimal.Decimal {
	switch spotType {
	case constant.ParkingSpotTypeMotorcycle:
	case constant.ParkingSpotTypeHandicapped:
		return decimal.NewFromInt(5)
	case constant.ParkingSpotTypeCompact:
		return decimal.NewFromInt(10)
	case constant.ParkingSpotTypeLarge:
		return decimal.NewFromInt(20)
	case constant.ParkingSpotTypeEVChargeStation:
		return decimal.NewFromInt(50)
	}

	return decimal.NewFromInt(0)
}

func getFeeAccordingToDuration(duration time.Duration) decimal.Decimal {
	hours := duration.Hours()
	if hours < 1 {
		return decimal.NewFromInt(10)
	}

	if hours < 2 {
		return decimal.NewFromFloat(17.5)
	}

	if hours < 3 {
		return decimal.NewFromFloat(22.5)
	}

	return decimal.NewFromInt(40)
}
