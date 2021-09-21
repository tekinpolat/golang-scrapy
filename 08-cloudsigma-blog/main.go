package main 


import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
)

func main()  {
	color.Set(color.FgYellow)
	fmt.Println("=>https://www.cloudsigma.com/blog/")
	color.Unset()

	res, err := http.Get("https://www.cloudsigma.com/blog/")

	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)

	page_count, _ := strconv.Atoi(doc.Find("a.last").Text())
	fmt.Printf("%v %T", page_count, page_count)

	var counter int = 1
	for i := 1; i <= page_count; i++ {
		res, err := http.Get("https://www.cloudsigma.com/blog/page/" +strconv.Itoa(i)+ "/")
		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()
		fmt.Println(res)
		doc, err = goquery.NewDocumentFromReader(res.Body)

		if err != nil {
			log.Fatalln(err)
		}

		
		doc.Find("article.post").Each(func(i int, s *goquery.Selection) {
			title := s.Find("header.entry-header h2.entry-title a").Text()
			desc := s.Find("div.entry-content.excerpt p").Text()
			link := s.Find("div.entry-content.excerpt a.more-link").AttrOr("href", "")
			var tags []string

			s.Find("footer.entry-footer.cf a").Each(func(i int, s *goquery.Selection) {
				tags = append(tags, s.Text())
			})
			//fmt.Println("LL")
			color.Set(color.FgGreen)
			fmt.Printf("=>%v %s %s\n", counter, title, link)
			fmt.Printf(" %s\n", desc)
			fmt.Printf(" %s\n", tags)
			color.Unset()
			fmt.Println("***********************")
			counter++
		})
		
	}
}