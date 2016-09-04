package v2

import (
    "github.com/gin-gonic/gin"
    "github.com/toefel18/address-service/geodan"
)

func GetAllAddressesSummary(dao geodan.AddressesNLDao) func(c *gin.Context) {
    return func (c *gin.Context) {
        count, error := dao.AddressCount()
        if error != nil {
            c.AbortWithError(500, error)
        } else {
            c.JSON(200, gin.H{"totalAddresses": count})
        }
    }
}

func GetAddressByKixcode(dao geodan.AddressesNLDao) func(c *gin.Context) {
    return func (c *gin.Context) {
        kixcode := c.Param("kixcode")
        address, error := dao.AddressByKixcode(kixcode)
        if error != nil {
            c.AbortWithError(500, error)
        } else {
            c.JSON(200, address)
        }
    }
}
