// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	"sampleapi/pls-shared/constant"
	daprsvr "sampleapi/pls-shared/dapr/server"
)

func MapServiceInvocationRoutes(server daprsvr.Server, handler PaymentHandler) error {
	if err := server.AddServiceInvocationHandler(constant.ProcessPayment, handler.ProcessPayment); err != nil {
		return err
	}

	return nil
}
