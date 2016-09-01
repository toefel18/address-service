package api

import (
	"github.com/gin-gonic/gin"
	"github.com/toefel18/address-service/api/v1"
	"github.com/toefel18/address-service/geodan"
)

func Publish(db *geodan.CsvExtract) {
	router := gin.Default()
	router.GET( "/v1/addresses/:kixcode", v1.AddressByKixcode(db))
	router.POST("/v1/addresses/search", v1.MultipleAddresses(db))
	router.Run(":8888")

}
