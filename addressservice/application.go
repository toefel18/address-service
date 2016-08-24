package main

import (
	"github.com/toefel18/address-service/geodan"
	"github.com/toefel18/address-service/api"
)

func main() {
	geodan := geodan.CreateFromFile("/home/hestersco/adressen_latlong_201507_landelijk.csv")
	api.Publish(geodan)
}
