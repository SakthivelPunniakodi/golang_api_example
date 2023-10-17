// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
	daprsvr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/dapr/server"
)

func MapServiceInvocationRoutes(server daprsvr.Server, handler TicketHandler) error {
	if err := server.AddServiceInvocationHandler(constant.GenerateTicket, handler.GenerateTicket); err != nil {
		return err
	}

	if err := server.AddServiceInvocationHandler(constant.CalculateFees, handler.GetTicketInfo); err != nil {
		return err
	}

	if err := server.AddServiceInvocationHandler(constant.DiscardTicket, handler.DiscardTicket); err != nil {
		return err
	}

	return nil
}
