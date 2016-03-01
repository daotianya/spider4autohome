package process

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"regexp"
	"spider/dao"
	"spider/model"
	"spider/utils"
	"strconv"
	"strings"
	"time"
)

func processAricleDetail(page *page.Page) {
	sourceURL := page.GetRequest().GetUrl()
	doc := page.GetHtmlParser()

	re := regexp.MustCompile("http://www.autohome.com.cn/(\\w+)/(\\d+)/(\\d+).html")
	groups := re.FindStringSubmatch(sourceURL)

	articleCategoryTag := groups[1]
	articleProvider := doc.Find(".article-info span").Eq(1).Find("a").Eq(0).Text()
	articleType := doc.Find(".article-info span").Eq(2).Text()
	articleType = strings.Split(articleType, "：")[1]
	articleTagsSlice := []string{}
	articleTags := ""
	urlmd5 := utils.MD5(sourceURL)

	//文章标签
	doc.Find(".article-tags .tags a").Each(func(i int, s *goquery.Selection) {
		articleTagsSlice = append(articleTagsSlice, s.Text())
	})

	articleTags = strings.Join(articleTagsSlice, "|")

	//文章内容

	//所有图片
	doc.Find("#articleContent img").Each(func(i int, s *goquery.Selection) {
		imgsrc, _ := s.Eq(0).Attr("src")
		fmt.Println("imgsrc: " + imgsrc)
	})

	articleContent, _ := doc.Find("#articleContent").Eq(0).Html()

	//总页数
	totalPages := 1

	pTxt := doc.Find(".page .page-item-info").Text()
	if b, _ := regexp.MatchString("共(\\d+)页", pTxt); b {
		reg := regexp.MustCompile("共(\\d+)页")
		pages := reg.FindStringSubmatch(pTxt)[1]
		totalPages, _ = strconv.Atoi(pages)
	}

	article := model.Article{
		ArticleTags:        articleTags,
		ArticleProvider:    articleProvider,
		ArticleType:        articleType,
		ArticleCategoryTag: articleCategoryTag,
		ArticleContent:     articleContent,
		TotoalPages:        totalPages,
		Urlmd5:             urlmd5,
		UpdateTime:         time.Now().Format("2006-01-02 15:04:05"),
	}

	//fmt.Println(article)
	dao.UpdateArticel(&article)

	if totalPages > 1 {
		for i := 2; i <= totalPages; i++ {
			nextSpiderUrl := fmt.Sprintf("http://www.autohome.com.cn/%s/%s/%s-%d.html", groups[1], groups[2], groups[3], i)
			//fmt.Println("nextSpiderUrl:" + nextSpiderUrl)
			page.AddTargetRequest(nextSpiderUrl, "html")
		}
	}
}
