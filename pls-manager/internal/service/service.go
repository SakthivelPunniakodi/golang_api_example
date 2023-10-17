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

	"sampleapi/pls-manager/internal/dto"
	"sampleapi/pls-manager/internal/payment"
	"sampleapi/pls-manager/internal/space"
	"sampleapi/pls-manager/internal/ticket"
	daprclt "sampleapi/pls-shared/dapr/client"
	spaceEvents "sampleapi/pls-shared/events/space"
	"sampleapi/pls-shared/logger"
	"sampleapi/pls-shared/utils"
)

type Option func(*manager)

func WithLogger(logger logger.Logger) Option {
	return func(o *manager) {
		o.logger = logger
	}
}

func WithDaprClient(daprClient daprclt.Client) Option {
	return func(o *manager) {
		o.daprClient = daprClient
	}
}

func WithPaymentSvc(paymentSvc payment.Payment) Option {
	return func(o *manager) {
		o.paymentSvc = paymentSvc
	}
}

func WithSpaceSvc(spaceSvc space.Space) Option {
	return func(o *manager) {
		o.spaceSvc = spaceSvc
	}
}

func WithTicketSvc(ticketSvc ticket.Ticket) Option {
	return func(o *manager) {
		o.ticketSvc = ticketSvc
	}
}

func WithUnoccupiedSpotsChannel(unoccupiedSpotsCh chan<- []spaceEvents.Floor) Option {
	return func(o *manager) {
		o.unoccupiedSpotsCh = unoccupiedSpotsCh
	}
}

type Manager interface {
	AllocateSpot(ctx context.Context, req dto.AllocateSpotReq) (dto.AllocateSpotRes, error)
	GetTicketInfo(ctx context.Context, req dto.GetTicketInfoReq) (dto.GetTicketInfoRes, error)
	DeallocateSpot(ctx context.Context, req dto.DeallocateSpotReq) error
	SpotChanged(ctx context.Context, req dto.SpotChangedReq)
}

type manager struct {
	logger            logger.Logger
	daprClient        daprclt.Client
	paymentSvc        payment.Payment
	spaceSvc          space.Space
	ticketSvc         ticket.Ticket
	unoccupiedSpotsCh chan<- []spaceEvents.Floor
}

// NewManager svc instance.
func NewManager(opts ...Option) Manager {
	o := &manager{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (o manager) AllocateSpot(ctx context.Context, req dto.AllocateSpotReq) (dto.AllocateSpotRes, error) {
	spotType, err := utils.GetParkingSpotType(
		req.VehicleTypeEnum,
		req.IsNeedHandicappedSpot,
		req.IsNeedEVChargeStationSpot,
	)
	if err != nil {
		return dto.AllocateSpotRes{}, fmt.Errorf("utils.GetParkingSpotType: %w", err)
	}

	allocateSpaceRes, err := o.spaceSvc.AllocateSpace(ctx, space.AllocateSpaceReq{
		SpotType: spotType,
	})
	if err != nil {
		return dto.AllocateSpotRes{}, fmt.Errorf("spaceSvc.AllocateSpace: %w", err)
	}

	generateTicketRes, err := o.ticketSvc.GenerateTicket(ctx, ticket.GenerateTicketReq{
		VehicleNumber: req.VehicleNumber,
		VehicleType:   req.VehicleTypeEnum,
		FloorNumber:   allocateSpaceRes.FloorNumber,
		SpotNumber:    allocateSpaceRes.SpotNumber,
		SpotType:      allocateSpaceRes.SpotType,
	})
	if err != nil {
		return dto.AllocateSpotRes{}, fmt.Errorf("ticketSvc.GenerateTicket: %w", err)
	}

	return dto.AllocateSpotRes{
		TicketID:      generateTicketRes.TicketID,
		VehicleNumber: generateTicketRes.VehicleNumber,
		VehicleType:   utils.GetVehicleTypeString(generateTicketRes.VehicleType),
		EntryTime:     generateTicketRes.EntryTime,
		ValidationKey: generateTicketRes.ValidationKey,
		FloorNumber:   generateTicketRes.FloorNumber,
		SpotNumber:    generateTicketRes.SpotNumber,
		SpotType:      generateTicketRes.SpotType,
	}, nil
}

func (o manager) GetTicketInfo(ctx context.Context, req dto.GetTicketInfoReq) (dto.GetTicketInfoRes, error) {
	getTicketInfoRes, err := o.ticketSvc.GetTicketInfo(ctx, ticket.GetTicketInfoReq{
		TicketID:      req.TicketID,
		ValidationKey: req.ValidationKey,
	})
	if err != nil {
		return dto.GetTicketInfoRes{}, fmt.Errorf("ticketSvc.GetTicketInfo: %w", err)
	}

	return dto.GetTicketInfoRes{
		TicketID:      getTicketInfoRes.TicketID,
		VehicleNumber: getTicketInfoRes.VehicleNumber,
		VehicleType:   utils.GetVehicleTypeString(getTicketInfoRes.VehicleType),
		EntryTime:     getTicketInfoRes.EntryTime,
		FloorNumber:   getTicketInfoRes.FloorNumber,
		SpotNumber:    getTicketInfoRes.SpotNumber,
		SpotType:      getTicketInfoRes.SpotType,
		Fees:          getTicketInfoRes.Fees,
	}, nil
}

func (o manager) DeallocateSpot(ctx context.Context, req dto.DeallocateSpotReq) error {
	getTicketInfoRes, err := o.ticketSvc.GetTicketInfo(ctx, ticket.GetTicketInfoReq{
		TicketID:      req.TicketID,
		ValidationKey: req.ValidationKey,
	})
	if err != nil {
		return fmt.Errorf("ticketSvc.GetTicketInfo: %w", err)
	}

	if err := o.paymentSvc.ProcessPayment(ctx, payment.ProcessPaymentReq{
		TicketID:      req.TicketID,
		Amount:        getTicketInfoRes.Fees,
		CVV:           req.CVV,
		CardNumber:    req.CardNumber,
		CardType:      req.CardType,
		PaymentMethod: req.PaymentMethod,
	}); err != nil {
		return fmt.Errorf("paymentSvc.ProcessPayment: %w", err)
	}

	if err := o.spaceSvc.DeallocateSpace(ctx, space.DeallocateSpaceReq{
		FloorNumber: getTicketInfoRes.FloorNumber,
		SpotNumber:  getTicketInfoRes.SpotNumber,
		SpotType:    getTicketInfoRes.SpotType,
	}); err != nil {
		return fmt.Errorf("spaceSvc.DeallocateSpace: %w", err)
	}

	if err := o.ticketSvc.DiscardTicket(ctx, ticket.DiscardTicketReq{
		TicketID:      req.TicketID,
		ValidationKey: req.ValidationKey,
	}); err != nil {
		return fmt.Errorf("ticketSvc.DiscardTicket: %w", err)
	}

	return nil
}

func (o manager) SpotChanged(ctx context.Context, req dto.SpotChangedReq) {
	o.unoccupiedSpotsCh <- req.Floors
}
