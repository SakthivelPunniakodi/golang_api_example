// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package plsbestfit

import (
	"context"
	"fmt"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
	daprclt "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/dapr/client"
	spaceEvents "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/events/space"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-space/internal/common"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-space/internal/dto"
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
	floorNumber, spotNumber, err := o.parkingLot.AllocateSpot(req.SpotType)
	if err != nil {
		return dto.AllocateSpaceRes{}, err
	}

	fmt.Printf("Allocated parking spot on Floor %d, Spot %d, Type: %s\n", floorNumber, spotNumber, req.SpotType)

	if err := o.PublishSpotChangedEvent(ctx); err != nil {
		return dto.AllocateSpaceRes{}, err
	}

	return dto.AllocateSpaceRes{
		FloorNumber: floorNumber,
		SpotNumber:  spotNumber,
		SpotType:    req.SpotType,
	}, nil

}

func (o *spaceService) DeallocateSpace(ctx context.Context, req dto.DeallocateSpaceReq) error {
	if err := o.parkingLot.DeallocateSpot(req.FloorNumber, req.SpotNumber, req.SpotType); err != nil {
		return err
	}

	fmt.Printf("Deallocated parking slspotot on Floor %d, Slot %d\n", req.FloorNumber, req.SpotNumber)

	if err := o.PublishSpotChangedEvent(ctx); err != nil {
		return err
	}

	return nil
}

func (o *spaceService) PublishSpotChangedEvent(ctx context.Context) error {
	spotChangedInput := spaceEvents.SpotChangedInput{}

	for _, floor := range o.parkingLot.Floors {
		spots := map[constant.ParkingSpotType][]spaceEvents.Spot{}

		for _, spot := range floor.CompactAvailableSpots {
			spots[constant.ParkingSpotTypeCompact] = append(spots[constant.ParkingSpotTypeCompact], spaceEvents.Spot{
				SpotNumber: spot.SpotNumber,
			})
		}

		for _, spot := range floor.LargeAvailableSpots {
			spots[constant.ParkingSpotTypeLarge] = append(spots[constant.ParkingSpotTypeLarge], spaceEvents.Spot{
				SpotNumber: spot.SpotNumber,
			})
		}

		for _, spot := range floor.HandicappedAvailableSpots {
			spots[constant.ParkingSpotTypeHandicapped] = append(spots[constant.ParkingSpotTypeHandicapped], spaceEvents.Spot{
				SpotNumber: spot.SpotNumber,
			})
		}

		for _, spot := range floor.EVChargeStationAvailableSpots {
			spots[constant.ParkingSpotTypeEVChargeStation] = append(spots[constant.ParkingSpotTypeEVChargeStation], spaceEvents.Spot{
				SpotNumber: spot.SpotNumber,
			})
		}

		for _, spot := range floor.MotorcycleAvailableSpots {
			spots[constant.ParkingSpotTypeMotorcycle] = append(spots[constant.ParkingSpotTypeMotorcycle], spaceEvents.Spot{
				SpotNumber: spot.SpotNumber,
			})
		}

		spotChangedInput.Floors = append(spotChangedInput.Floors, spaceEvents.Floor{
			FloorNumber: floor.FloorNumber,
			SpotTypes:   spots,
		})
	}

	event := spaceEvents.CreateSpotChangedEvent(spotChangedInput)

	if err := o.daprClient.PublishEvent(ctx, event); err != nil {
		return fmt.Errorf("daprClient.PublishEvent: %w", err)
	}

	return nil
}
