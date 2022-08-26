package services

import (
	"unisun/api/class-room-price-mapping-processor-listener/src/models"
	"unisun/api/class-room-price-mapping-processor-listener/src/ports/repository"
)

type ClassRoomPriceServiceAdapter struct {
	Repo repository.ClassRoomPricePort
}

func NewClassRoomPriceServiceAdapter(repo repository.ClassRoomPricePort) *ClassRoomPriceServiceAdapter {
	return &ClassRoomPriceServiceAdapter{
		Repo: repo,
	}
}

func (srv *ClassRoomPriceServiceAdapter) GetClassRoomPrice(id int) *models.ClassRoomPrice {
	result := srv.Repo.GetByClassRoomId(id)
	return &models.ClassRoomPrice{
		Id:           result.Id,
		ClassRoomId:  result.ClassRoomId,
		RegularPrice: result.RegularPrice,
		SpecialPrice: result.SpecialPrice,
		Advisors:     result.Advisors,
		Categories:   result.Categories,
	}
}
