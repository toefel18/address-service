package main

import (
	//"github.com/toefel18/address-service/api"
	"fmt"
	"github.com/toefel18/address-service/geodan"
)

func main() {
	geodan := geodan.ReadGeodan("/home/hestersco/adressen_latlong_201507_landelijk.csv")
	fmt.Println("Record 1")
	fmt.Println(geodan.Records[0])
	fmt.Println("Record 2")
	fmt.Println(geodan.Records[1])
	//api.RunRest()
}
