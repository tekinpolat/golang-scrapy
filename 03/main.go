package main 
import (
	"log"
	"fmt"
	"strings"
	"strconv"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)


func main(){
	for page:= 1; page < 18; page++{
		res, err := http.Get("https://coreyms.com/page/" + strconv.Itoa(page))
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(doc)

		doc.Find("article.post").Each(func(i int, s *goquery.Selection) {
			//link := s.Find("a.entry-title-link")
			title := strings.TrimSpace(s.Find("a.entry-title-link").Text())
			desc := s.Find("div.entry-content p").Text()
			time := s.Find("time.entry-time").Text()
			author := s.Find("span.entry-author a").Text()
			comment_count := s.Find("span.entry-comments-link a").Text()
			tags := ""
			categories := ""
			s.Find("span.entry-tags a").Each(func(i int, s *goquery.Selection) {
				tags += strings.TrimSpace(s.Text()) + "-"
			})

			s.Find("span.entry-categories a").Each(func(i int, s *goquery.Selection) {
				categories += strings.TrimSpace(s.Text()) + "-"
			})
			fmt.Printf("%d =>\n", (i+1)*page)
			fmt.Printf("Author => %s\n", author)
			fmt.Printf("Comment Count => %s\n", comment_count)
			fmt.Printf("Title => %s\n", title)
			fmt.Printf("Desc  => %s\n", desc)
			fmt.Printf("Tags  => %s\n", tags)
			fmt.Printf("Time  => %s\n", time)
			fmt.Printf("Categories  => %s\n", categories)
		});
	}
}