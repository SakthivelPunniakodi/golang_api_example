// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package dto

import "github.com/shopspring/decimal"

type ProcessPaymentReq struct {
	TicketID      string          `json:"ticketId"`
	Amount        decimal.Decimal `json:"amount"`
	CVV           string          `json:"cvv"`
	CardNumber    string          `json:"creditCardNumber"`
	CardType      string          `json:"cardType"`
	PaymentMethod string          `json:"paymentMethod"`
}
