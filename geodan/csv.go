package geodan

import (
	"errors"
	"fmt"
	"strings"
)

type CsvExtract struct {
	Headers   map[string]int
	Addresses []Address
}

func NewCsvExtract() *CsvExtract {
	return &CsvExtract{Headers: make(map[string]int)}
}

func (db *CsvExtract) FindByKixcode(kixcode string) (Address, error) {
	for _, address := range db.Addresses {
		if address.Kixcode() == kixcode {
			return address, nil
		}
	}
	return Address{}, errors.New(kixcode + "not found")
}

func (db *CsvExtract) Append(address Address) {
	address.db = db
	db.Addresses = append(db.Addresses, address)
}

type Address struct {
	fields []string
	db     *CsvExtract
}

func (address Address) Kixcode() string {
	postcalCode := address.GetField("postcode")
	housenr := address.GetField("huisnummer")
	houseletter := address.GetField("huisletter")
	addition := address.GetField("huisnrtoev")
	return strings.ToUpper(fmt.Sprintf("NL%v%06vX%v%v", postcalCode, housenr, houseletter, addition))
}

func (address Address) Type() string {
	if address.GetField("gebrksdoel") == "woonfunctie" {
		return CUSTOMER
	} else {
		return BUSINESS
	}
}

func (address Address) GetField(fieldName string) string {
	return address.fields[address.db.Headers[fieldName]]
}


