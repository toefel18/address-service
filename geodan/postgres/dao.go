package postgres

import (
	"database/sql"
	"fmt"
	"github.com/toefel18/address-service/geodan"
)

type Dao struct {
	ConnectionString string
	db               *sql.DB
}

func NewDao(connectionString string) (*Dao, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &Dao{connectionString, db}, nil
}

func (dao *Dao) AddressCount() (int, error) {
	row := dao.db.QueryRow("SELECT COUNT(*) FROM addressesnetherlands;")
	var count int
	row.Scan(&count)
	return count, nil
}

func (dao *Dao) AddressByKixcode(kixcode string) (geodan.AddressNL, error) {
	row := dao.db.QueryRow("SELECT aobjectid, kixcode, nraandid, oruimteid, straat, huisnummer, huisletter, huisnrtoev, postcode, wnpcode, woonplaats, gemcode, gemeente, provcode, provincie, buurtcode, buurtnr, straatnen, aotype, status, oppvlakte, gebrksdoel, x_rd, y_rd, lat, long FROM addressesnetherlands WHERE kixcode = $1", kixcode)
	address := geodan.AddressNL{}
	row.Scan(
		&address.Aobjectid,
		&address.Kixcode,
		&address.Nraandid,
		&address.Oruimteid,
		&address.Straat,
		&address.Huisnummer,
		&address.Huisletter,
		&address.Huisnrtoev,
		&address.Postcode,
		&address.Wnpcode,
		&address.Woonplaats,
		&address.Gemcode,
		&address.Gemeente,
		&address.Provcode,
		&address.Provincie,
		&address.Buurtcode,
		&address.Buurtnr,
		&address.Straatnen,
		&address.Aotype,
		&address.Status,
		&address.Oppvlakte,
		&address.Gebrksdoel,
		&address.X_rd,
		&address.Y_rd,
		&address.Lat,
		&address.Long)
	return address, nil
}

func (dao *Dao) Close() {
	if dao.db != nil {
		err := dao.db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}
