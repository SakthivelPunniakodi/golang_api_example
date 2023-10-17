// ---------------------------------------------------------------------------------------------------------------------
//  Author: Gayan Madushanka
//  Email: gayanmadushanka2@gmail.com
//  Created On: 4/9/2023
//  Purpose: <Small description about file>
// ---------------------------------------------------------------------------------------------------------------------

package handler

import (
	"github.com/go-chi/chi/v5"
)

func MapExternalDaprRoutes(mux *chi.Mux, handler Manager) {
	mux.Route("/api/v1", func(r chi.Router) {
		r.Post("/spot/allocate", handler.AllocateSpot)
		r.Get("/ticket", handler.GetTicketInfo)
		r.Post("/spot/deallocate", handler.DeallocateSpot)
	})
}

func MapExternalWebsocketRoutes(mux *chi.Mux, handler Manager) {
	mux.Route("/api/v1", func(r chi.Router) {
		r.Get("/spot/unoccupied", handler.NotifyUnoccupiedSpots)
	})
}
