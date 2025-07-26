package dto

import (
	"time"
)

type PaymentDTO struct {
	Id        string
	Provider  string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
