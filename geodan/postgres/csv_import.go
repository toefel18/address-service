package postgres

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"
    "github.com/toefel18/address-service/geodan"
    "database/sql"
    _ "github.com/lib/pq" // import with _ prefixed means importing solely for it's side-effects
    "strings"
    "fmt"
)

func Import(filename string)  {
	log.Print("Reading GEODAN file")
	start := time.Now()

	file, err := os.Open(filename)
    defer file.Close()
    if err != nil {
        panic(err)
    }
    defer log.Print("Done in ", time.Since(start).Seconds(), " Seconds")
    ImportFromReader(file)
}

func ImportFromReader(reader io.Reader) {
    csvReader := csv.NewReader(bufio.NewReader(reader))
    csvReader.Comma = ';'
    csvExtract := geodan.NewCsvExtract()
    readIntoDatabase(csvExtract, csvReader)
}

func readIntoDatabase(csvExtract *geodan.CsvExtract, csvReader *csv.Reader) {
    readHeaderIntoCsvExtract(csvExtract, csvReader)
    insertRecordsIntoDatabase(csvExtract, csvReader)

}
func readHeaderIntoCsvExtract(csvExtract *geodan.CsvExtract, csvReader *csv.Reader) {
    header, err := csvReader.Read()
    if err == io.EOF {
        log.Print("WARN, could not read HEADER")
        panic(err)
    }
    for index, element := range header {
        csvExtract.Headers[element] = index
    }
}

func extractKixcode(csvExtract *geodan.CsvExtract, record []string) string {
    postcalCode := record[csvExtract.Headers["postcode"]]
    housenr :=  record[csvExtract.Headers["huisnummer"]]
    houseletter :=  record[csvExtract.Headers["huisletter"]]
    addition :=  record[csvExtract.Headers["huisnrtoev"]]
    return strings.ToUpper(fmt.Sprintf("NL%v%06vX%v%v", postcalCode, housenr, houseletter, addition))
}

func insertRecordsIntoDatabase(csvExtract *geodan.CsvExtract, csvReader *csv.Reader) {
    db, err := sql.Open("postgres", "postgres://postgres:root@localhost/postgres?sslmode=disable")

    if err != nil {
        panic (err)
    }
    defer db.Close()

    counter := 0

    stmt, _  := db.Prepare("INSERT INTO addressesnetherlands (" +
        "aobjectid," +
        "kixcode," +
        "nraandid," +
        "oruimteid," +
        "straat," +
        "huisnummer," +
        "huisletter," +
        "huisnrtoev," +
        "postcode," +
        "wnpcode," +
        "woonplaats," +
        "gemcode," +
        "gemeente," +
        "provcode," +
        "provincie," +
        "buurtcode," +
        "buurtnr," +
        "straatnen," +
        "aotype," +
        "status," +
        "oppvlakte," +
        "gebrksdoel," +
        "x_rd," +
        "y_rd," +
        "lat," +
        "long) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)")

    tx, _ := db.Begin()

    for {
        record, err := csvReader.Read()
        if err == io.EOF {
            tx.Commit()
            break
        } else if err != nil {
            log.Fatal(err)
        }
        _, err = stmt.Exec(
            record[0],
            extractKixcode(csvExtract, record),
            record[1],
            record[2],
            record[3],
            record[4],
            record[5],
            record[6],
            record[7],
            record[8],
            record[9],
            record[10],
            record[11],
            record[12],
            record[13],
            record[14],
            record[15],
            record[16],
            record[17],
            record[18],
            record[19],
            record[20],
            record[21],
            record[22],
            record[23],
            record[24])

        counter += 1

        if err != nil {
            panic(err)
        }

        if counter % 9999 == 0 {
            tx.Commit() // to see progress while import is running
            tx, _ = db.Begin()
            log.Println("Inserted ", counter, " addresses" )
        }
    }
}

