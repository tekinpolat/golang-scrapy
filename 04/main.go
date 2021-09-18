package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func main(){
	fmt.Println("https://en.wikipedia.org/wiki/List_of_Nobel_laureates => scrapy")

	res, err := http.Get("https://en.wikipedia.org/wiki/List_of_Nobel_laureates")
	
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)
	
	doc.Find("table.wikitable tbody tr").Each(func(i int, s *goquery.Selection) {
		
		year := strings.TrimSpace(s.Find("td:nth-child(1)").Text())

		var  physics []string
		s.Find("td:nth-child(2) span.fn a").Each(func(i int, s *goquery.Selection) {
			physics = append(physics, strings.TrimSpace(s.Text()))
		})

		var chemistry []string
		s.Find("td:nth-child(3) span.fn a").Each(func(i int, s *goquery.Selection) {
			chemistry = append(chemistry, strings.TrimSpace(s.Text()))
		})

		var physics_or_medicine []string
		s.Find("td:nth-child(4) span.fn a").Each(func(i int, s *goquery.Selection) {
			physics_or_medicine = append(physics_or_medicine, strings.TrimSpace(s.Text()))
		})

		var literature [] string
		s.Find("td:nth-child(5) span.fn a").Each(func(i int, s *goquery.Selection) {
			literature = append(literature, strings.TrimSpace(s.Text()))
		})

		var peace [] string
		s.Find("td:nth-child(6) span.fn a").Each(func(i int, s *goquery.Selection) {
			peace = append(peace, strings.TrimSpace(s.Text()))
		})

		var economics []string
		s.Find("td:nth-child(7) span.fn a").Each(func(i int, s *goquery.Selection) {
			economics = append(economics, strings.TrimSpace(s.Text()))
		})


		fmt.Printf("%s\n", year)
		fmt.Printf("Physics => %v %T\n",physics, physics)
		fmt.Printf("Chemistry => %v\n",chemistry)
		fmt.Printf("Physics Or Medicine => %v\n",physics_or_medicine)
		fmt.Printf("Literature => %v\n",literature)
		fmt.Printf("Peace => %v\n",peace)
		fmt.Printf("Economics => %v\n",economics)
		fmt.Println("************************")
	})

}