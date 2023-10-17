// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package main

import (
	"context"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
	daprclt "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/dapr/client"
	daprsvr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/dapr/server"
	ilogger "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/utils"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-space/internal/common"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-space/internal/env"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-space/internal/handler"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-space/internal/plsbestfit"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-space/internal/plsfirstfit"
)

func main() {
	logger := ilogger.Default(constant.Space)
	defer utils.CheckErr(logger, logger.Sync)

	logger.Info("service starting")

	env, err := env.ParseEnv()
	if err != nil {
		logger.Fatalf("env.ParseEnv: %v", err)
	}

	logger.Debugf("envs: %v", env)

	logger = ilogger.NewLogger(&ilogger.Config{
		ServiceName:    constant.Space,
		ServiceVersion: env.Service.Version,
		LogLevel:       env.Service.LogLevel,
	})

	daprClient, err := daprclt.NewClient(
		daprclt.WithLogger(logger),
		daprclt.WithPubSubAndTopic(constant.PubSubName, constant.Space),
	)
	if err != nil {
		logger.Fatalf("dapr.NewClient: %v", err)
	}
	defer daprClient.Close()

	service := getService(logger, daprClient, env.IsBestFit)

	registerHandlers(logger, service, env.Service.Port)
}

func getService(logger ilogger.Logger, daprClient daprclt.Client, isBestFit bool) common.SpaceService {
	floors := [][]constant.ParkingSpotType{
		{
			constant.ParkingSpotTypeCompact,
			constant.ParkingSpotTypeLarge,
			constant.ParkingSpotTypeHandicapped,
			constant.ParkingSpotTypeMotorcycle,
		},
		{
			constant.ParkingSpotTypeCompact,
			constant.ParkingSpotTypeLarge,
			constant.ParkingSpotTypeEVChargeStation,
			constant.ParkingSpotTypeCompact,
		},
	}

	floorCount := len(floors)
	floorCapacity := len(floors[0])

	if isBestFit {
		parkingLot := plsbestfit.NewParkingLot(floorCount, floorCapacity, floors...)

		service := plsbestfit.NewSpaceService(
			plsbestfit.WithLogger(logger),
			plsbestfit.WithDaprClient(daprClient),
			plsbestfit.WithParkingLot(parkingLot),
		)

		service.PublishSpotChangedEvent(context.Background())

		return service
	}

	parkingLot := plsfirstfit.NewParkingLot(floorCount, floorCapacity, floors...)

	service := plsfirstfit.NewSpaceService(
		plsfirstfit.WithLogger(logger),
		plsfirstfit.WithDaprClient(daprClient),
		plsfirstfit.WithParkingLot(parkingLot),
	)

	service.PublishSpotChangedEvent(context.Background())

	return service
}

func registerHandlers(logger ilogger.Logger, service common.SpaceService, port string) {
	h := handler.NewSpaceHandler(
		handler.WithLogger(logger),
		handler.WithSvc(service),
	)

	daprServer := daprsvr.NewServer(
		port,
		daprsvr.WithLogger(logger),
		daprsvr.WithPubSub(constant.PubSubName),
	)

	if err := handler.MapServiceInvocationRoutes(daprServer, h); err != nil {
		logger.Fatalf("handler.MapServiceInvocationRoutes: %v", err)
	}

	if err := daprServer.Start(); err != nil {
		logger.Fatalf("server.Start: %v", err)
	}
}
