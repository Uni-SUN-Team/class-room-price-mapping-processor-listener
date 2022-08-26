package controllers

import (
	"net/http"
	"strconv"
	"unisun/api/class-room-price-mapping-processor-listener/src/models"
	"unisun/api/class-room-price-mapping-processor-listener/src/ports/service"

	"github.com/gin-gonic/gin"
)

type ClassRoomPriceControllerAdapter struct {
	ClassRoomPriceService service.ClassRoomPriceServicePort
}

func NewClassRoomPriceControllerAdapter(classRoomPriceService service.ClassRoomPriceServicePort) *ClassRoomPriceControllerAdapter {
	return &ClassRoomPriceControllerAdapter{
		ClassRoomPriceService: classRoomPriceService,
	}
}

func (srv *ClassRoomPriceControllerAdapter) GetClassRoomPriceById(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
			Code  int    `json:"code"`
		}{
			Error: err.Error(),
			Code:  http.StatusInternalServerError,
		})
	}
	data := srv.ClassRoomPriceService.GetClassRoomPrice(id)
	c.AbortWithStatusJSON(http.StatusOK, struct {
		Data models.ClassRoomPrice `json:"data"`
	}{
		Data: *data,
	})
}
