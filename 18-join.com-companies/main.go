package main

import (
	//"fmt"
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
)

type Company struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

const URL = "https://join.com/companies/"

func init(){
	color.Cyan(URL)
}

func main() {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)

	var company []Company

	doc.Find("a.pcd_list_company_link").Each(func(i int, s *goquery.Selection) {
		company_name := s.Text()
		company_link := s.AttrOr("href", "")

		company = append(company, Company{Name:company_name, Link: company_link})

		//fmt.Printf("%d => %s %s\n", i, company,  company_link )
	})

	fmt.Printf("%#v",company)

	resultJson, _ := json.MarshalIndent(company, "", "\t")
	//resultJson, _ := json.Marshal(result)

	fmt.Println(string(resultJson))
	//fmt.Println(resultJson)
}