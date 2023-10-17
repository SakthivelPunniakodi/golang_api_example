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

func MapServiceInvocationRoutes(server daprsvr.Server, handler SpaceHandler) error {
	if err := server.AddServiceInvocationHandler(constant.AllocateSpace, handler.AllocateSpace); err != nil {
		return err
	}

	if err := server.AddServiceInvocationHandler(constant.DeallocateSpace, handler.DeallocateSpace); err != nil {
		return err
	}

	return nil
}
