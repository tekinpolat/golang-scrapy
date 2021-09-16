package main 
import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main(){
	res, err := http.Get("https://imdb.com/chart/top")
	if err != nil {
	  log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	//fmt.Println(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
	  log.Fatal(err)
	}
	
	doc.Find("tbody.lister-list tr").Each(func(i int, s *goquery.Selection) {
		title := s.Find("td.titleColumn a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
	
}