package geodan

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"
)

type GeodanDB struct {
	Headers map[string]int
	Records [][]string
}

func ReadGeodan(filename string) *GeodanDB {
	log.Print("Reading GEODAN file")
	start := time.Now()

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(bufio.NewReader(file))
	csvReader.Comma = ';'
	db := GeodanDB{}
	readIntoDatabase(&db, csvReader)
	log.Print("Done in", time.Since(start).Seconds(), "Seconds")
	return &db
}

func readIntoDatabase(db *GeodanDB, csvReader *csv.Reader) {
    readHeaderIntoDb(db, csvReader)
    readRecordsIntoDb(db, csvReader)

}
func readHeaderIntoDb(db *GeodanDB, csvReader *csv.Reader) {
    header, err := csvReader.Read()
    if err == io.EOF {
        log.Print("WARN, could not read HEADER")
        panic(err)
    }
    for index, element := range header {
        db.Headers[element] = index
    }
}

func readRecordsIntoDb(db *GeodanDB, csvReader *csv.Reader) {
    for {
        //log.Print("length of ", len(db.Records), cap(db.Records))
        record, err := csvReader.Read()
        if err == io.EOF || len(db.Records) > 100000 {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        db.Records = append(db.Records, record)
    }
}