package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
)

func main()  {
	color.Cyan("=>Go diliyle ulkelerin uluslararasi harf kodlari\n")
	color.Blue("=>https://turev.net/Ulkelerin-Uluslararasi-Harf-Kodlari/\n")
	fmt.Println("Starting...")

	res, err := http.Get("https://turev.net/Ulkelerin-Uluslararasi-Harf-Kodlari/")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
	//fmt.Println(res.Status)
	//fmt.Println(res.Header)
	//fmt.Println(res.ContentLength)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)

	doc.Find("div.news-text table tbody tr").Each(func(i int, s *goquery.Selection) {
		country := s.Find("td:nth-child(1)").Text()
		customs_code := s.Find("td:nth-child(2)").Text()
		country_code := s.Find("td:nth-child(3)").Text()

		fmt.Printf("=>%s %s %s\n", country, customs_code, country_code)
	})

	//link_sql := doc.Find("div.news-text p a").AttrOr("href", "")
	//fmt.Println(link_sql)

}