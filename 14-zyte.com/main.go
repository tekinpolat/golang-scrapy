package main 


import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"

	"database/sql"
	_ "github.com/lib/pq"
)

const BASE_URL = "https://www.zyte.com"


const (
	DB_USER 	= "tknplt"
	DB_PASSWORD = "2121"
	DB_ADDR 	= "localhost"
	DB_NAME 	= "scrapy_go"
	DB_PORT 	= 5432
)

func init() {
	fmt.Println("Starting...")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main()  {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()
	
	res, err := http.Get(BASE_URL + "/blog/")
    if err != nil {
        log.Fatalf("Error fetching url %s: %v", BASE_URL, err)
    }

    defer res.Body.Close()

	fmt.Println(res)


	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)

	page_count_str := doc.Find("div.oxy-easy-posts-pages .page-numbers:nth-child(5)").Text()
	page_count,_ := strconv.Atoi(page_count_str) 
	fmt.Println(page_count)

	var counter int = 1
	for i := 2; i <= page_count; i++ {
		color.Blue("Page: %d\n", i)
		res, err := http.Get(BASE_URL + "/blog/page/" + strconv.Itoa(i))
		if err != nil {
			log.Fatalf("Error fetching url %s: %v", BASE_URL, err)
		}

		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("div.oxy-post").Each(func(i int, s *goquery.Selection) {
			date := strings.TrimSpace(s.Find("div.oxy-post-image-date-overlay").Text())
			title := strings.TrimSpace(s.Find("a.oxy-post-title").Text())
			author := strings.TrimSpace(s.Find("div.oxy-post-meta-author.oxy-post-meta-item:first-child").Text())
			link := BASE_URL + s.Find("a.oxy-read-more").AttrOr("href", "")

			minutes := strings.TrimSpace(s.Find("div.oxy-post-meta-author.oxy-post-meta-item:nth-child(2)").Text())

			color.Green("%d\n", counter)
			fmt.Printf("%s\n", date)
			fmt.Printf("%s\n", title)
			fmt.Printf("%s\n", author)
			fmt.Printf("%s\n", minutes)
			fmt.Printf("%s\n", link)
			fmt.Println("*********")
			counter++

			sql := `INSERT INTO zyte_blog (link, author, title, date, minutes) VALUES ($1, $2, $3, $4, $5)`
			_, err = db.Exec(sql, link, author, title, date, minutes)
			if err != nil {
				panic(err)
			}

		})
	}
}