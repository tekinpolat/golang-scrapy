package main 

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/PuerkitoBio/goquery"

	"database/sql"
	"github.com/lib/pq"
)

func main()  {
	fmt.Println("=>http://quotes.toscrape.com Crawling...")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",DB_ADDR, DB_PORT, DB_USER, DB_PASSWRD, DB_DB))

	if err != nil {
		log.Fatalln(err)
	}

	var counter int = 1
	for i := 1; i <= 10; i++ {

		URL := "http://quotes.toscrape.com/page/" + strconv.Itoa(i) +"/"

		res, err := http.Get(URL)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()
		//fmt.Println(res)


		doc, err := goquery.NewDocumentFromReader(res.Body)
			
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(doc)

		sql := `INSERT INTO quotes (quote,author, tags) VALUES ($1, $2, $3)`

		doc.Find("div.quote").Each(func(i int, s *goquery.Selection) {
			quote := s.Find("span.text").Text()
			author := s.Find("small.author").Text()

			var tags []string
			s.Find("div.tags a").Each(func(i int, s *goquery.Selection) {
				tags = append(tags, s.Text())
			})

			result, err := db.Exec(sql, quote, author , pq.Array(tags))
			if err != nil {
				log.Println(err)
			}

			fmt.Println(result)

			fmt.Printf("%d => %s %s\n", counter, quote, author)
			fmt.Printf("%#v\n", tags)
			fmt.Println("*******")
			counter++
		})
	}
}