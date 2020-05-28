package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var CHECKS = []struct {
	checkname string
	checker   Accumulator
}{
	{
		"Двойные номера сертификатов",
		&DoubleNCertif{make([]DoubleNCertifErr, 0)},
	},
	{
		"Двойные СНИЛСы",
		&DoubleSnils{make([]DoubleSnilsErr, 0)},
	},
	{
		"Разные ФИО, ДР",
		&FIODif{make([]FIODifErr, 0)},
	},
	{
		"Разные СНИЛС",
		&SnilsDif{make([]SnilsDifErr, 0)},
	},
	{
		"Разные паспортные данные",
		&PassportDif{make([]PassportDifErr, 0)},
	},
	{
		"Разные даты сертификатов",
		&DCertifDif{make([]DCertifDifErr, 0)},
	},
	{
		"Разные обменные карты",
		&CardDif{make([]CardDifErr, 0)},
	},
	{
		"Разные полисы",
		&PolicyDif{make([]PolicyDifErr, 0)},
	},
	{
		"Разные листки нетрудоспособности",
		&LeafDif{make([]LeafDifErr, 0)},
	},
	{
		"Пересечение дат 3-1 и 3-2",
		&Intersect3{make([]Intersect3Err, 0)},
	},
	{
		"Дата постановки = дата родов",
		&Pobs{make([]PobsErr, 0)},
	},
	{
		"Разные ФИО, дата рождения, дата постановки в 3-1 и 3-2",
		&Dif3{make([]Dif3Err, 0)},
	},
}

func main() {
	var talonData = make([]TalonInfo, 0)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	logfile, err := os.Create("Технические ошибки программы.txt")
	if err != nil {
		log.Fatalf("Невозможно создать файл: %v", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recovered: %v", r)
		}
	}()
	//загрузка данных
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && (strings.ToUpper(filepath.Ext(path)) == ".XML") {
			if err != nil {
				log.Printf("Путь: %s, ошибка %v\n", path, err)
			}
			f, err2 := os.Open(path)
			if err2 != nil {
				log.Printf("Ошибка при открытии файла: %s, ошибка %v\n", path, err2)
			}
			defer f.Close()
			res, err2 := Xml2Struct(f, path)
			if err2 != nil {
				log.Printf("Ошибка при разборе файла: %s, ошибка %v\n", path, err2)
			} else {
				talonData = append(talonData, res...)
			}
		}
		return nil
	})
	//прогон проверок
	for _, cinfo := range CHECKS {
		f, err := os.Create(cinfo.checkname + ".html")
		if err != nil {
			log.Printf("Ошибка при создании файла: %s, ошибка %v\n", cinfo.checkname, err)
		}
		defer f.Close()
		if cinfo.checker.UsePair() {
			for i := 0; i < len(talonData); i++ {
				for j := i + 1; j < len(talonData); j++ {
					cinfo.checker.PutPair(talonData[i], talonData[j])
				}
			}
		} else {
			for _, ti := range talonData {
				cinfo.checker.PutOne(ti)
			}
		}
		f.WriteString("<h3>" + cinfo.checkname + "</h3>")
		f.WriteString("<table style=\"border-color:black; border-collapse:collapse\" cellpadding=\"2\" cellspacing=\"1\" border=\"1\"><tr>")
		for _, s := range cinfo.checker.Header() {
			f.WriteString("<td>" + s + "</td>")
		}
		f.WriteString("</tr>")
		for i := 0; i < cinfo.checker.Len(); i++ {
			f.WriteString("<tr>")
			for _, s := range cinfo.checker.Get(i) {
				f.WriteString("<td>" + s + "</td>")
			}
			f.WriteString("</tr>")
		}
		f.WriteString("</table>")
	}
}
