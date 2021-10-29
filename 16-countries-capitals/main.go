package main 

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	
	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
)

type CountryCapital struct{
	Country string `json:"country"`
	Capital string `json:"capital"`
}

const URL = "https://geographyfieldwork.com/WorldCapitalCities.htm"

func init(){
	color.Cyan(URL)
}

func main(){
	resp, err := http.Get(URL)
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp)
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)

	var result []CountryCapital
	doc.Find("table#anyid  tbody tr").Each(func(i int, s *goquery.Selection) {
		country := s.Find("td:nth-child(1)").Text()
		if i != 0 && country != "200" {
			
			capital := s.Find("td:nth-child(2)").Text()
			/*
			if country == capital {
				fmt.Println("Ülke ile Başkent isimleri aynı =>", country)
			}
			*/

			//fmt.Println(i,country, capital)
			//fmt.Println("****************")

			result = append(result, CountryCapital{Country:country, Capital:capital})
		}
		
	})

	//fmt.Printf("%#v\n", result)
	resultJson, _ := json.MarshalIndent(result, "", "\t")

	fmt.Println(string(resultJson))
}