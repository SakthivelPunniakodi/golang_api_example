// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package main

import (
	"github.com/SakthivelPunniakodi/golang_api_example/pls-payment/internal/env"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-payment/internal/handler"
	iservice "github.com/SakthivelPunniakodi/golang_api_example/pls-payment/internal/service"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
	daprclt "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/dapr/client"
	daprsvr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/dapr/server"
	ilogger "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/utils"
)

func main() {
	logger := ilogger.Default(constant.Payment)
	defer utils.CheckErr(logger, logger.Sync)

	logger.Info("service starting")

	env, err := env.ParseEnv()
	if err != nil {
		logger.Fatalf("env.ParseEnv: %v", err)
	}

	logger.Debugf("envs: %v", env)

	logger = ilogger.NewLogger(&ilogger.Config{
		ServiceName:    constant.Payment,
		ServiceVersion: env.Service.Version,
		LogLevel:       env.Service.LogLevel,
	})

	daprClient, err := daprclt.NewClient(
		daprclt.WithLogger(logger),
		daprclt.WithPubSubAndTopic(constant.PubSubName, constant.Payment),
	)
	if err != nil {
		logger.Fatalf("dapr.NewClient: %v", err)
	}
	defer daprClient.Close()

	service := getService(logger, daprClient)

	registerHandlers(logger, service, env.Service.Port)
}

func getService(logger ilogger.Logger, daprClient daprclt.Client) iservice.PaymentService {
	return iservice.NewPaymentService(
		iservice.WithLogger(logger),
		iservice.WithDaprClient(daprClient),
	)
}

func registerHandlers(logger ilogger.Logger, service iservice.PaymentService, port string) {
	h := handler.NewPaymentHandler(
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
