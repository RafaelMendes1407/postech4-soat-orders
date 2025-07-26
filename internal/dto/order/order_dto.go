package dto

import (
	"time"
)

type OrderStatus string

const (
	OrderStatusReceived  OrderStatus = "received"
	OrderStatusPreparing OrderStatus = "preparing"
	OrderStatusReady     OrderStatus = "ready"
	OrderStatusCompleted OrderStatus = "completed"
)

type OrderDTO struct {
	Id        string
	Number    int
	Status    string
	ClientId  string
	PaymentId string
	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
