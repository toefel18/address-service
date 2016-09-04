package api

import (
	"github.com/gin-gonic/gin"
	"github.com/toefel18/address-service/api/v1"
	"github.com/toefel18/address-service/geodan"
	"github.com/toefel18/address-service/api/v2"
)

func Publish(db geodan.AddressesNLDao) {
	router := gin.Default()
	router.GET( "/v1/addresses/:kixcode", v1.AddressByKixcode(db))
	router.POST("/v1/addresses/search", v1.MultipleAddresses(db))
	router.GET("/v2/addresses", v2.GetAllAddressesSummary(db))
	router.GET("/v2/addresses/:kixcode", v2.GetAddressByKixcode(db))
	router.Run(":8888")

}
