package model

import "github.com/google/uuid"

type OrderItem struct {
	ID         uuid.UUID `json:"id"`
	Product    Product   `json:"product"`
	Status     int       `json:"status"`
	GroupID    uuid.UUID `json:"group_id"`
	CreateDate float32   `json:"create_date"`
	Count      int       `json:"count"`
}
