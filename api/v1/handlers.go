package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/toefel18/address-service/geodan"
	"net/http"
)

type KixcodeAndType struct {
	KixCode     string `json:"kixCode"`
	KixCodeType string `json:"type"`
}

func AddressByKixcode(db geodan.AddressesNLDao) func(c *gin.Context) {
	return func(c *gin.Context) {
		kixcode := c.Param("kixcode")
		c.IndentedJSON(http.StatusOK, createResult(db, kixcode))
	}
}

func MultipleAddresses(db geodan.AddressesNLDao) func(ctx *gin.Context) {
    return func(c *gin.Context) {
        data := make([]string,0)
        c.BindJSON(&data)
        kixTypeArray := make([]KixcodeAndType, 0)
        for _, kixcode := range data {
            kixTypeArray = append(kixTypeArray, createResult(db, kixcode))
        }
        c.IndentedJSON(http.StatusOK, kixTypeArray)
    }
}


func createResult(db geodan.AddressesNLDao, kixcode string) KixcodeAndType {
	kixcodeType := geodan.BUSINESS
	address, err := db.AddressByKixcode(kixcode)
	if err == nil {
		if address.Gebrksdoel == "woonfunctie" {
			kixcodeType = "2C"
		} else {
			kixcodeType = "2B"
		}
	}
	return KixcodeAndType{kixcode, kixcodeType}
}

