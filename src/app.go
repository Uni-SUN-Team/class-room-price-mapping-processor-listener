package src

import (
	config "unisun/api/class-room-price-mapping-processor-listener/src/configs"
	"unisun/api/class-room-price-mapping-processor-listener/src/controllers"
	"unisun/api/class-room-price-mapping-processor-listener/src/repositories"
	"unisun/api/class-room-price-mapping-processor-listener/src/routes"
	"unisun/api/class-room-price-mapping-processor-listener/src/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func App() *gin.Engine {
	config.ConnectDatabase()
	r := gin.Default()
	g := r.Group(viper.GetString("app.context_path") + viper.GetString("app.root_path") + "/v1")
	{
		handleClassRoomPrice().ClassRoomPriceRoute(g)
	}
	return r
}

func handleClassRoomPrice() *routes.ClassRoomPriceRouteAdapter {
	repo := repositories.NewClassRoomPriceRepositoriesAdapter()
	service := services.NewClassRoomPriceServiceAdapter(repo)
	controller := controllers.NewClassRoomPriceControllerAdapter(service)
	router := routes.NewClassRoomPriceRouteAdapter(controller)
	return router
}
