package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/toefel18/address-service/geodan"
	"net/http"
    "fmt"
)

type KixcodeAndType struct {
	KixCode     string `json:"kixCode"`
	KixCodeType string `json:"type"`
}

func AddressByKixcode(db *geodan.GeodanDB) func(c *gin.Context) {
	return func(c *gin.Context) {
		kixcode := c.Param("kixcode")
		c.JSON(http.StatusOK, createResult(db, kixcode))
	}
}

func MultipleAddresses(db *geodan.GeodanDB) func(ctx *gin.Context) {
    return func(c *gin.Context) {
        data := make([]string,0)
        c.BindJSON(&data)
        kixTypeArray := make([]KixcodeAndType, 0)
        for _, kixcode := range data {
            kixTypeArray = append(kixTypeArray, createResult(db, kixcode))
        }
        c.JSON(http.StatusOK, kixTypeArray)
    }
}


func createResult(db *geodan.GeodanDB, kixcode string) KixcodeAndType {
	kixcodeType := geodan.BUSINESS
	address, err := db.FindByKixcode(kixcode)
	if err == nil {
		kixcodeType = address.Type()
	}
	return KixcodeAndType{kixcode, kixcodeType}
}

