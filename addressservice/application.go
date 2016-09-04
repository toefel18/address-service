package main

import (
	//"github.com/toefel18/address-service/geodan"
	//"github.com/toefel18/address-service/api"
	"github.com/toefel18/address-service/geodan/postgres"
	"fmt"
	"github.com/toefel18/address-service/api"
)

func main() {
	dao, err := postgres.NewDao("postgres://postgres:root@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Println("Could not open database connection")
		return
	}
	fmt.Println("Publishing API")
	//postgres.Import("/home/hestersco/adressen_latlong_201507_landelijk.csv")
	api.Publish(dao)
}
