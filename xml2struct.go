package main

import (
	"encoding/xml"
	"io"
	"strings"

	"golang.org/x/net/html/charset"
)

func getAttributeValue(src []xml.Attr, key string) string {
	for _, v := range src {
		if v.Name.Local == key {
			return v.Value
		}
	}
	return ""
}

func Xml2Struct(r io.Reader, filename string) ([]TalonInfo, error) {
	var res []TalonInfo = make([]TalonInfo, 0)
	var talontype string
	var currentTalon TalonInfo
	var currentChild ChildInfo
	var isChildProcessing bool = false
	dec := xml.NewDecoder(r)
	dec.CharsetReader = charset.NewReaderLabel
	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return res, err
		}
		switch t := token.(type) {
		case xml.ProcInst:
			{
				if t.Target == "BSERTIF1" {
					talontype = "1"
				}
				if t.Target == "BSERTIF2" {
					talontype = "2"
				}
				if t.Target == "BSERTIF3_1" {
					talontype = "3-1"
				}
				if t.Target == "BSERTIF3_2" {
					talontype = "3-2"
				}
			}
		case xml.StartElement:
			{
				if strings.Index(t.Name.Local, "BTAL") == 0 { //начало талона
					currentTalon = TalonInfo{}
					currentTalon.Filename = filename
					currentTalon.NTalon = talontype
					currentTalon.Childs = make([]ChildInfo, 0)
					currentTalon.DReg = getAttributeValue(t.Attr, "DREG")
				}
				if t.Name.Local == "CERTIF" {
					currentTalon.SCertif = getAttributeValue(t.Attr, "SCERTIF")
					currentTalon.NCertif = getAttributeValue(t.Attr, "NCERTIF")
					currentTalon.DCertif = getAttributeValue(t.Attr, "DCERTIF")
				}
				if t.Name.Local == "PERSON" {
					currentTalon.Snils = getAttributeValue(t.Attr, "SNILS")
					currentTalon.LName = getAttributeValue(t.Attr, "LNAME")
					currentTalon.FName = getAttributeValue(t.Attr, "FNAME")
					currentTalon.MName = getAttributeValue(t.Attr, "MNAME")
					currentTalon.BDate = getAttributeValue(t.Attr, "BDATE")
				}
				if t.Name.Local == "PASSPORT" {
					currentTalon.PassportTDoc = getAttributeValue(t.Attr, "TDOC")
					currentTalon.PassportTDocName = TDocName(currentTalon.PassportTDoc)
					currentTalon.PassportNDoc = getAttributeValue(t.Attr, "NDOC")
					currentTalon.PassportSDoc = getAttributeValue(t.Attr, "SDOC")
					currentTalon.PassportDDoc = getAttributeValue(t.Attr, "DDOC")
					currentTalon.PassportODoc = getAttributeValue(t.Attr, "ODOC")
				}
				if t.Name.Local == "POLICY" {
					if isChildProcessing {
						currentChild.SPolicy = getAttributeValue(t.Attr, "SPOLICY")
						currentChild.NPolicy = getAttributeValue(t.Attr, "NPOLICY")
					} else {
						currentTalon.SPolicy = getAttributeValue(t.Attr, "SPOLICY")
						currentTalon.NPolicy = getAttributeValue(t.Attr, "NPOLICY")
					}
				}
				if t.Name.Local == "SICKLIST" {
					currentTalon.SLeaf = getAttributeValue(t.Attr, "SLEAF")
					currentTalon.NLeaf = getAttributeValue(t.Attr, "NLEAF")
				}
				if t.Name.Local == "EXCCARD" {
					currentTalon.NCard = getAttributeValue(t.Attr, "NCARD")
					currentTalon.DCard = getAttributeValue(t.Attr, "DCARD")
				}
				if t.Name.Local == "CHILD" {
					if talontype == "2" {
						currentTalon.ChildDiag = getAttributeValue(t.Attr, "DIAG")
						currentTalon.ChildGrowth = getAttributeValue(t.Attr, "CROWTH")
						currentTalon.ChildWeight = getAttributeValue(t.Attr, "WEIGHT")
						currentTalon.ChildSex = getAttributeValue(t.Attr, "SEX")
					} else {
						currentChild = ChildInfo{}
						currentChild.LName = getAttributeValue(t.Attr, "LNAME")
						currentChild.FName = getAttributeValue(t.Attr, "FNAME")
						currentChild.MName = getAttributeValue(t.Attr, "MNAME")
						currentChild.BDate = getAttributeValue(t.Attr, "BDATE")
						currentChild.DatePobs = getAttributeValue(t.Attr, "DATEPOBS")
						currentChild.DateBobs = getAttributeValue(t.Attr, "DATEBOBS")
						currentChild.DateEobs = getAttributeValue(t.Attr, "DATEEOBS")
					}

				}
			}
		case xml.EndElement:
			{
				if (t.Name.Local == "CHILD") && (talontype != "2") {
					currentTalon.Childs = append(currentTalon.Childs, currentChild)
				}
				if strings.Index(t.Name.Local, "BTAL") == 0 { //конец талона
					res = append(res, currentTalon)
				}
			}
		}
	}
	return res, nil
}
