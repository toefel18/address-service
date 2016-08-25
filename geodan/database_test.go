package geodan_test

import (
	"github.com/toefel18/address-service/geodan"
	"strings"
	"testing"
	"fmt"
)

const GEODAN_TESTSET = `"aobjectid";"nraandid";"oruimteid";"straat";"huisnummer";"huisletter";"huisnrtoev";"postcode";"wnpcode";"woonplaats";"gemcode";"gemeente";"provcode";"provincie";"buurtcode";"buurtnr";"straatnen";"aotype";"status";"oppvlakte";"gebrksdoel";"x_rd";"y_rd";"lat";"long"
"0003010000126006";"0003200000134086";"0003300000117006";"Burgemeester Lewe van Aduardstraat";52;"g";"5";"9902NN";"3386";"Appingedam";"0003";"Appingedam";"20";"Groningen";"00030002";30002;"Burg Lewe van Aduardstr";"v";"Verblijfsobject in gebruik";10;"overige gebruiksfunctie";253463.07;592651.11;53.311231855738;6.86450306650747
"0003010000126007";"0003200000134087";"0003300000117006";"Burgemeester Lewe van Aduardstraat";52;"g";"4";"9902NN";"3386";"Appingedam";"0003";"Appingedam";"20";"Groningen";"00030002";30002;"Burg Lewe van Aduardstr";"v";"Verblijfsobject in gebruik";11;"overige gebruiksfunctie";253462.74;592648.51;53.3112085618188;6.8644972714828
"0003010000126008";"0003200000134088";"0003300000117006";"Burgemeester Lewe van Aduardstraat";52;"g";"3";"9902NN";"3386";"Appingedam";"0003";"Appingedam";"20";"Groningen";"00030002";30002;"Burg Lewe van Aduardstr";"v";"Verblijfsobject in gebruik";13;"overige gebruiksfunctie";253462.07;592645.64;53.3111829400536;6.86448638500512
"0003010000126009";"0003200000134089";"0003300000117006";"Burgemeester Lewe van Aduardstraat";52;"g";"2";"9902NN";"3386";"Appingedam";"0003";"Appingedam";"20";"Groningen";"00030002";30002;"Burg Lewe van Aduardstr";"v";"Verblijfsobject in gebruik";12;"overige gebruiksfunctie";253461.87;592643.11;53.3111602234592;6.86448260544979
"0003010000126010";"0003200000134090";"0003300000117006";"Burgemeester Lewe van Aduardstraat";52;"g";"1";"9902NN";"3386";"Appingedam";"0003";"Appingedam";"20";"Groningen";"00030002";30002;"Burg Lewe van Aduardstr";"v";"Verblijfsobject in gebruik";12;"overige gebruiksfunctie";253461.34;592640.04;53.3111327804224;6.86447366724567
"0003010000126363";"0003200000134263";"0003300000117175";"Professor Cleveringaplein";1;"";"";"9901AZ";"3386";"Appingedam";"0003";"Appingedam";"20";"Groningen";"00030000";30000;"Prof Cleveringaplein";"v";"Verblijfsobject in gebruik";1336;"kantoorfunctie";252828.00;593714.00;53.3208963439582;6.85530301640351
"0003010000126364";"0003200000134264";"0003300000117175";"Professor Cleveringaplein";3;"";"";"9901AZ";"3386";"Appingedam";"0003";"Appingedam";"20";"Groningen";"00030000";30000;"Prof Cleveringaplein";"v";"Verblijfsobject in gebruik";1392;"winkelfunctie";252842.00;593684.00;53.3206242938808;6.85550386388695
"0003010000126329";"0003200000134229";"0003300000117038";"Dijkstraat";40;"";"";"9901AT";"3386";"Appingedam";"0003";"Appingedam";"20";"Groningen";"00030000";30000;"Dijkstraat";"v";"Verblijfsobject in gebruik";102;"winkelfunctie";253027.00;593697.00;53.3207071331691;6.85828322530634
"0003010000126365";"0003200000134265";"0003300000117175";"Professor Cleveringaplein";4;"";"";"9901AZ";"3386";"Appingedam";"0003";"Appingedam";"20";"Groningen";"00030000";30000;"Prof Cleveringaplein";"v";"Verblijfsobject in gebruik";158;"winkelfunctie";252866.12;593681.63;53.3205986202316;6.85586492868056
"0003010000128481";"0003200000136871";"0003300000117390";"Professor d Blécourtstraat";1;"";"";"9902EC";"3386";"Appingedam";"0003";"Appingedam";"20";"Groningen";"00030002";30002;"Prof d Blécourtstraat";"v";"Verblijfsobject in gebruik";110;"woonfunctie";253738.00;593589.00;53.3196059682132;6.86891629013341
"0003010000128482";"0003200000136872";"0003300000117390";"Professor d Blécourtstraat";6;"";"";"9902EC";"3386";"Appingedam";"0003";"Appingedam";"20";"Groningen";"00030002";30002;"Prof d Blécourtstraat";"v";"Verblijfsobject in gebruik";164;"woonfunctie";253689.00;593602.00;53.3197318042909;6.86818522423476
`

func TestBasicIntegration(t *testing.T) {
	db := geodan.CreateFromReader(strings.NewReader(GEODAN_TESTSET))
	address, err := db.FindByKixcode("NL9902NN000052XG5")
	if err != nil {
		t.Fatal("db.FindByKixcode returned an error")
	}
	if address.Kixcode() != db.Addresses[0].Kixcode() {
		t.Fatal("the returned address is not the expected: ", address)
	}
	if address.Type() != geodan.BUSINESS {
		t.Fatal("the returned address is not 2B but", address.Type())
	}
}
