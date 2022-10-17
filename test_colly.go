// package main

// import (
// 	"fmt"
// 	"log"
// 	"regexp"
// 	"strings"

// 	"github.com/gocolly/colly"
// )

// func main() {
// 	c := colly.NewCollector()
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})

// 	c.OnError(func(_ *colly.Response, err error) {
// 		log.Println("Something went wrong:", err)
// 	})

// 	c.OnResponse(func(r *colly.Response) {
// 		fmt.Println("Visited", r.Request.URL)
// 	})

// 	// 获取每一个文章
// 	c.OnHTML(".article-panel", func(e *colly.HTMLElement) {
// 		title := e.ChildText(".a-post .title") // 标题
// 		fmt.Printf("标题: %v\n", title)

// 		content := e.ChildText(".a-post .content p") // 摘要
// 		fmt.Printf("摘要: %v\n", content)

// 		e.ForEach(".a-meta .mr-3", func(i int, h *colly.HTMLElement) {
// 			if i == 0 {
// 				fmt.Printf("日期: %v\n", h.Text) // 日期
// 			}
// 		})

// 		coverStyle := e.ChildAttr("a.a-thumb div", "style") // 获取封面style
// 		r, _ := regexp.Compile("url(.*);")                  // 获取背景url
// 		cover := r.FindString(coverStyle)
// 		cover = strings.Replace(cover, "url(", "", 1)
// 		cover = strings.Replace(cover, ");", "", 1)
// 		fmt.Printf("cover: %v\n", cover)

// 		detail := e.ChildAttr("a.a-thumb", "href")
// 		fmt.Printf("详情: %v\n", detail)
// 		fmt.Println("")
// 	})

// 	// 访问下一页
// 	c.OnHTML(".paginations a.next[href]", func(e *colly.HTMLElement) {
// 		e.Request.Visit(e.Attr("href"))
// 	})

// 	c.OnScraped(func(r *colly.Response) {
// 		fmt.Println("Finished", r.Request.URL)
// 	})

// 	c.Visit("https://rmb.ee")
// }
