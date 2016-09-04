package geodan

const (
    BUSINESS = "2B"
    CUSTOMER = "2C")

// TODO fix camelcasing json fields
type AddressNL struct {
    Aobjectid  string  `json:"objectId"`
    Kixcode    string  `json:"kixCode"`
    Nraandid   string  `json:"nraandid"`  //TODO find out what this means by looking in geodan column deff
    Oruimteid  string  `json:"ruimteId"`
    Straat     string  `json:"straat"`
    Huisnummer int     `json:"huisnummer"`
    Huisletter string  `json:"huisletter"`
    Huisnrtoev string  `json:"huisnummerToevoeging"`
    Postcode   string  `json:"postcode"`
    Wnpcode    string  `json:"woonplaatsCode"`
    Woonplaats string  `json:"woonplaats"`
    Gemcode    string  `json:"gemeenteCode"`
    Gemeente   string  `json:"gemeente"`
    Provcode   string  `json:"provincieCode"`
    Provincie  string  `json:"provincie"`
    Buurtcode  string  `json:"buurtCode"`
    Buurtnr    int     `json:"buurtNummer"`
    Straatnen  string  `json:"straatnen"` //TODO lookup meaning
    Aotype     string  `json:"objectType"`
    Status     string  `json:"status"`
    Oppvlakte  int     `json:"oppervlakte"`
    Gebrksdoel string  `json:"gebruiksdoel"`
    X_rd       float64 `json:"xRijksdriehoek"`
    Y_rd       float64 `json:"yRijksdriehoek"`
    Lat        float64 `json:"latitude"`
    Long       float64 `json:"longitude"`
}

func (address *AddressNL) Type () string {
    if address.Gebrksdoel == "woonfunctie" {
        return CUSTOMER
    } else {
        return BUSINESS
    }
}

type AddressesNLDao interface {
    AddressCount() (int, error)
    AddressByKixcode(kixcode string) (AddressNL, error)
    AddressesByQuery(postcode string, huisnummer int, huisletter string, huisnummerToevoeging string) ([]AddressNL, error)
    Close()
}
