package main

import (
	"fmt"
	"reflect"
)

type Accumulator interface {
	UsePair() bool
	PutOne(ti TalonInfo)
	PutPair(t1 TalonInfo, t2 TalonInfo)
	Header() []string
	Len() int
	Get(i int) []string
}

/*
//Тестовая проверка
type CTestErr struct {
	SNCertif string
	NTalon   string
	Comment  string
}
type CTest struct {
	errs []CTestErr
}

func (a *CTest) UsePair() bool {
	return true
}

func (a *CTest) PutOne(ti TalonInfo) {
	//pass
}

func (a *CTest) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if (t1.NTalon == "1") && (t2.NTalon == "1") && !reflect.DeepEqual(t1, t2) {
		a.errs = append(a.errs, CTestErr{t1.SCertif + " " + t1.NCertif,
			t1.NTalon,
			t1.LName + " " + t1.FName + " " + t1.MName + " (" + t1.Filename + ")" +
				" / " +
				t2.LName + " " + t2.FName + " " + t2.MName + " (" + t2.Filename + ")"})
	}
}

func (a *CTest) Header() []string {
	return []string{"Серия номер сертификата", "Номер талона", "Комментарий"}
}

func (a *CTest) Len() int {
	return len(a.errs)
}

func (a *CTest) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].NTalon, a.errs[i].Comment}
}
*/

// Двойные номера сертификатов
type DoubleNCertifErr struct {
	SNCertif string
	NTalon   string
	Comment  string
}
type DoubleNCertif struct {
	errs []DoubleNCertifErr
}

func (a *DoubleNCertif) UsePair() bool {
	return true
}

func (a *DoubleNCertif) PutOne(ti TalonInfo) {
	//pass
}

func (a *DoubleNCertif) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if (t1.NTalon == t2.NTalon) && (t1.SCertif == t2.SCertif) &&
		(t1.NCertif == t2.NCertif) && !reflect.DeepEqual(t1, t2) {
		a.errs = append(a.errs, DoubleNCertifErr{t1.SCertif + " " + t1.NCertif,
			t1.NTalon,
			t1.LName + " " + t1.FName + " " + t1.MName + " (" + t1.Filename + ")" +
				" <br/> " +
				t2.LName + " " + t2.FName + " " + t2.MName + " (" + t2.Filename + ")"})
	}
}

func (a *DoubleNCertif) Header() []string {
	return []string{"Серия номер сертификата", "Номер талона", "Комментарий"}
}

func (a *DoubleNCertif) Len() int {
	return len(a.errs)
}

func (a *DoubleNCertif) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].NTalon, a.errs[i].Comment}
}

///////////////
// Двойные СНИЛС
type DoubleSnilsErr struct {
	Snils   string
	NTalon  string
	Comment string
}
type DoubleSnils struct {
	errs []DoubleSnilsErr
}

func (a *DoubleSnils) UsePair() bool {
	return true
}

func (a *DoubleSnils) PutOne(ti TalonInfo) {
	//pass
}

func (a *DoubleSnils) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if (t1.NTalon == t2.NTalon) && (t1.Snils > "") && (t1.Snils == t2.Snils) &&
		((t1.LName != t2.LName) || (t1.FName != t2.FName) || (t1.MName != t2.MName)) {
		a.errs = append(a.errs, DoubleSnilsErr{t1.Snils,
			t1.NTalon,
			t1.SCertif + " " + t1.NCertif + ", " + t1.LName + " " + t1.FName + " " + t1.MName + " (" + t1.Filename + ")" +
				" <br/> " +
				t2.SCertif + " " + t2.NCertif + ", " + t2.LName + " " + t2.FName + " " + t2.MName + " (" + t2.Filename + ")"})
	}
}

func (a *DoubleSnils) Header() []string {
	return []string{"СНИЛС", "Номер талона", "Комментарий"}
}

func (a *DoubleSnils) Len() int {
	return len(a.errs)
}

func (a *DoubleSnils) Get(i int) []string {
	return []string{a.errs[i].Snils, a.errs[i].NTalon, a.errs[i].Comment}
}

///////////////
// Разные ФИО, ДР
type FIODifErr struct {
	SNCertif string
	Comment  string
}
type FIODif struct {
	errs []FIODifErr
}

func (a *FIODif) UsePair() bool {
	return true
}

func (a *FIODif) PutOne(ti TalonInfo) {
	//pass
}

func (a *FIODif) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if (t1.SCertif == t2.SCertif) && (t1.NCertif == t2.NCertif) &&
		((t1.LName != t2.LName) || (t1.FName != t2.FName) || (t1.MName != t2.MName) || (t1.BDate != t2.BDate)) {
		inf1 := "Номер талона " + t1.NTalon + " "
		inf2 := "Номер талона " + t2.NTalon + " "
		if t1.LName != t2.LName {
			inf1 += "<b>" + t1.LName + "</b> "
			inf2 += "<b>" + t2.LName + "</b> "
		} else {
			inf1 += t1.LName + " "
			inf2 += t2.LName + " "
		}
		if t1.FName != t2.FName {
			inf1 += "<b>" + t1.FName + "</b> "
			inf2 += "<b>" + t2.FName + "</b> "
		} else {
			inf1 += t1.FName + " "
			inf2 += t2.FName + " "
		}
		if t1.MName != t2.MName {
			inf1 += "<b>" + t1.MName + "</b> "
			inf2 += "<b>" + t2.MName + "</b> "
		} else {
			inf1 += t1.MName + " "
			inf2 += t2.MName + " "
		}
		if t1.BDate != t2.BDate {
			inf1 += "<b>" + t1.BDate + "</b> "
			inf2 += "<b>" + t2.BDate + "</b> "
		} else {
			inf1 += t1.BDate + " "
			inf2 += t2.BDate + " "
		}
		inf1 += "(" + t1.Filename + ")"
		inf2 += "(" + t2.Filename + ")"
		a.errs = append(a.errs, FIODifErr{t1.SCertif + " " + t1.NCertif,
			inf1 + " <br/> " + inf2})
	}
}

func (a *FIODif) Header() []string {
	return []string{"Серия номер сертификата", "Комментарий"}
}

func (a *FIODif) Len() int {
	return len(a.errs)
}

func (a *FIODif) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].Comment}
}

///////////////
// Разные СНИЛС
type SnilsDifErr struct {
	SNCertif string
	Comment  string
}
type SnilsDif struct {
	errs []SnilsDifErr
}

func (a *SnilsDif) UsePair() bool {
	return true
}

func (a *SnilsDif) PutOne(ti TalonInfo) {
	//pass
}

func (a *SnilsDif) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if (t1.SCertif == t2.SCertif) && (t1.NCertif == t2.NCertif) &&
		(t1.Snils != t2.Snils) {
		inf1 := "Номер талона " + t1.NTalon + " "
		inf2 := "Номер талона " + t2.NTalon + " "
		inf1 += "<b>" + t1.Snils + "</b> "
		inf2 += "<b>" + t2.Snils + "</b> "
		inf1 += t1.LName + " "
		inf2 += t2.LName + " "
		inf1 += t1.FName + " "
		inf2 += t2.FName + " "
		inf1 += t1.MName + " "
		inf2 += t2.MName + " "
		inf1 += t1.BDate + " "
		inf2 += t2.BDate + " "
		inf1 += "(" + t1.Filename + ")"
		inf2 += "(" + t2.Filename + ")"
		a.errs = append(a.errs, SnilsDifErr{t1.SCertif + " " + t1.NCertif,
			inf1 + " <br/> " + inf2})
	}
}

func (a *SnilsDif) Header() []string {
	return []string{"Серия номер сертификата", "Комментарий"}
}

func (a *SnilsDif) Len() int {
	return len(a.errs)
}

func (a *SnilsDif) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].Comment}
}

///////////////
// Разные паспортные данные
type PassportDifErr struct {
	SNCertif string
	Comment  string
}
type PassportDif struct {
	errs []PassportDifErr
}

func (a *PassportDif) UsePair() bool {
	return true
}

func (a *PassportDif) PutOne(ti TalonInfo) {
	//pass
}

func (a *PassportDif) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if (t1.SCertif == t2.SCertif) && (t1.NCertif == t2.NCertif) &&
		((t1.PassportTDoc != t2.PassportTDoc) ||
			(t1.PassportNDoc != t2.PassportNDoc) ||
			(t1.PassportSDoc != t2.PassportSDoc) ||
			(t1.PassportDDoc != t2.PassportDDoc) ||
			(t1.PassportODoc != t2.PassportODoc)) {
		inf1 := fmt.Sprintf("%s %s %s (%s)", t1.LName, t1.FName, t1.MName, t1.Filename)
		inf2 := fmt.Sprintf("%s %s %s (%s)", t2.LName, t2.FName, t2.MName, t2.Filename)

		if t1.PassportTDoc != t2.PassportTDoc {
			inf1 += "<b>" + t1.PassportTDocName + "</b> "
			inf2 += "<b>" + t2.PassportTDocName + "</b> "
		} else {
			inf1 += t1.PassportTDocName + " "
			inf2 += t2.PassportTDocName + " "
		}
		if t1.PassportSDoc != t2.PassportSDoc {
			inf1 += "<b>" + t1.PassportSDoc + "</b> "
			inf2 += "<b>" + t2.PassportSDoc + "</b> "
		} else {
			inf1 += t1.PassportSDoc + " "
			inf2 += t2.PassportSDoc + " "
		}
		if t1.PassportNDoc != t2.PassportNDoc {
			inf1 += "<b>" + t1.PassportNDoc + "</b> "
			inf2 += "<b>" + t2.PassportNDoc + "</b> "
		} else {
			inf1 += t1.PassportNDoc + " "
			inf2 += t2.PassportNDoc + " "
		}
		if t1.PassportDDoc != t2.PassportDDoc {
			inf1 += "<b>" + t1.PassportDDoc + "</b> "
			inf2 += "<b>" + t2.PassportDDoc + "</b> "
		} else {
			inf1 += t1.PassportDDoc + " "
			inf2 += t2.PassportDDoc + " "
		}
		if t1.PassportODoc != t2.PassportODoc {
			inf1 += "<b>" + t1.PassportODoc + "</b> "
			inf2 += "<b>" + t2.PassportODoc + "</b> "
		} else {
			inf1 += t1.PassportODoc + " "
			inf2 += t2.PassportODoc + " "
		}
		a.errs = append(a.errs, PassportDifErr{t1.SCertif + " " + t1.NCertif,
			inf1 + " <br/> " + inf2})
	}
}

func (a *PassportDif) Header() []string {
	return []string{"Серия номер сертификата", "Комментарий"}
}

func (a *PassportDif) Len() int {
	return len(a.errs)
}

func (a *PassportDif) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].Comment}
}

///////////////
// Разные даты сертификатов
type DCertifDifErr struct {
	SNCertif string
	Comment  string
}
type DCertifDif struct {
	errs []DCertifDifErr
}

func (a *DCertifDif) UsePair() bool {
	return true
}

func (a *DCertifDif) PutOne(ti TalonInfo) {
	//pass
}

func (a *DCertifDif) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if (t1.SCertif == t2.SCertif) && (t1.NCertif == t2.NCertif) &&
		(t1.DCertif != t2.DCertif) {
		inf1 := fmt.Sprintf("<b>%s</b> %s %s %s (%s)", t1.DCertif, t1.LName, t1.FName, t1.MName, t1.Filename)
		inf2 := fmt.Sprintf("<b>%s</b> %s %s %s (%s)", t2.DCertif, t2.LName, t2.FName, t2.MName, t2.Filename)

		a.errs = append(a.errs, DCertifDifErr{t1.SCertif + " " + t1.NCertif,
			inf1 + " <br/> " + inf2})
	}
}

func (a *DCertifDif) Header() []string {
	return []string{"Серия номер сертификата", "Комментарий"}
}

func (a *DCertifDif) Len() int {
	return len(a.errs)
}

func (a *DCertifDif) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].Comment}
}

///////////////
// Разные обменные карты
type CardDifErr struct {
	SNCertif string
	Comment  string
}
type CardDif struct {
	errs []CardDifErr
}

func (a *CardDif) UsePair() bool {
	return true
}

func (a *CardDif) PutOne(ti TalonInfo) {
	//pass
}

func (a *CardDif) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if (t1.SCertif == t2.SCertif) && (t1.NCertif == t2.NCertif) &&
		((t1.NCard != t2.NCard) || (t1.DCard != t2.DCard)) {
		inf1 := fmt.Sprintf("Обм карта <b>%s %s</b> %s %s %s (%s)", t1.NCard, t1.DCard, t1.LName, t1.FName, t1.MName, t1.Filename)
		inf2 := fmt.Sprintf("Обм карта <b>%s %s</b> %s %s %s (%s)", t2.NCard, t2.DCard, t2.LName, t2.FName, t2.MName, t2.Filename)

		a.errs = append(a.errs, CardDifErr{t1.SCertif + " " + t1.NCertif,
			inf1 + " <br/> " + inf2})
	}
}

func (a *CardDif) Header() []string {
	return []string{"Серия номер сертификата", "Комментарий"}
}

func (a *CardDif) Len() int {
	return len(a.errs)
}

func (a *CardDif) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].Comment}
}

///////////////
// Разные страховые полисы
type PolicyDifErr struct {
	SNCertif string
	Comment  string
}
type PolicyDif struct {
	errs []PolicyDifErr
}

func (a *PolicyDif) UsePair() bool {
	return true
}

func (a *PolicyDif) PutOne(ti TalonInfo) {
	//pass
}

func (a *PolicyDif) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if (t1.SCertif == t2.SCertif) && (t1.NCertif == t2.NCertif) &&
		((t1.NPolicy != t2.NPolicy) || (t1.SPolicy != t2.SPolicy)) {
		inf1 := fmt.Sprintf("Полис <b>%s %s</b> %s %s %s (%s)", t1.SPolicy, t1.NPolicy, t1.LName, t1.FName, t1.MName, t1.Filename)
		inf2 := fmt.Sprintf("Полис <b>%s %s</b> %s %s %s (%s)", t2.SPolicy, t2.NPolicy, t2.LName, t2.FName, t2.MName, t2.Filename)

		a.errs = append(a.errs, PolicyDifErr{t1.SCertif + " " + t1.NCertif,
			inf1 + " <br/> " + inf2})
	}
}

func (a *PolicyDif) Header() []string {
	return []string{"Серия номер сертификата", "Комментарий"}
}

func (a *PolicyDif) Len() int {
	return len(a.errs)
}

func (a *PolicyDif) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].Comment}
}

///////////////
// Разные листки нетрудоспособности
type LeafDifErr struct {
	SNCertif string
	Comment  string
}
type LeafDif struct {
	errs []LeafDifErr
}

func (a *LeafDif) UsePair() bool {
	return true
}

func (a *LeafDif) PutOne(ti TalonInfo) {
	//pass
}

func (a *LeafDif) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if ((t1.NTalon == "1") && (t2.NTalon == "2") || (t1.NTalon == "2") && (t2.NTalon == "1")) &&
		(t1.SCertif == t2.SCertif) && (t1.NCertif == t2.NCertif) &&
		((t1.SLeaf != t2.SLeaf) || (t1.NLeaf != t2.NLeaf) || (t1.DLeaf != t2.DLeaf)) {
		inf1 := "Листок: &lt;" + t1.SLeaf
		inf2 := "Листок: &lt;" + t2.SLeaf
		if t1.NLeaf != t2.NLeaf {
			inf1 += " <b>" + t1.NLeaf + "</b>"
			inf2 += " <b>" + t2.NLeaf + "</b>"
		} else {
			inf1 += " " + t1.NLeaf
			inf2 += " " + t2.NLeaf
		}
		if t1.DLeaf != t2.DLeaf {
			inf1 += " <b>" + t1.DLeaf + "</b>"
			inf2 += " <b>" + t2.DLeaf + "</b>"
		} else {
			inf1 += " " + t1.DLeaf
			inf2 += " " + t2.DLeaf
		}
		inf1 = fmt.Sprintf("%s &gt; %s %s %s (%s)", inf1, t1.LName, t1.FName, t1.MName, t1.Filename)
		inf2 = fmt.Sprintf("%s &gt; %s %s %s (%s)", inf2, t2.LName, t2.FName, t2.MName, t2.Filename)

		a.errs = append(a.errs, LeafDifErr{t1.SCertif + " " + t1.NCertif,
			inf1 + " <br/> " + inf2})
	}
}

func (a *LeafDif) Header() []string {
	return []string{"Серия номер сертификата", "Комментарий"}
}

func (a *LeafDif) Len() int {
	return len(a.errs)
}

func (a *LeafDif) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].Comment}
}

///////////////
// Пересечение дат 3-1 и 3-2
type Intersect3Err struct {
	SNCertif string
	Comment  string
}
type Intersect3 struct {
	errs []Intersect3Err
}

func (a *Intersect3) UsePair() bool {
	return true
}

func (a *Intersect3) PutOne(ti TalonInfo) {
	//pass
}

func (a *Intersect3) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if ((t1.NTalon == "3-1") && (t2.NTalon == "3-2") || (t1.NTalon == "3-2") && (t2.NTalon == "3-1")) &&
		(t1.SCertif == t2.SCertif) && (t1.NCertif == t2.NCertif) {
		for _, c1 := range t1.Childs {
			for _, c2 := range t2.Childs {
				if (c1.DateBobs <= c2.DateEobs) && (c1.DateBobs >= c2.DateBobs) ||
					(c2.DateBobs <= c1.DateEobs) && (c2.DateBobs >= c1.DateBobs) {
					inf1 := fmt.Sprintf("%s %s %s %s <b>%s - %s</b> %s %s %s (%s)",
						t1.NTalon, c1.LName, c1.FName, c1.MName, c1.DateBobs, c1.DateEobs,
						t1.LName, t1.FName, t1.MName, t1.Filename)
					inf2 := fmt.Sprintf("%s %s %s %s <b>%s - %s</b> %s %s %s (%s)",
						t2.NTalon, c2.LName, c2.FName, c2.MName, c2.DateBobs, c2.DateEobs,
						t2.LName, t2.FName, t2.MName, t2.Filename)
					a.errs = append(a.errs, Intersect3Err{t1.SCertif + " " + t1.NCertif,
						inf1 + " <br/> " + inf2})

				}
			}
		}

	}
}

func (a *Intersect3) Header() []string {
	return []string{"Серия номер сертификата", "Комментарий"}
}

func (a *Intersect3) Len() int {
	return len(a.errs)
}

func (a *Intersect3) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].Comment}
}

///////////////
// Дата постановки = дата родов
type PobsErr struct {
	SNCertif string
	Comment  string
}
type Pobs struct {
	errs []PobsErr
}

func (a *Pobs) UsePair() bool {
	return false
}

func (a *Pobs) PutOne(ti TalonInfo) {
	if (ti.NTalon == "3-1") || (ti.NTalon == "3-2") {
		for _, c1 := range ti.Childs {
			if c1.DatePobs <= c1.BDate {
				a.errs = append(a.errs, PobsErr{ti.SCertif + " " + ti.NCertif,
					fmt.Sprintf("%s %s %s %s <b>ДПост %s <= ДР %s</b> %s %s %s (%s)",
						ti.NTalon, c1.LName, c1.FName, c1.MName, c1.DatePobs, c1.BDate,
						ti.LName, ti.FName, ti.MName, ti.Filename)})
			}
		}
	}
}

func (a *Pobs) PutPair(t1 TalonInfo, t2 TalonInfo) {
	//pass
}

func (a *Pobs) Header() []string {
	return []string{"Серия номер сертификата", "Комментарий"}
}

func (a *Pobs) Len() int {
	return len(a.errs)
}

func (a *Pobs) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].Comment}
}

///////////////
// Разные ФИО, дата рождения, дата постановки в 3-1 и 3-2
type Dif3Err struct {
	SNCertif string
	Comment  string
}
type Dif3 struct {
	errs []Dif3Err
}

func (a *Dif3) UsePair() bool {
	return true
}

func (a *Dif3) PutOne(ti TalonInfo) {
	//pass
}

func cmpChild(c1 ChildInfo, c2 ChildInfo) bool {
	return (c1.LName == c2.LName) && (c1.FName == c2.FName) &&
		(c1.MName == c2.MName) && (c1.BDate == c2.BDate) &&
		(c1.DatePobs == c2.DatePobs)
}

func (a *Dif3) PutPair(t1 TalonInfo, t2 TalonInfo) {
	if ((t1.NTalon == "3-1") && (t2.NTalon == "3-2") || (t1.NTalon == "3-2") && (t2.NTalon == "3-1")) &&
		(t1.SCertif == t2.SCertif) && (t1.NCertif == t2.NCertif) {
		inf1 := ""
		for _, c1 := range t1.Childs {
			report := true
			for _, c2 := range t2.Childs {
				if cmpChild(c1, c2) {
					report = false
				}
			}
			if report {
				inf1 += fmt.Sprintf("%s <b>%s %s %s ДР %s ДПост %s</b> %s %s %s (%s)<br/>",
					t1.NTalon, c1.LName, c1.FName, c1.MName, c1.BDate, c1.DatePobs,
					t1.LName, t1.FName, t1.MName, t1.Filename)

			}
		}
		for _, c1 := range t2.Childs {
			report := true
			for _, c2 := range t1.Childs {
				if cmpChild(c1, c2) {
					report = false
				}
			}
			if report {
				inf1 += fmt.Sprintf("%s <b>%s %s %s ДР %s ДПост %s</b> %s %s %s (%s)<br/>",
					t1.NTalon, c1.LName, c1.FName, c1.MName, c1.BDate, c1.DatePobs,
					t1.LName, t1.FName, t1.MName, t1.Filename)
			}
		}
		if inf1 > "" {
			a.errs = append(a.errs, Dif3Err{t1.SCertif + " " + t1.NCertif,
				inf1})
		}
	}
}

func (a *Dif3) Header() []string {
	return []string{"Серия номер сертификата", "Комментарий"}
}

func (a *Dif3) Len() int {
	return len(a.errs)
}

func (a *Dif3) Get(i int) []string {
	return []string{a.errs[i].SNCertif, a.errs[i].Comment}
}
