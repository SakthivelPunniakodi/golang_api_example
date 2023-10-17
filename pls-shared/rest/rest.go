// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	ierr "github.com/SakthivelPunniakodi/golang_api_example/pls-shared/errors"
	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/logger"
)

type errorResponse struct {
	Error string `json:"error"`
}

func ErrorRes(httpResWtr http.ResponseWriter, logger logger.Logger, msg string, err error) {
	logger.Errorf("%s: %v", msg, err)

	errRes := errorResponse{
		Error: err.Error(),
	}

	var internalErr *ierr.Error
	if !errors.As(err, &internalErr) {
		Res(httpResWtr, http.StatusInternalServerError, errRes)

		return
	}

	switch internalErr.Code() {
	case ierr.NotFound:
		Res(httpResWtr, http.StatusNotFound, errRes)
	case ierr.InvalidArgument:
		Res(httpResWtr, http.StatusBadRequest, errRes)
	case ierr.Unknown:
		fallthrough
	default:
		Res(httpResWtr, http.StatusInternalServerError, errRes)
	}
}

func Res(httpResWtr http.ResponseWriter, statusCode int, res ...any) {
	if res == nil {
		httpResWtr.WriteHeader(statusCode)

		return
	}

	// TEMP
	httpResWtr.WriteHeader(statusCode)
	json.NewEncoder(httpResWtr).Encode(res[0])

	return
}
