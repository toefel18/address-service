package v2

import (
    "github.com/gin-gonic/gin"
    "github.com/toefel18/address-service/geodan/postgres/dao"
)

func getAllAddressesSummary(dao * dao.AddressesNLDao) func(c *gin.Context) {
    return func (c *gin.Context) {
        count, error := dao.AddressCount()
        if error != nil {
            c.AbortWithError(500, error)
        } else {
            c.JSON(200, gin.H{"totalAddresses", int32(count)})
        }
    }
}
