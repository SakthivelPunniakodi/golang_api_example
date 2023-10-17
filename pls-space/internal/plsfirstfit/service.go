// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package plsfirstfit

import (
	"context"
	"fmt"
	"sync"

	daprclt "sampleapi/pls-shared/dapr/client"
	"sampleapi/pls-shared/logger"
	"sampleapi/pls-space/internal/common"
	"sampleapi/pls-space/internal/dto"
)

type Option func(*spaceService)

func WithLogger(logger logger.Logger) Option {
	return func(o *spaceService) {
		o.logger = logger
	}
}

func WithDaprClient(daprClient daprclt.Client) Option {
	return func(o *spaceService) {
		o.daprClient = daprClient
	}
}

func WithParkingLot(parkingLot *ParkingLot) Option {
	return func(o *spaceService) {
		o.parkingLot = parkingLot
	}
}

type spaceService struct {
	logger     logger.Logger
	daprClient daprclt.Client
	parkingLot *ParkingLot
	mu         sync.Mutex
}

// NewSpaceService svc instance.
func NewSpaceService(opts ...Option) common.SpaceService {
	o := &spaceService{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (o *spaceService) AllocateSpace(ctx context.Context, req dto.AllocateSpaceReq) (dto.AllocateSpaceRes, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	for _, floor := range o.parkingLot.Floors {
		spotNumber, err := floor.AllocateSlot(req.SpotType)
		if err != nil {
			continue
		}

		fmt.Printf("Allocated parking spot on Floor %d, Spot %d, Type: %s\n", floor.FloorNumber, spotNumber, req.SpotType)

		if err := o.PublishSpotChangedEvent(ctx); err != nil {
			return dto.AllocateSpaceRes{}, err
		}

		return dto.AllocateSpaceRes{
			FloorNumber: floor.FloorNumber,
			SpotNumber:  spotNumber,
			SpotType:    req.SpotType,
		}, nil
	}

	return dto.AllocateSpaceRes{}, fmt.Errorf("aa")
}

func (o *spaceService) DeallocateSpace(ctx context.Context, req dto.DeallocateSpaceReq) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	floor := o.parkingLot.Floors[req.FloorNumber-1]
	if err := floor.DeallocateSlot(req.SpotNumber); err != nil {
		return err
	}

	fmt.Printf("Deallocated parking slspotot on Floor %d, Slot %d\n", req.FloorNumber, req.SpotNumber)

	if err := o.PublishSpotChangedEvent(ctx); err != nil {
		return err
	}

	return nil
}

func (o *spaceService) PublishSpotChangedEvent(ctx context.Context) error {
	// spotChangedInput := spaceEvents.SpotChangedInput{}

	// for _, floor := range o.parkingLot.Floors {
	// 	spots := []spaceEvents.Spot{}

	// 	for _, spot := range floor.Spots {
	// 		spots = append(spots, spaceEvents.Spot{
	// 			SpotNumber: spot.SpotNumber,
	// 			SpotType:   spot.SpotType,
	// 			IsOccupied: spot.IsOccupied,
	// 		})
	// 	}

	// 	spotChangedInput.Floors = append(spotChangedInput.Floors, spaceEvents.Floor{
	// 		FloorNumber: floor.FloorNumber,
	// 		Spots:       spots,
	// 	})
	// }

	// event := spaceEvents.CreateSpotChangedEvent(spotChangedInput)

	// if err := o.daprClient.PublishEvent(ctx, event); err != nil {
	// 	return fmt.Errorf("daprClient.PublishEvent: %w", err)
	// }

	return nil
}
