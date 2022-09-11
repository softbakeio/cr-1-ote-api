package main

import (
	"bytes"
	"fmt"
	"github.com/shakinm/xlsReader/xls"
	"io/ioutil"
	"net/http"
	"time"
)

const(
	ExcelColStartIndex = 0
	ExcelColEndIndex = 1
	ExcelRowStartIndex = 6
	ExcelRowEndIndex = 29
	ExcelOteDailyElectricityUrl = "https://www.ote-cr.cz/pubweb/attachments/01/%d/month%02d/day%02d/DT_%02d_%02d_%d_CZ.xls"
)

//NewExcelElectricityDataService new instance of ExcelElectricityDataService
func NewExcelElectricityDataService() ElectricityDataService {
	newInstance := new(ExcelElectricityDataService)
	return newInstance
}

type ExcelElectricityDataService struct {}

func (eeds *ExcelElectricityDataService) GetElectricityHourData() (ehd *ElectricityHourData, err error) {

	client := http.Client{}
	data := make([]float64, 0)
	year, month, day := time.Now().Date()
	excelFileUrl := fmt.Sprintf(ExcelOteDailyElectricityUrl, year, month, day, day, month, year)

	// get excel file and read it by excel lib
	resp, err := client.Get(excelFileUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	file , err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	workbook, err := xls.OpenReader(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}
	sheet, err := workbook.GetSheet(0)
	if err != nil {
		return nil, err
	}

	for rowIndex := ExcelRowStartIndex; rowIndex <= ExcelRowEndIndex; rowIndex++ {
		if row, err := sheet.GetRow(rowIndex); err == nil {
			val, _ := row.GetCol(1)
			data = append(data, val.GetFloat64())
		}
	}

	ehd = CreateElectricityHourData(data)

	return ehd, nil
}

