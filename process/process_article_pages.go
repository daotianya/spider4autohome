package process

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"regexp"
	"spider/dao"
	"spider/model"
	"spider/utils"
	"time"
)

func processAriclePages(page *page.Page) {
	sourceURL := page.GetRequest().GetUrl()
	doc := page.GetHtmlParser()

	re := regexp.MustCompile("http://www.autohome.com.cn/(\\w+)/(\\d+)/(\\d+)-(\\d+).html")
	groups := re.FindStringSubmatch(sourceURL)

	articleUrlMd5 := utils.MD5(fmt.Sprintf("http://www.autohome.com.cn/%s/%s/%s.html", groups[1], groups[2], groups[3]))
	pageNo := groups[4]
	urlmd5 := utils.MD5(sourceURL)
	pageContent, _ := doc.Find("#articleContent").Eq(0).Html()

	//所有图片
	doc.Find("#articleContent img").Each(func(i int, s *goquery.Selection) {
		imgsrc, _ := s.Eq(0).Attr("src")
		fmt.Println("imgsrc: " + imgsrc)
	})

	//fmt.Printf("articleUrlMd5: %s\n", articleUrlMd5)
	//fmt.Printf("pageNo: %s\n", pageNo)
	//fmt.Printf("urlmd5: %s\n", urlmd5)
	//fmt.Printf("pageContent: %s\n", pageContent)

	pageStruct := model.ArticlePage{
		ArticleUrlMd5: articleUrlMd5,
		PageNo:        pageNo,
		PageContent:   pageContent,
		PageLink:      sourceURL,
		Urlmd5:        urlmd5,
		UpdateTime:    time.Now().Format("2006-01-02 15:04:05"),
	}

	dao.SaveArticlePage(&pageStruct)
}
