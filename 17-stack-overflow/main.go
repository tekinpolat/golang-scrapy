package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"strconv"
	"time"
	"database/sql"

	"github.com/lib/pq"
	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "tknplt"
	password = "2121"
	dbname   = "scrapy_go"
  )
  

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
  	defer db.Close()

	var url string 
	var counter int = 1
	sql := `INSERT INTO stack_overflow (question_link, question_title, question_desc, vote_count, answered_count, view_count, relativetime, "user", tags) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	
	for i := 435921; i >= 1 ; i-- {
		url = "https://stackoverflow.com/questions?tab=newest&page=" + strconv.Itoa(i)
		fmt.Println("=>", url)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("<=", doc)

		
		doc.Find("div#questions div.question-summary").Each(func(i int, s *goquery.Selection) {
			var tags []string
			question_link := s.Find("h3 a.question-hyperlink").AttrOr("href", "")
			question_title := strings.TrimSpace(s.Find("h3 a.question-hyperlink").Text())
			question_desc := strings.TrimSpace(s.Find("div.excerpt").Text())
			vote_count := strings.TrimSpace(s.Find("span.vote-count-post").Text())
			answered_count := strings.TrimSpace(s.Find("div.status.unanswered strong").Text())
			if answered_count == ""{
				answered_count = "0"
			}
			view_count := strings.TrimSpace(s.Find("div.views").Text())
			relativetime := strings.TrimSpace(s.Find("span.relativetime").Text())
			user := strings.TrimSpace(s.Find("div.user-details a").Text())

			s.Find("a.post-tag.flex--item").Each(func(i int, s *goquery.Selection) {
				tags = append(tags, strings.TrimSpace(s.Text()))
			})

			fmt.Printf("%d: %s\n", counter, question_title)
			fmt.Printf("\t%s\n", question_desc)
			fmt.Printf("\t%s\n", question_link)
			fmt.Printf("\t%s Votes - Answer %s \n", vote_count, answered_count)
			fmt.Printf("\t%s Views\n", view_count)
			fmt.Printf("\t%s\n", relativetime)
			fmt.Printf("\t%s\n", user)
			fmt.Printf("\t%#v\n", tags)
			fmt.Println("*********************************")
			counter++

			_, err = db.Exec(sql, question_link, question_title, question_desc, vote_count, answered_count, view_count, relativetime, user, pq.Array(tags))
			if err != nil {
				color.Red("Hata %v", err)
			}

		})

		time.Sleep(10 * time.Second)
	}
}