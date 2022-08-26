package service

import "unisun/api/class-room-price-mapping-processor-listener/src/models"

type ClassRoomPriceServicePort interface {
	GetClassRoomPrice(id int) *models.ClassRoomPrice
}
