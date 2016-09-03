package main

import (
	//"github.com/toefel18/address-service/geodan"
	//"github.com/toefel18/address-service/api"
	"github.com/toefel18/address-service/geodan/postgres"
	"fmt"
)

func main() {
	dao, err := postgres.NewDao("postgres://postgres:root@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		fmt.Println("Could not open database connection")
		return
	}
	fmt.Println(dao.AddressCount())
	//fmt.Println(dao.AddressByKixcode("NL4576AB000006X"))

	//postgres.Import("/home/hestersco/adressen_latlong_201507_landelijk.csv")
	//geodan := geodan.CreateFromFile("/home/hestersco/adressen_latlong_201507_landelijk.csv")
	//api.Publish(geodan)
}
