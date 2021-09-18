package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"strconv"
	"github.com/PuerkitoBio/goquery"
)

func main(){
	fmt.Println("https://internshala.com/internships => scrapy")

	for page := 1; page < 388; page++{
		res, err := http.Get("https://internshala.com/internships/page-" +  strconv.Itoa(page))

		if err != nil {
			log.Fatal(err)
		}
	
		defer res.Body.Close()
	
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}
	
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}
	
		fmt.Println(doc)
	
		doc.Find("div.container-fluid.individual_internship.visibilityTrackerItem").Each(func(i int, s *goquery.Selection) { 
			company_tip := s.Find("div.heading_4_5.profile a").Text()
			company_name := s.Find("div.heading_6.company_name").Text()
			company_name = strings.TrimSpace(company_name)
			location := s.Find("a.location_link").Text()
			stipend := s.Find("span.stipend").Text()
			stipend = strings.TrimSpace(stipend)
			apply_by := s.Find("div.other_detail_item.apply_by div.item_body").Text()
	
			fmt.Printf("Company Tip => %s\n", company_tip)
			fmt.Printf("Company 	=> %s\n", company_name)
			fmt.Printf("Location 	=> %s\n", location)
			fmt.Printf("Stipend 	=> %s\n", stipend)
			fmt.Printf("Apply by 	=> %s\n", apply_by)
			fmt.Println("*************")
		});
	}

	
}