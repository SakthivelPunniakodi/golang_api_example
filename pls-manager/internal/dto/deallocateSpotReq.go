// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package dto

type DeallocateSpotReq struct {
	TicketID      string `json:"ticketId"`
	ValidationKey string `json:"validationKey"`
	CVV           string `json:"cvv"`
	CardNumber    string `json:"cardNumber"`
	CardType      string `json:"cardType"`
	PaymentMethod string `json:"paymentMethod"`
}
