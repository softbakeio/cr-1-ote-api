package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"time"
)

const(
	//HtmlOteDailyElectricityUrl source HTML url of daily electricity data
	HtmlOteDailyElectricityUrl = "https://www.ote-cr.cz/cs/kratkodobe-trhy/elektrina/denni-trh"
)
//NewHtmlElectricityDataService new instance of HtmlElectricityDataService
func NewHtmlElectricityDataService() ElectricityDataService {
	newInstance := new(HtmlElectricityDataService)
	return newInstance
}

type HtmlElectricityDataService struct {}

func currentDate() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
}


func (heds *HtmlElectricityDataService) GetElectricityHourData() (ehd *ElectricityHourData, err error) {

	apiUrl := fmt.Sprintf("%s?date=%s", HtmlOteDailyElectricityUrl, currentDate())
	data := make([]float64, 0)

	// Request the HTML page.
	res, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("error get html content from OTE")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}


	// Find the review items
	tables := doc.Find(".report_table").Each(func (i int, s *goquery.Selection) {})

	if tables.Length() == 0 {
		return nil, errors.New("not table data found")
	}

	tables.Each(func(i int, s *goquery.Selection) {

		// its second report table
		if i == 1 {
			s.Find("tbody tr").Each(func(i1 int, s1 *goquery.Selection) {
				if i1 >= 0 && i1 < 24 {
					s1.Find("td").Each(func(i2 int, s2 *goquery.Selection) {
						if i2 == 0 {
							num, _ := strconv.ParseFloat(NormalizeAmerican(s2.Text()), 64)
							data = append(data, num)
						}
					})
				}
			})
		}
	})

	if err != nil {
		return nil, err
	}

	ehd = CreateElectricityHourData(data)

	return ehd, nil
}