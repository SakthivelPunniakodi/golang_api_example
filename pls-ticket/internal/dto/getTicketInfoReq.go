// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package dto

type GetTicketInfoReq struct {
	TicketID      string `json:"ticketId"`
	ValidationKey string `json:"validationKey"`
}
