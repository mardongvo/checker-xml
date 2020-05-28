package main

//"reflect"
//"strings"
//"testing"

/*
var xmltest = []struct {
	xmltext string
	err     error
	result  []TalonInfo
}{
	{
		`<?xml version="1.0" encoding="WINDOWS-1251"?><?BSERTIF1 version="0.2"?><BIRTH_SERTIF xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <HEADER DATE="000" TO="" KPP="" INN="" FROM=""/>
  <BTAL1 PREM="0" MULP="0" QTW2="0" QTW1="0" PHELP="0" CHK="1" DREG="2020.01.01">
    <CERTIF DCERTIF="2020.01.01" NCERTIF="111111" SCERTIF="G"/>
    <PERSON ADDRESS="" BDATE="1989.01.01" MNAME="O" FNAME="I" LNAME="F" SNILS="000-000-000 11"/>
    <PASSPORT ODOC="ФМС" DDOC="2015.01.01" SDOC="1515" NDOC="111111" TDOC="14"/>
    <POLICY NPOLICY="1234567" SPOLICY=""/>
    <SICKLIST NLEAF="N1" SLEAF="S1"/>
    <EXCCARD DCARD="2019.01.01" NCARD="1/1"/>
  </BTAL1>
  <BTAL1 PREM="0" MULP="0" QTW2="0" QTW1="0" PHELP="0" CHK="1" DREG="2020.02.02">
    <CERTIF DCERTIF="2020.02.02" NCERTIF="222222" SCERTIF="Г"/>
    <PERSON ADDRESS="" BDATE="1990.01.01" MNAME="O" FNAME="I" LNAME="F" SNILS="000-000-000 22"/>
    <PASSPORT ODOC="ФМС" DDOC="2016.01.01" SDOC="1616" NDOC="222222" TDOC="14"/>
    <POLICY NPOLICY="1234567" SPOLICY=""/>
    <SICKLIST NLEAF="" SLEAF=""/>
    <EXCCARD DCARD="2019.02.02" NCARD="2/2"/>
  </BTAL1>
</BIRTH_SERTIF>`,
		nil,
		[]TalonInfo{
			TalonInfo{
				NTalon:           "1",
				DReg:             "2020.01.01",
				SCertif:          "",
				NCertif:          "",
				DCertif:          "",
				Snils:            "",
				LName:            "",
				FName:            "",
				MName:            "",
				BDate:            "",
				PassportTDoc:     "",
				PassportTDocName: "",
				PassportNDoc:     "",
				PassportSDoc:     "",
				PassportDDoc:     "",
				PassportODoc:     "",
				SPolicy:          "",
				NPolicy:          "",
				SLeaf:            "",
				NLeaf:            "",
				NCard:            "",
				DCard:            "",
			},
		},
	},
}

func TestParse(t *testing.T) {
	for _, test := range xmltest {
		r := strings.NewReader(test.xmltext)
		ti, err := Xml2Struct(r)
		if err != test.err {
			t.Errorf("Xml2Struct error gives \"%#v\", expected \"%#v\"",
				err, test.err)
		}
		if !reflect.DeepEqual(ti, test.result) {
			t.Errorf("Xml2Struct gives \"%#v\", expected \"%#v\"",
				ti, test.result)
		}
	}

}
*/
