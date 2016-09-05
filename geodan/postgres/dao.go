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
	row, err := dao.db.Query("SELECT COUNT(*) FROM addressesnetherlands;")
	if err != nil {
		return 0, err
	}
	var count int
	row.Scan(&count)
	return count, nil
}

func (dao *Dao) AddressByKixcode(kixcode string) (geodan.AddressNL, error) {
	//TODO rewrite and use scanRowsAsAddresses to avoid duplication
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

func (dao *Dao) AddressesByQuery(postcode string, huisnummer int, huisletter string, huisnummerToevoeging string) ([]geodan.AddressNL, error) {
	//TODO create dynamic where clause
	rows, err := dao.db.Query("SELECT aobjectid, kixcode, nraandid, oruimteid, straat, huisnummer, huisletter, huisnrtoev, postcode, wnpcode, woonplaats, gemcode, gemeente, provcode, provincie, buurtcode, buurtnr, straatnen, aotype, status, oppvlakte, gebrksdoel, x_rd, y_rd, lat, long FROM addressesnetherlands WHERE ");
	if err != nil {
		return make([]geodan.AddressNL, 0), err
	}
	defer rows.Close()
	return dao.scanRowsAsAddresses(rows)
}

func (dao *Dao) scanRowsAsAddresses(rows *sql.Rows) ([]geodan.AddressNL, error) {
	addresses := make([]geodan.AddressNL, 0)
	for rows.Next() {
		address := geodan.AddressNL{}
		err := rows.Scan(
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
		if err != nil {
			return addresses, err
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

func (dao *Dao) Close() {
	if dao.db != nil {
		err := dao.db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}
