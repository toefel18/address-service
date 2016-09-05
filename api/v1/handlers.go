package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/toefel18/address-service/geodan"
	"net/http"
	"github.com/toefel18/address-service/api/area"
)

type KixcodeAndType struct {
	KixCode     string `json:"kixCode"`
	KixCodeType string `json:"type"`
}

func AddressByKixcode(db geodan.AddressesNLDao, areaStore *area.AreaStore) func(c *gin.Context) {
	return func(c *gin.Context) {
		kixcode := c.Param("kixcode")
		c.IndentedJSON(http.StatusOK, createResult(db, kixcode, areaStore))
	}
}

func MultipleAddresses(db geodan.AddressesNLDao, areaStore *area.AreaStore) func(ctx *gin.Context) {
    return func(c *gin.Context) {
        data := make([]string,0)
        c.BindJSON(&data)
        kixTypeArray := make([]KixcodeAndType, 0)
        for _, kixcode := range data {
            kixTypeArray = append(kixTypeArray, createResult(db, kixcode, areaStore))
        }
        c.IndentedJSON(http.StatusOK, kixTypeArray)
    }
}


func createResult(db geodan.AddressesNLDao, kixcode string, areaStore * area.AreaStore) KixcodeAndType {
	kixcodeType := geodan.BUSINESS
	address, err := db.AddressByKixcode(kixcode)
	if err == nil {
		if address.Gebrksdoel == "woonfunctie" && !areaStore.Excludes(kixcode){
			kixcodeType = "2C"
		} else {
			kixcodeType = "2B"
		}
	}
	return KixcodeAndType{kixcode, kixcodeType}
}

func ExcludeArea(areaStore *area.AreaStore) func(c *gin.Context) {
	return func(c *gin.Context) {
		kixcodeBegin := c.Param("kixcodeBeginIncl")
		kixcodeEnd := c.Param("kixcodeEndIncl")
		area, found := areaStore.Get(kixcodeBegin, kixcodeEnd)
		if found {
			c.JSON(200, area)
		} else {
			c.Status(404)
		}
	}
}

func AddExcludeArea(areaStore *area.AreaStore) func(c *gin.Context) {
	return func(c *gin.Context) {
		kixcodeBegin := c.Param("kixcodeBeginIncl")
		kixcodeEnd := c.Param("kixcodeEndIncl")
		newArea := area.Area{KixcodeBegin: kixcodeBegin, KixcodeEndInclusive: kixcodeEnd}
		areaStore.Add(newArea)
		c.Status(204)
	}
}

func DeleteExcludeArea(areaStore *area.AreaStore) func(c *gin.Context) {
	return func(c *gin.Context) {
		kixcodeBegin := c.Param("kixcodeBeginIncl")
		kixcodeEnd := c.Param("kixcodeEndIncl")
		areaStore.Remove(kixcodeBegin, kixcodeEnd)
		c.Status(204)
	}
}