// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package main

import (
	"context"

	"sampleapi/pls-shared/constant"
	daprclt "sampleapi/pls-shared/dapr/client"
	daprsvr "sampleapi/pls-shared/dapr/server"
	ilogger "sampleapi/pls-shared/logger"
	"sampleapi/pls-shared/utils"
	"sampleapi/pls-space/internal/common"
	"sampleapi/pls-space/internal/env"
	"sampleapi/pls-space/internal/handler"
	"sampleapi/pls-space/internal/plsbestfit"
	"sampleapi/pls-space/internal/plsfirstfit"
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
