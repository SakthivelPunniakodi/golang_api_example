package utils

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strconv"

	"github.com/google/uuid"

	"github.com/SakthivelPunniakodi/golang_api_example/pls-shared/constant"
)

func GetVehicleTypeString(vehicleType constant.VehicleType) string {
	switch vehicleType {
	case constant.VehicleTypeMotorcycle:
		return "Motorcycle"
	case constant.VehicleTypeCar:
		return "Car"
	case constant.VehicleTypeVan:
		return "Van"
	case constant.VehicleTypeTruck:
		return "Truck"
	default:
		return "Unknown"
	}
}

func GetVehicleType(vehicleType string) constant.VehicleType {
	switch vehicleType {
	case "Motorcycle":
		return constant.VehicleTypeMotorcycle
	case "Car":
		return constant.VehicleTypeCar
	case "Van":
		return constant.VehicleTypeVan
	case "Truck":
		return constant.VehicleTypeTruck
	}

	return 0
}

func IsValidVehicleNumber(vehicleNumber string) bool {
	pattern := `^[A-Z]{3}-\d{4}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(vehicleNumber)
}

func GetParkingSpotType(vehicleType constant.VehicleType, isNeedHandicappedSpot, isNeedEVChargeStationSpot bool) (constant.ParkingSpotType, error) {
	if isNeedHandicappedSpot {
		return constant.ParkingSpotTypeHandicapped, nil
	}

	if isNeedEVChargeStationSpot {
		return constant.ParkingSpotTypeEVChargeStation, nil
	}

	switch vehicleType {
	case constant.VehicleTypeMotorcycle:
		return constant.ParkingSpotTypeMotorcycle, nil
	case constant.VehicleTypeCar, constant.VehicleTypeVan:
		return constant.ParkingSpotTypeCompact, nil
	case constant.VehicleTypeTruck:
		return constant.ParkingSpotTypeLarge, nil
	}

	return "", fmt.Errorf("")
}

func HashUint64ToUUID(inputUint64 uint64) (uuid.UUID, error) {
	intBytes := []byte(strconv.FormatUint(inputUint64, 10))

	hashBytes := sha256.Sum256(intBytes)

	uuidBytes := hashBytes[:16]

	u, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		return uuid.Nil, err
	}

	return u, nil
}

func ExtractNumber(input string) (uint64, error) {
	pattern := `(\d+)`

	regex := regexp.MustCompile(pattern)

	match := regex.FindStringSubmatch(input)

	if len(match) >= 2 {
		number, err := strconv.ParseUint(match[1], 10, 64)
		if err != nil {
			return 0, err
		}
		return number, nil
	}

	return 0, fmt.Errorf("Number not found in input")
}

func IsValidUUID(input string) bool {
	_, err := uuid.Parse(input)
	return err == nil
}

func IsValidTicketFormat(input string) bool {
	pattern := `^TICKET-\d+$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)
}
