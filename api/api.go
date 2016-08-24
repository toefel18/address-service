package api

import (
	"github.com/gin-gonic/gin"
	"github.com/toefel18/address-service/geodan"
)

func Publish(db *geodan.GeodanDB) {
	router := gin.Default()
	router.GET("/v1/addresses/:kixcode", v1Address(db))
	router.Run("0.0.0.0:8888")
}

func v1Address(db *geodan.GeodanDB) func (ctx *gin.Context) {
    return func(ctx *gin.Context) {
        kix := ctx.Param("kixcode")
        kixcodeType := geodan.BUSINESS
        address, err := db.FindByKixcode(kix)
        if err != nil {
            kixcodeType = address.Type()
        }

        ctx.JSON(200, gin.H{
            "kixCode": kix,
            "type": kixcodeType,
        })
    }
}


//package api
//
//import (
//"github.com/gin-gonic/gin"
//"github.com/toefel18/address-service/geodan"
//)
//
//type RestApi struct {
//    db *geodan.GeodanDB
//}
//
//func NewRestApi(db *geodan.GeodanDB) *RestApi {
//    return &RestApi{db}
//}
//
//func (api *RestApi) Start() {
//    router := gin.Default()
//    router.GET("/v1/addresses/:kixcode", api.v1Address)
//    router.Run("0.0.0.0:8888") // listen and server on 0.0.0.0:8080
//}
//
//func (api *RestApi) v1Address(ctx *gin.Context) {
//    kix := ctx.Param("kixcode")
//    kixcodeType := geodan.BUSINESS
//    address, err := api.db.FindByKixcode(kix)
//    if err != nil {
//        kixcodeType = address.Type()
//    }
//
//    ctx.JSON(200, gin.H{
//        "kixCode": kix,
//        "type": kixcodeType,
//    })
//}
