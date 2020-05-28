package main

type ChildInfo struct {
	LName    string //BTAL*->CHILD->LNAME
	FName    string //BTAL*->CHILD->FNAME
	MName    string //BTAL*->CHILD->MNAME
	BDate    string //BTAL*->CHILD->BDATE
	DatePobs string //BTAL*->CHILD->DATEPOBS
	DateBobs string //BTAL*->CHILD->DATEBOBS
	DateEobs string //BTAL*->CHILD->DATEEOBS
	SPolicy  string //BTAL*->CHILD->POLICY->SPOLICY
	NPolicy  string //BTAL*->CHILD->POLICY->NPOLICY
}

type TalonInfo struct {
	Filename         string
	NTalon           string      //Номер талона: 1, 2, 3-1, 3-2
	DReg             string      //BTAL*->DREG (1,2,3)
	SCertif          string      //BTAL*->CERTIF->SCERTIF (1,2,3)
	NCertif          string      //BTAL*->CERTIF->NCERTIF (1,2,3)
	DCertif          string      //BTAL*->CERTIF->DCERTIF (1,2,3)
	Snils            string      //BTAL*->PERSON->SNILS (1,2,3)
	LName            string      //BTAL*->PERSON->LNAME (1,2,3)
	FName            string      //BTAL*->PERSON->FNAME (1,2,3)
	MName            string      //BTAL*->PERSON->MNAME (1,2,3)
	BDate            string      //BTAL*->PERSON->BDATE (1,2,3)
	PassportTDoc     string      //BTAL*->PASSPORT->TDOC (1,2,3)
	PassportTDocName string      //BTAL*->PASSPORT->TDOC (1,2,3)
	PassportNDoc     string      //BTAL*->PASSPORT->NDOC (1,2,3)
	PassportSDoc     string      //BTAL*->PASSPORT->SDOC (1,2,3)
	PassportDDoc     string      //BTAL*->PASSPORT->DDOC (1,2,3)
	PassportODoc     string      //BTAL*->PASSPORT->ODOC (1,2,3)
	SPolicy          string      //BTAL*->POLICY->SPOLICY (1,2,3)
	NPolicy          string      //BTAL*->POLICY->NPOLICY (1,2,3)
	SLeaf            string      //BTAL*->SICKLIST->SLEAF (1,2)
	NLeaf            string      //BTAL*->SICKLIST->NLEAF (1,2)
	DLeaf            string      //BTAL*->SICKLIST->DLEAF (1,2)
	NCard            string      //BTAL*->EXCCARD->NCARD (1,2,3)
	DCard            string      //BTAL*->EXCCARD->DCARD (1,2,3)
	ChildDiag        string      //BTAL*->CHILD->DIAG (2)
	ChildGrowth      string      //BTAL*->CHILD->GROWTH (2)
	ChildWeight      string      //BTAL*->CHILD->WEIGHT (2)
	ChildSex         string      //BTAL*->CHILD->SEX (2)
	Childs           []ChildInfo //BTAL*->CHILD (3)
}

func TDocName(td string) string {
	switch td {
	case "1":
		return "Паспорт гражданина СССР"
	case "2":
		return "Загранпаспорт гражданина СССР"
	case "3":
		return "Свидетельство о рождении"
	case "4":
		return "Удостоверение личности офицера"
	case "5":
		return "Справка об освобождении из места лишения свободы"
	case "6":
		return "Паспорт Минморфлота"
	case "7":
		return "Военный билет солдата (матроса, сержанта, старшины)"
	case "8":
		return "Дипломатический паспорт гражданина РФ"
	case "9":
		return "Иностранный паспорт"
	case "10":
		return "Свидетельство о регистрации ходатайства о признании иммигранта беженцем"
	case "11":
		return "Вид на жительство"
	case "12":
		return "Удостоверение беженца в РФ"
	case "13":
		return "Временное удостоверение личности гражданина РФ"
	case "14":
		return "Паспорт гражданина России"
	case "15":
		return "Загранпаспорт гражданина РФ"
	case "16":
		return "Паспорт моряка"
	case "17":
		return "Военный билет офицера запаса"
	case "18":
		return "Иные документы, выдаваемые органами МВД"
	default:
		return "Неизвестный документ"
	}
}
