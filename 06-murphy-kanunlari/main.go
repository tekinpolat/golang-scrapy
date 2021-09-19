package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/fatih/color"
	"github.com/PuerkitoBio/goquery"
)

func main(){
	color.Set(color.FgGreen)
	fmt.Println("=>https://www.ugureskici.com/notlarim-makalelerim/murphy-kanunlari")
	color.Unset()

	res, err := http.Get("https://www.ugureskici.com/notlarim-makalelerim/murphy-kanunlari")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)

	doc.Find("article#post-2492 div.entry__content ul li").Each(func(i int, s *goquery.Selection) {
		kanun := s.Text()
		if i%4 == 0{
			color.Blue("%d->%s", i+1, kanun)
		}else{
			color.Cyan("%d->%s", i+1, kanun)
		}
	})
	color.Unset()
}