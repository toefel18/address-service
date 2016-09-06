package api

import (
    "github.com/gin-gonic/gin"
    "github.com/toefel18/address-service/api/v1"
    "github.com/toefel18/address-service/api/v2"
    "github.com/toefel18/address-service/geodan"
    "github.com/toefel18/address-service/api/area"
)


func Publish(db geodan.AddressesNLDao) {
    areaStore := &area.AreaStore{}

    router := gin.Default()
    router.GET("/v1/addresses/:kixcode", v1.AddressByKixcode(db, areaStore))
    router.POST("/v1/addresses/search", v1.MultipleAddresses(db, areaStore))
    router.GET("/v1/excludeAreas/:kixcodeBeginIncl/:kixcodeEndIncl", v1.ExcludeArea(areaStore))
    router.PUT("/v1/excludeAreas/:kixcodeBeginIncl/:kixcodeEndIncl", v1.AddExcludeArea(areaStore))
    router.DELETE("/v1/excludeAreas/:kixcodeBeginIncl/:kixcodeEndIncl", v1.DeleteExcludeArea(areaStore))

    router.GET("/v2/addresses", v2.GetAllAddressesSummary(db))
    router.GET("/v2/addresses/:kixcode", v2.GetAddressByKixcode(db))
    router.Run(":8888")

}
