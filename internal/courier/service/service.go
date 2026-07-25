package service

import (
	"errors"
	"math"
	"time"

	"github.com/maruki00/deligo/internal/courier/domain"
)

type CourierService struct {
	repo domain.CourierRepository
}

func NewCourierService(repo domain.CourierRepository) *CourierService {
	return &CourierService{repo: repo}
}

func (s *CourierService) RegisterCourier(id string, vehicle string, lat, lon float64, active bool) error {
	c := &domain.Courier{
		ID:               id,
		VehicleType:      vehicle,
		IsActive:         active,
		CurrentLatitude:  lat,
		CurrentLongitude: lon,
	}
	return s.repo.CreateCourier(c)
}

func (s *CourierService) HandleOrderConfirmed(orderID string, restaurantLat, restaurantLon float64) error {
	assignment := &domain.OrderCourier{
		OrderID: orderID,
		Status:  "searching",
	}

	if err := s.repo.CreateOrderCourier(assignment); err != nil {
		return err
	}

	couriers, err := s.repo.GetActiveCouriers()
	if err != nil || len(couriers) == 0 {
		return nil
	}

	var bestCourier *domain.Courier
	minDistance := math.MaxFloat64

	for i := range couriers {
		busy, err := s.repo.GetCourierBusyStatus(couriers[i].ID)
		if err != nil || busy {
			continue
		}

		dist := haversine(couriers[i].CurrentLatitude, couriers[i].CurrentLongitude, restaurantLat, restaurantLon)
		if dist < minDistance {
			minDistance = dist
			bestCourier = &couriers[i]
		}
	}

	if bestCourier != nil {
		now := time.Now()
		assignment.CourierID = &bestCourier.ID
		assignment.AssignedAt = &now
		return s.repo.UpdateOrderCourier(assignment)
	}

	return nil
}

func (s *CourierService) AcceptOrder(orderID string, courierID string) error {
	assignment, err := s.repo.GetOrderCourierByOrder(orderID)
	if err != nil {
		return err
	}
	if assignment.CourierID == nil || *assignment.CourierID != courierID {
		return errors.New("unauthorized state: order not assigned to this courier")
	}
	if assignment.Status != "searching" {
		return errors.New("invalid lifecycle status transition to accepted")
	}

	assignment.Status = "accepted"
	return s.repo.UpdateOrderCourier(assignment)
}

func (s *CourierService) ArriveAtRestaurant(orderID string, courierID string) error {
	assignment, err := s.repo.GetOrderCourierByOrder(orderID)
	if err != nil {
		return err
	}
	if assignment.CourierID == nil || *assignment.CourierID != courierID {
		return errors.New("unauthorized state")
	}
	if assignment.Status != "accepted" {
		return errors.New("invalid transition to at_restaurant")
	}

	assignment.Status = "at_restaurant"
	return s.repo.UpdateOrderCourier(assignment)
}

func (s *CourierService) StartCourier(orderID string, courierID string) error {
	assignment, err := s.repo.GetOrderCourierByOrder(orderID)
	if err != nil {
		return err
	}
	if assignment.CourierID == nil || *assignment.CourierID != courierID {
		return errors.New("unauthorized state")
	}
	if assignment.Status != "at_restaurant" {
		return errors.New("invalid transition to picked_up")
	}

	now := time.Now()
	assignment.Status = "picked_up"
	assignment.PickedUpAt = &now
	return s.repo.UpdateOrderCourier(assignment)
}

func (s *CourierService) CompleteCourier(orderID string, courierID string) error {
	assignment, err := s.repo.GetOrderCourierByOrder(orderID)
	if err != nil {
		return err
	}
	if assignment.CourierID == nil || *assignment.CourierID != courierID {
		return errors.New("unauthorized state")
	}
	if assignment.Status != "picked_up" {
		return errors.New("invalid transition to delivered")
	}

	now := time.Now()
	assignment.Status = "delivered"
	assignment.DeliveredAt = &now
	return s.repo.UpdateOrderCourier(assignment)
}

func (s *CourierService) TrackLocation(courierID string, orderID string, lat, lon float64) error {
	if err := s.repo.UpdateCourierLocation(courierID, lat, lon); err != nil {
		return err
	}

	if orderID != "" {
		trackingLog := &domain.OrderTracking{
			OrderID:   orderID,
			Latitude:  lat,
			Longitude: lon,
		}
		return s.repo.SaveTrackingLog(trackingLog)
	}
	return nil
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371.0
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0

	l1 := lat1 * math.Pi / 180.0
	l2 := lat2 * math.Pi / 180.0

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(l1)*math.Cos(l2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
