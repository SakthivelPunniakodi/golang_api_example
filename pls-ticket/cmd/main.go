// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package main

import (
	"sampleapi/pls-shared/constant"
	daprclt "sampleapi/pls-shared/dapr/client"
	daprsvr "sampleapi/pls-shared/dapr/server"
	ilogger "sampleapi/pls-shared/logger"
	"sampleapi/pls-shared/utils"
	"sampleapi/pls-ticket/internal/env"
	"sampleapi/pls-ticket/internal/handler"
	iservice "sampleapi/pls-ticket/internal/service"
)

func main() {
	logger := ilogger.Default(constant.Ticket)
	defer utils.CheckErr(logger, logger.Sync)

	logger.Info("service starting")

	env, err := env.ParseEnv()
	if err != nil {
		logger.Fatalf("env.ParseEnv: %v", err)
	}

	logger.Debugf("envs: %v", env)

	logger = ilogger.NewLogger(&ilogger.Config{
		ServiceName:    constant.Ticket,
		ServiceVersion: env.Service.Version,
		LogLevel:       env.Service.LogLevel,
	})

	daprClient, err := daprclt.NewClient(
		daprclt.WithLogger(logger),
		daprclt.WithPubSubAndTopic(constant.PubSubName, constant.Ticket),
	)
	if err != nil {
		logger.Fatalf("dapr.NewClient: %v", err)
	}
	defer daprClient.Close()

	service := getService(logger, daprClient)

	registerHandlers(logger, service, env.Service.Port)
}

func getService(logger ilogger.Logger, daprClient daprclt.Client) iservice.TicketService {
	return iservice.NewTicketService(
		iservice.WithLogger(logger),
		iservice.WithDaprClient(daprClient),
	)
}

func registerHandlers(logger ilogger.Logger, service iservice.TicketService, port string) {
	h := handler.NewTicketHandler(
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