package postgres

import (
    "github.com/toefel18/address-service/geodan"
    "database/sql"
    "fmt"
)

type Dao struct {
    ConnectionString string
    db *sql.DB
}

func NewDao (connectionString string) (*Dao, error) {
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
    row := dao.db.QueryRow("SELECT aobjectid, kixcode, nraandid, oruimteid, straat, huisnummer, huisletter, huisnrtoev, postcode, wnpcode, woonplaats, gemcode, gemeente, provcode, provincie, buurtcode, buurtnr, straatnen, aotype, status, oppvlakte, gebrksdoel, x_rd, y_rd, lat, long FROM addressesnetherlands;")
    address := geodan.AddressNL{}
    row.Scan(&address.aobjectid)
    return count, nil
    return geodan.AddressNL{}, nil
}

func (dao *Dao) Close() {
    if dao.db != nil {
        err := dao.db.Close()
        if err != nil {
            fmt.Println(err)
        }
    }
}
