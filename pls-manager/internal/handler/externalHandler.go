// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-manager/internal/dto"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
	ierr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/errors"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/rest"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/utils"
)

type externalHandler interface {
	AllocateSpot(httpResWtr http.ResponseWriter, httpReq *http.Request)
	GetTicketInfo(httpResWtr http.ResponseWriter, httpReq *http.Request)
	DeallocateSpot(httpResWtr http.ResponseWriter, httpReq *http.Request)
	NotifyUnoccupiedSpots(httpResWtr http.ResponseWriter, httpReq *http.Request)
}

func (o manager) AllocateSpot(httpResWtr http.ResponseWriter, httpReq *http.Request) {
	req, err := validateAndGetAllocateSpotReq(httpReq)
	if err != nil {
		rest.ErrorRes(httpResWtr, o.logger, "Invalid request.", err)

		return
	}

	res, err := o.managerSvc.AllocateSpot(httpReq.Context(), req)
	if err != nil {
		rest.ErrorRes(httpResWtr, o.logger, "Spot allocation failed.", err)

		return
	}

	rest.Res(httpResWtr, http.StatusCreated, res)
}

func (o manager) GetTicketInfo(httpResWtr http.ResponseWriter, httpReq *http.Request) {
	req, err := validateAndGetTicketInfoReq(httpReq)
	if err != nil {
		rest.ErrorRes(httpResWtr, o.logger, "Invalid request.", err)

		return
	}

	res, err := o.managerSvc.GetTicketInfo(httpReq.Context(), req)
	if err != nil {
		rest.ErrorRes(httpResWtr, o.logger, "Get fees failed.", err)

		return
	}

	rest.Res(httpResWtr, http.StatusOK, res)
}

func (o manager) DeallocateSpot(httpResWtr http.ResponseWriter, httpReq *http.Request) {
	req, err := validateAndGetDeallocateSpotReq(httpReq)
	if err != nil {
		rest.ErrorRes(httpResWtr, o.logger, "Invalid request.", err)

		return
	}

	if err := o.managerSvc.DeallocateSpot(httpReq.Context(), req); err != nil {
		rest.ErrorRes(httpResWtr, o.logger, "Spot deallocation failed.", err)

		return
	}

	rest.Res(httpResWtr, http.StatusCreated)
}

func (o manager) NotifyUnoccupiedSpots(httpResWtr http.ResponseWriter, httpReq *http.Request) {
	conn, err := o.websocketUpgrader.Upgrade(httpResWtr, httpReq, nil)
	if err != nil {
		rest.ErrorRes(httpResWtr, o.logger, "websocketUpgrader.Upgrade:", err)

		return
	}
	defer conn.Close()

	for {
		s := <-o.unoccupiedSpotsCh

		jsonData, err := json.Marshal(s)
		if err != nil {
			log.Println(err)
			return
		}

		if err := conn.WriteMessage(websocket.TextMessage, jsonData); err != nil {
			rest.ErrorRes(httpResWtr, o.logger, "conn.WriteMessage", err)

			return
		}
	}
}

func validateAndGetAllocateSpotReq(httpReq *http.Request) (dto.AllocateSpotReq, error) {
	req := dto.AllocateSpotReq{}

	decoder := json.NewDecoder(httpReq.Body)
	if err := decoder.Decode(&req); err != nil {
		return req, ierr.WrapErrorf(err, ierr.InvalidArgument, "decoder.Decode")
	}

	if len(req.VehicleNumber) == 0 {
		return req, ierr.NewErrorf(ierr.InvalidArgument, "Vehicle number is empty")
	}

	if !utils.IsValidVehicleNumber(req.VehicleNumber) {
		return req, ierr.NewErrorf(ierr.InvalidArgument, "Vehicle number is invalid")
	}

	if len(req.VehicleType) == 0 {
		return req, ierr.NewErrorf(ierr.InvalidArgument, "Vehicle type is empty")
	}

	req.VehicleTypeEnum = utils.GetVehicleType(req.VehicleType)

	if !constant.AllVehicle.HasFlag(req.VehicleTypeEnum) {
		return req, ierr.NewErrorf(ierr.InvalidArgument, "Vehicle type is invalid")
	}

	if req.IsNeedHandicappedSpot && req.IsNeedEVChargeStationSpot {
		return req, ierr.NewErrorf(ierr.InvalidArgument, "You can only select one of Handicapped or EVChargeStation spot")
	}

	return req, nil
}

func validateAndGetTicketInfoReq(httpReq *http.Request) (dto.GetTicketInfoReq, error) {
	req := dto.GetTicketInfoReq{}

	queryStrings := httpReq.URL.Query()

	ticketID := queryStrings.Get("ticketId")

	if err := validateTicketID(ticketID); err != nil {
		return req, err
	}

	req.TicketID = ticketID

	validationKey := queryStrings.Get("validationKey")

	if err := validateValidationKey(validationKey); err != nil {
		return req, err
	}

	req.ValidationKey = validationKey

	return req, nil
}

func validateAndGetDeallocateSpotReq(httpReq *http.Request) (dto.DeallocateSpotReq, error) {
	req := dto.DeallocateSpotReq{}

	decoder := json.NewDecoder(httpReq.Body)
	if err := decoder.Decode(&req); err != nil {
		return req, ierr.WrapErrorf(err, ierr.InvalidArgument, "decoder.Decode")
	}

	if err := validateTicketID(req.TicketID); err != nil {
		return req, err
	}

	if err := validateValidationKey(req.ValidationKey); err != nil {
		return req, err
	}

	// TODO validate card info

	return req, nil
}

func validateTicketID(ticketID string) error {
	if len(ticketID) == 0 {
		return ierr.NewErrorf(ierr.InvalidArgument, "Ticket id is empty")
	}

	if !utils.IsValidTicketFormat(ticketID) {
		return ierr.NewErrorf(ierr.InvalidArgument, "Ticket id is invalid")
	}

	return nil
}

func validateValidationKey(validationKey string) error {
	if len(validationKey) == 0 {
		return ierr.NewErrorf(ierr.InvalidArgument, "Validation key is empty")
	}

	if !utils.IsValidUUID(validationKey) {
		return ierr.NewErrorf(ierr.InvalidArgument, "Validation key is invalid")
	}

	return nil
}
