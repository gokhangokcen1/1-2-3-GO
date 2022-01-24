package main

import (
	"fmt"
  "github.com/PuerkitoBio/goquery"
  "net/http"
)

func main() {
	res, _ := http.Get("http://www.imdb.com/chart/top")
  if res.StatusCode != 200 {
    fmt.Println("Hata",res.StatusCode)
    return
  }
  doc, _ := goquery.NewDocumentFromReader(res.Body)

  doc.Find(".titleColumn").Each(func(i int, selection *goquery.Selection){
    title := selection.Find("a").Text()
    fmt.Println(i+1, title)
  })
}
