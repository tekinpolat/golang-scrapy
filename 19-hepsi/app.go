package main

//go mod init scrapy-go
//go mod tidy
//go get github.com/gocolly/colly

//https://developpaper.com/how-to-use-goquery/
import (
	"fmt"
	//"io/ioutil"
	"strings"
	"log"
	"net/http"
	"database/sql"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"db"
)

func main(){

	fmt.Println("Starting")

	db.yazdir()
	//database
	db, err := sql.Open("mysql", "tknplt:Tknplt21.@tcp(127.0.0.1:3306)/hepsi")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()


	resp, err := http.Get("http://discitakvimi.com/")
	if err != nil {
		log.Fatalln(err)
	}

	/*
	fmt.Printf("%T\n",resp.Body)
	fmt.Println(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
 	//Convert the body to type string
	//sb := string(body)
	fmt.Printf("%T\n", body)
	*/


	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
	  	log.Fatal(err)
	}

	sql := "INSERT INTO product(title, content) VALUES (?, ?)"
	doc.Find("section.section.home-tile-section div.col-lg-6.mb-3").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(i)
		title := strings.TrimSpace(s.Find("h4").Text())
		content := strings.TrimSpace(s.Find("div.card.text-center.p-3.border-info p").Text())

		_, err = db.Exec(sql, title, content)
		if err != nil {
			panic(err)
		}

		fmt.Println(i, title,content)

		
	})
}