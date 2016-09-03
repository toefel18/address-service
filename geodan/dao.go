package geodan

type AddressNL struct {
    Aobjectid  string
    Kixcode    string
    Nraandid   string
    Oruimteid  string
    Straat     string
    Huisnummer int
    Huisletter string
    Huisnrtoev string
    Postcode   string
    Wnpcode    string
    Woonplaats string
    Gemcode    string
    Gemeente   string
    Provcode   string
    Provincie  string
    Buurtcode  string
    Buurtnr    int
    Straatnen  string
    Aotype     string
    Status     string
    Oppvlakte  int
    Gebrksdoel string
    X_rd       float64
    Y_rd       float64
    Lat        float64
    Long       float64
}

type AddressesNLDao interface {
    AddressCount() (int, error)
    AddressByKixcode(kixcode string) (AddressNL, error)
    Close()
}