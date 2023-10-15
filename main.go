package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ParkingSpace struct {
	ID            int    `json:"id"`
	Level         int    `json:"level"`
	Type          string `json:"type"`
	Occupied      bool   `json:"occupied"`
	Reserved      bool   `json:"reserver"`
	ReservedBy    string `json:"ReservedBy"`
	ReservationID string `json:"ReservationID`
}

var parkingSpaces []ParkingSpace

func main() {
	router := gin.Default()

	// Initialize some parking spaces
	parkingSpaces = []ParkingSpace{
		{ID: 1, Level: 1, Type: "Compact", Occupied: false, Reserved: false},
		{ID: 2, Level: 1, Type: "Large", Occupied: false, Reserved: false},
		{ID: 3, Level: 2, Type: "Compact", Occupied: false, Reserved: false},
		{ID: 4, Level: 2, Type: "Motorcycle", Occupied: false, Reserved: false},
	}

	// API routes
	router.GET("/parking-spaces", GetParkingSpaces)
	router.POST("/park", ParkCar)
	router.GET("/unpark", unParkCar)
	router.POST("/reserve-park", ParkReserve)
	// Start the server
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func GetParkingSpaces(c *gin.Context) {
	c.JSON(http.StatusOK, parkingSpaces)
}

func unParkCar(c *gin.Context) {
	SpaceID := c.DefaultQuery("space_id", "")
	i, err := strconv.Atoi(SpaceID)
	space, found := findParkingSpace(i)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parking space not found"})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parking space not found"})
		return
	}
	if !space.Occupied {
		c.JSON(http.StatusConflict, gin.H{"error": "Parking space is already unparked"})
		return
	}
	space.Occupied = false
	space.ReservationID = ""
	space.Reserved = false
	space.ReservedBy = ""
	c.JSON(http.StatusOK, gin.H{"message": "Car unparked successfully"})
}

func ParkCar(c *gin.Context) {
	var req struct {
		SpaceID       int    `json:"space_id"`
		ReservationID string `json:"reservation_id"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	space, found := findParkingSpace(req.SpaceID)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parking space not found"})
		return
	}

	if space.Occupied {
		c.JSON(http.StatusConflict, gin.H{"error": "Parking space is already occupied"})
		return
	}

	if space.ReservationID != req.ReservationID && space.ReservationID != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "Parking space is already reserved by other."})
		return
	}

	space.Occupied = true
	c.JSON(http.StatusOK, gin.H{"message": "Car parked successfully"})
}

func ParkReserve(c *gin.Context) {
	var req struct {
		SpaceID    int    `json:"space_id"`
		ReservedBy string `json:"recerved_by"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	space, found := findParkingSpace(req.SpaceID)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parking space not found"})
		return
	}

	if space.Occupied {
		c.JSON(http.StatusConflict, gin.H{"error": "Parking space is already occupied"})
		return
	}
	if space.Reserved {
		c.JSON(http.StatusConflict, gin.H{"error": "Parking space is already reserved"})
		return
	}
	space.Reserved = true
	space.ReservedBy = req.ReservedBy
	uuid := uuid.New()
	space.ReservationID = uuid.String()
	c.JSON(http.StatusOK, gin.H{"message": "Parking reserved successfully", "recervation_id": space.ReservationID})
}

func findParkingSpace(spaceID int) (*ParkingSpace, bool) {
	for i, space := range parkingSpaces {
		if space.ID == spaceID {
			return &parkingSpaces[i], true
		}
	}
	return nil, false
}
