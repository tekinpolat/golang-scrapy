package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"net/http"
	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
	"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)


func init(){
	fmt.Println("Starting....")
}

func main()  {
	
	//db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", DB_USER, DB_PASSWRD, DB_ADDR, DB_DB))

	//db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s  sslmode=disable",DB_ADDR, DB_PORT, DB_USER, DB_PASSWRD, DB_DB))
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",DB_ADDR, DB_PORT, DB_USER, DB_PASSWRD, DB_DB))


	
	if err != nil {
		log.Fatalln(err)
	}
	

	//defer db.Close()

	var url string = "https://www.zappos.com/women-sneakers-athletic-shoes/CK_XARC81wHAAQHiAgMBAhg.zso"
	color.Set(color.FgYellow)
	fmt.Printf("=> %s\n",url)
	color.Unset() 	

	var counter int = 1
	for i:= 0; i <= 104;i++{
		res, err := http.Get(url + "?p="+ strconv.Itoa(i))
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

		sql := `INSERT INTO zappos_com (link, price, description, brand, comment_count, star) VALUES ($1, $2, $3, $4, $5, $6)`
	
		doc.Find("article.Xx-z").Each(func(i int, s *goquery.Selection) {
			fmt.Println("hello")
			link := s.Find("a.Jw-z").AttrOr("href", "")
			link = fmt.Sprintf("https://www.zappos.com%s", link)
			price := s.Find("span.Gz-z").Text()
			desc := s.Find("dd.Mw-z").Text()
			brand := s.Find("dd.Lw-z").Text()
			star := s.Find("span.Eh-z").Text()
			comment_count := strings.TrimRight(s.Find("span.Fh-z").Text(),")")
			comment_count = strings.TrimLeft(comment_count,"(")

			if comment_count == "" {
				comment_count = "0"
			}

			if star == "" {
				star = "0"
			}
	
	
			fmt.Printf("%v=>%s\n", counter, link)
			fmt.Printf("Price: %s\n", price)
			fmt.Printf("Desc: %s\n", desc)
			fmt.Printf("Brand: %s\n", brand)
			fmt.Printf("Star: %s\n", star)
			fmt.Printf("Comment Count: %s\n", comment_count)
			fmt.Println("*******************")
			counter++

			
			result, err := db.Exec(sql, link, price, desc, brand, comment_count, star)
			if err != nil {
				log.Println(err)
			}

			log.Println(result)

		})
	}

	
}