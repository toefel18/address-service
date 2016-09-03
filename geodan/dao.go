package geodan

type AddressNL struct {
    aobjectid  string
    kixcode    string
    nraandid   string
    oruimteid  string
    straat     string
    huisnummer int
    huisletter string
    huisnrtoev string
    postcode   string
    wnpcode    string
    woonplaats string
    gemcode    string
    gemeente   string
    provcode   string
    provincie  string
    buurtcode  string
    buurtnr    int
    straatnen  string
    aotype     string
    status     string
    oppvlakte  int
    gebrksdoel string
    x_rd       float64
    y_rd       float64
    lat        float64
    long       float64
}

type AddressesNLDao interface {
    AddressCount() (int, error)
    AddressByKixcode(kixcode string) (AddressNL, error)
    Close()
}