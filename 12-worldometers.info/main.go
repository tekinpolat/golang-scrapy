package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main(){
	const URL = "https://www.worldometers.info/world-population/population-by-country/"
	fmt.Println("=>",URL)

	res, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	fmt.Println(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
		
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(doc)

	doc.Find("table#example2 tbody tr").Each(func(i int, s *goquery.Selection) {
		country := s.Find("td:nth-child(2)").Text()
		population := s.Find("td:nth-child(3)").Text()
		year_change := s.Find("td:nth-child(4)").Text()
		net_change := s.Find("td:nth-child(5)").Text()
		density := s.Find("td:nth-child(6)").Text()
		land_area := s.Find("td:nth-child(7)").Text()

		fmt.Printf("%d => %s => %s => %s => %s => %s => %s\n", i + 1, 
			country, population, year_change, net_change, density, land_area)
		
	})
}