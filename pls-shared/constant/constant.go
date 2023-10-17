// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package constant

// ---------------------------------------------------------------------------------------------------------------------

const (
	PubSubName = "pubsub"
)

// ---------------------------------------------------------------------------------------------------------------------
// Service names

const (
	Manager = "manager"
	Payment = "payment"
	Space   = "space"
	Ticket  = "ticket"
)

// ---------------------------------------------------------------------------------------------------------------------
// Invocation methods

// payment service
const (
	ProcessPayment = "ProcessPayment"
)

// space service
const (
	AllocateSpace   = "AllocateSpace"
	DeallocateSpace = "DeallocateSpace"
)

// ticket service
const (
	GenerateTicket = "GenerateTicket"
	CalculateFees  = "CalculateFees"
	DiscardTicket  = "DiscardTicket"
)

// ---------------------------------------------------------------------------------------------------------------------
// HTTP ContentTypes

const (
	ContentTypeJSON = "application/json; charset=utf-8"
)

// ---------------------------------------------------------------------------------------------------------------------
// HTTP Headers

const (
	OriginHeader                   = "Origin"
	ContentTypeHeader              = "Content-Type"
	AccessControlAllowOriginHeader = "Access-Control-Allow-Origin"
)

// ---------------------------------------------------------------------------------------------------------------------
// Vehicle Types

type VehicleType int

const (
	VehicleTypeMotorcycle VehicleType = 1 << iota
	VehicleTypeCar
	VehicleTypeVan
	VehicleTypeTruck
	AllVehicle = VehicleTypeMotorcycle | VehicleTypeCar | VehicleTypeVan | VehicleTypeTruck
)

func (vt VehicleType) HasFlag(flag VehicleType) bool {
	return vt&flag == flag
}

func (vt *VehicleType) AddFlag(flag VehicleType) {
	*vt |= flag
}

func (vt *VehicleType) RemoveFlag(flag VehicleType) {
	*vt &^= flag
}

// ---------------------------------------------------------------------------------------------------------------------
// Parking Spot Types

type ParkingSpotType string

const (
	ParkingSpotTypeCompact         ParkingSpotType = "Compact"
	ParkingSpotTypeLarge           ParkingSpotType = "Large"
	ParkingSpotTypeHandicapped     ParkingSpotType = "Handicapped"
	ParkingSpotTypeEVChargeStation ParkingSpotType = "EVChargeStation"
	ParkingSpotTypeMotorcycle      ParkingSpotType = "Motorcycle"
)

// ---------------------------------------------------------------------------------------------------------------------
