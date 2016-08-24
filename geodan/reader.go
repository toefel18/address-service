package geodan

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"
)

func CreateFromFile(filename string) *GeodanDB {
	log.Print("Reading GEODAN file")
	start := time.Now()

	file, err := os.Open(filename)
    defer file.Close()
    if err != nil {
        panic(err)
    }
    defer log.Print("Done in ", time.Since(start).Seconds(), " Seconds")
    return CreateFromReader(file)
}

func CreateFromReader(reader io.Reader) * GeodanDB {
    csvReader := csv.NewReader(bufio.NewReader(reader))
    csvReader.Comma = ';'
    db := NewDatabase()
    readIntoDatabase(db, csvReader)
    return db
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
        if err == io.EOF || len(db.Addresses) > 100000 {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        db.Append(Address{fields: record})
    }
}