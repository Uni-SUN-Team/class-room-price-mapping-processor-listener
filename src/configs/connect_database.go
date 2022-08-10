package config

import (
	"log"
	"os"
	"strings"
	"unisun/api/class-room-price-mapping-processor-schedule/src/constants"
	"unisun/api/class-room-price-mapping-processor-schedule/src/entitys"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	str := []string{
		"host=" + os.Getenv(constants.DB_HOST),
		"user=" + os.Getenv(constants.DB_USER),
		"password=" + os.Getenv(constants.DB_PASS),
		"dbname=" + os.Getenv(constants.DB_NAME),
		"port=" + os.Getenv(constants.DB_PORT),
		"sslmode=" + os.Getenv(constants.DB_SSL),
		"TimeZone=" + os.Getenv(constants.DB_TIMEZONE),
	}
	dsn := strings.Join(str, " ")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database!")
	}

	var ClassRoomPriceEntity entitys.ClassRoomPriceEntity
	database.AutoMigrate(&ClassRoomPriceEntity)
	DB = database
}
