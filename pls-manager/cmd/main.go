// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"sampleapi/pls-manager/internal/env"
	"sampleapi/pls-manager/internal/handler"
	"sampleapi/pls-manager/internal/payment"
	iservice "sampleapi/pls-manager/internal/service"
	"sampleapi/pls-manager/internal/space"
	"sampleapi/pls-manager/internal/ticket"
	"sampleapi/pls-shared/constant"
	daprclt "sampleapi/pls-shared/dapr/client"
	daprsvr "sampleapi/pls-shared/dapr/server"
	spaceEvents "sampleapi/pls-shared/events/space"
	ilogger "sampleapi/pls-shared/logger"
	"sampleapi/pls-shared/utils"
)

func main() {
	logger := ilogger.Default(constant.Manager)
	defer utils.CheckErr(logger, logger.Sync)

	logger.Info("service starting")

	env, err := env.ParseEnv()
	if err != nil {
		logger.Fatalf("env.ParseEnv: %+v", err)
	}

	logger.Debugf("envs: %v", env)

	logger = ilogger.NewLogger(&ilogger.Config{
		ServiceName:    constant.Manager,
		ServiceVersion: env.Service.Version,
		LogLevel:       env.Service.LogLevel,
	})

	daprClient, err := daprclt.NewClient(
		daprclt.WithLogger(logger),
		daprclt.WithPubSubAndTopic(constant.PubSubName, constant.Manager),
	)
	if err != nil {
		logger.Fatalf("dapr.NewClient: %v", err)
	}
	defer daprClient.Close()

	unoccupiedSpotsCh := make(chan []spaceEvents.Floor)

	service := getService(logger, daprClient, unoccupiedSpotsCh)
	registerHandlers(logger, service, env.Service.Port, unoccupiedSpotsCh)
}

func getService(logger ilogger.Logger, daprClient daprclt.Client, unoccupiedSpotsCh chan<- []spaceEvents.Floor) iservice.Manager {
	paymentSvc := payment.NewPayment(logger, daprClient)
	spaceSvc := space.NewSpace(logger, daprClient)
	ticketSvc := ticket.NewTicket(logger, daprClient)

	return iservice.NewManager(
		iservice.WithLogger(logger),
		iservice.WithDaprClient(daprClient),
		iservice.WithPaymentSvc(paymentSvc),
		iservice.WithSpaceSvc(spaceSvc),
		iservice.WithTicketSvc(ticketSvc),
		iservice.WithUnoccupiedSpotsChannel(unoccupiedSpotsCh),
	)
}

func registerHandlers(logger ilogger.Logger, service iservice.Manager, port string, unoccupiedSpotsCh <-chan []spaceEvents.Floor) {
	h := handler.NewManager(
		handler.WithLogger(logger),
		handler.WithSvc(service),
		handler.WithUnoccupiedSpotsChannel(unoccupiedSpotsCh),
	)

	websocketRouter := chi.NewRouter()

	handler.MapExternalWebsocketRoutes(websocketRouter, h)

	go func() {
		err := http.ListenAndServe(":8090", websocketRouter)
		if err != nil {
			logger.Fatalf("http.ListenAndServe: %v", err)
		}
	}()

	daprRouter := chi.NewRouter()

	handler.MapExternalDaprRoutes(daprRouter, h)

	daprServer := daprsvr.NewServerWithMux(
		daprRouter,
		port,
		daprsvr.WithLogger(logger),
		daprsvr.WithPubSub(constant.PubSubName),
	)

	if err := handler.MapTopicEventHandlerRoutes(daprServer, h); err != nil {
		logger.Fatalf("handler.MapTopicEventHandlerRoutes: %v", err)
	}

	if err := daprServer.Start(); err != nil {
		logger.Fatalf("daprServer.Start: %v", err)
	}
}
