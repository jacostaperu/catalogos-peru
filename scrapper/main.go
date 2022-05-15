package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type Catalog struct {
	Name        string `json:"n"`
	ValidFrom   int64  `json:"vf"`
	ValidUntil  int64  `json:"vu"`
	Urllink     string `json:"u"`
	ImageIdxMax int    `json:""`
}

var meses map[string]string = map[string]string{
	"enero":     "01",
	"febrero":   "02",
	"marzo":     "03",
	"abril":     "04",
	"mayo":      "05",
	"junio":     "06",
	"julio":     "07",
	"agost":     "08",
	"setiembre": "09",
	"octubre":   "10",
	"noviembre": "11",
	"diciembre": "12",
}

func ScrapeMetro() Catalog {
	c := colly.NewCollector()

	var metro Catalog
	metro.Name = "Metro"

	metro.Urllink = "https://catalogosmetro.metro.pe/catalogo-precios-mas-bajos/"
	metro.ImageIdxMax = 0

	re := regexp.MustCompile(`\d+`)
	reFecha := regexp.MustCompile(`Del\s+(?P<sDay>\d+)\s+de\s+(?P<sMonth>\w+)\s+al\s+(?P<eDay>\d+)\s+de\s+(?P<eMonth>\w+)\s+de\s+(?P<Year>\d+)`)

	c.OnHTML("meta[name='description']", func(e *colly.HTMLElement) {
		fecha_text := e.Attr("content")

		matches := reFecha.FindStringSubmatch(fecha_text)

		sDayIdx := reFecha.SubexpIndex("sDay")
		sMonthIdx := reFecha.SubexpIndex("sMonth")
		eDayIdx := reFecha.SubexpIndex("eDay")
		eMonthIdx := reFecha.SubexpIndex("eMonth")
		yearIdx := reFecha.SubexpIndex("Year")

		sDate := matches[yearIdx] + "-" + meses[matches[sMonthIdx]] + "-" + matches[sDayIdx]
		eDate := matches[yearIdx] + "-" + meses[matches[eMonthIdx]] + "-" + matches[eDayIdx]
		timeS, _ := time.Parse("2006-01-02", sDate)
		timeE, _ := time.Parse("2006-01-02", eDate)
		metro.ValidFrom = timeS.Unix()  // TODO get from the web page
		metro.ValidUntil = timeE.Unix() // TODO get from the web page

		//fmt.Println(sDate, timeS.Unix(), eDate, timeE.Unix())

	})

	c.OnHTML("div.contenido-interno.juntar", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))

		e.DOM.Find("img").Each(func(i int, ee *goquery.Selection) {
			src, ok := ee.Attr("src")

			if ok {
				//fmt.Println(e.Request.URL.Scheme + "://" + e.Request.URL.Host + e.Request.URL.Path + src)
				idxStr := string(re.Find([]byte(src)))
				idxInt, _ := strconv.Atoi(idxStr)

				if idxInt > metro.ImageIdxMax {
					metro.ImageIdxMax = idxInt
				}
			}

		})

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://catalogosmetro.metro.pe/catalogo-precios-mas-bajos/")

	return metro

}

func main() {

	// Find and visit all links
	metro := ScrapeMetro()
	jsonData, _ := json.Marshal(metro)
	fmt.Println(string(jsonData))

}
