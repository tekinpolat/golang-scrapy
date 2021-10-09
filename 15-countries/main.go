package main 

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
)

const URL = "https://history.state.gov/countries/all"
//const URL = "http://discitakvimi.com"

type Country struct{
	Name string `json:"name"`
}

func init(){
	color.Cyan("=>" + URL)
}

func main(){
	resp, err := http.Get(URL)

	if err != nil{
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}



	fmt.Println(resp)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
	  	log.Fatal(err)
	}

	fmt.Println(doc)

	var Countries []Country
	doc.Find("div.col-md-6 ul li").Each(func(i int, s *goquery.Selection) {
		country := s.Text()
		//fmt.Println(country) 

		Countries = append(Countries, Country{Name:country})
	});

	//fmt.Printf("%#v", Countries)

	country_json , _ := json.MarshalIndent(Countries, "", "\t")
	fmt.Println(string(country_json))
}