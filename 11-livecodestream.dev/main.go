package main 

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"strconv"
	"github.com/PuerkitoBio/goquery"
)

func init() {
	fmt.Println("=>https://livecodestream.dev/tags")
}

func main() {
	fmt.Println("Starting...")

	var url string = "https://livecodestream.dev/tags"

	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	fmt.Println(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
		
	if err != nil {
		log.Fatalln(err)
	}

	

	last_page_url := doc.Find("nav.relative.z-0.inline-flex.rounded-md.shadow-sm.-space-x-px a:last-child").AttrOr("href", "")

	s := strings.Split(last_page_url, "/")
	
	page_count, _ := strconv.Atoi(s[3])
	fmt.Printf("Total page count => %d\n", page_count)

	var URLS []string

	for i := 1; i <= page_count; i++{

		if i != 1{
			url = url + "/page/"+ strconv.Itoa(i)+"/"
		}

		res, err = http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		doc, err = goquery.NewDocumentFromReader(res.Body)

		
		if err != nil {
			log.Fatalln(err)
		}

		doc.Find("ul li.flex").Each(func(i int, s *goquery.Selection) {
			link := s.Find("a").AttrOr("href", "")
			link = fmt.Sprintf("https://livecodestream.dev%s", link)
			URLS = append(URLS, link)
		})

	}

	for _, URL := range URLS{
		fmt.Printf("%s scrapy starting...\n", URL)

		res, err := http.Get(URL)
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		doc, err = goquery.NewDocumentFromReader(res.Body)
		
		if err != nil {
			log.Fatalln(err)
		}

		last_page_url = doc.Find("nav.relative a:last-child").AttrOr("href", "")

		s := strings.Split(strings.Trim(last_page_url, "/"), "/")
		//fmt.Println("=>", s, len(s), s[len(s)-2])
		fmt.Printf("%T %+v %#v Length=>%v\n", s, s, s, len(s))

		//if len(s) 
		
	}

}