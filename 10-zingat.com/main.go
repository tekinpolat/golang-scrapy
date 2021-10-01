package main 

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"strconv"
	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
	"database/sql"
	"github.com/lib/pq"
)

const BASE_URL = "https://www.zingat.com"

func init(){
	color.Yellow("=>https://www.zingat.com/kiralik-konut")
}

func main()  {
	fmt.Println("Staring...")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",DB_ADDR, DB_PORT, DB_USER, DB_PASSWRD, DB_DB))

	if err != nil {
		log.Fatalln(err)
	}

	var counter int = 1
	for i := 1; i <= 48; i++ {

		var URL string = "https://www.zingat.com/kiralik-konut?page=" + strconv.Itoa(i)


		res, err := http.Get(URL)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()
		//fmt.Println("=>",res)

		doc, err := goquery.NewDocumentFromReader(res.Body)
		
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(doc)

		sql := `INSERT INTO zingat_com (link, price, title, tags, location) VALUES ($1, $2, $3, $4, $5)`

		doc.Find("ul.zc-viewport li.zl-card ").Each(func(i int, s *goquery.Selection) {
			link := s.Find("a.zl-card-inner").AttrOr("href", "")
			link = BASE_URL + link

			price := s.Find("div.feature-item.feature-price").Text()
			title := s.Find("div.zlc-title").Text()
			location := strings.TrimSpace(s.Find("div.zlc-location").Text())
			//image := s.Find("img.zlc-master-image").AttrOr("src", "")

			var tags []string

			s.Find("div.zlc-tags span").Each(func(i int, s *goquery.Selection) {
				tags = append(tags, s.Text())
			})

			fmt.Printf("%d=>%s\n", counter, link)
			fmt.Printf("Price:%s\n", price)
			fmt.Printf("Title:%s\n", title)
			fmt.Printf("Location:%s\n", location)
			//fmt.Printf("Image:%s\n", image)
			fmt.Printf("Tags:%s\n", tags)
			fmt.Println("-----------------------------------")
			counter++

			result, err := db.Exec(sql, link, price, title, pq.Array(tags), location)
			if err != nil {
				log.Println(err)
			}

			log.Println(result)
		})
	}
}