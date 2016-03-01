package process

import (
	"github.com/hu17889/go_spider/core/common/page"
	"regexp"
	"spider/dao"
	"spider/download"
	"spider/model"
	"spider/utils"
	"time"
)

func processAricle(page *page.Page) {
	sourceURL := page.GetRequest().GetUrl()
	re := regexp.MustCompile("http://www.autohome.com.cn/(\\d+)/0/0\\-0\\-[0-9]{1}\\-0/")
	seriesId := re.FindStringSubmatch(sourceURL)[1]

	seriesUrlMd5 := utils.MD5("http://www.autohome.com.cn/" + seriesId + "/#levelsource=000000000_0&pvareaid=102538")

	doc := page.GetHtmlParser()
	newList := doc.Find(".cont-info ul li")

	for i, l := 0, newList.Size(); i < l; i++ {
		li := newList.Eq(i)

		articleLogo, _ := li.Find(".newpic a img").Eq(0).Attr("src")
		articleLink, _ := li.Find("h3 a").Eq(0).Attr("href")
		articleTitle := li.Find("h3 a").Eq(0).Text()
		articleAuthor := li.Find(".name-tx span").Eq(0).Text()
		articlePubDate := li.Find(".name-tx span").Eq(1).Text()
		articleViewCount := li.Find(".name-tx span").Eq(2).Text()
		articleCommentCount := li.Find(".name-tx span").Eq(3).Text()

		//下载图标
		imageFileName := download.SaveImage(articleLogo, ImageRootDir+"articlelogo/")
		articleLogoLoc := ""

		if imageFileName != "" {
			articleLogoLoc = "articlelogo/" + imageFileName
		}

		aritlce := model.Article{
			SeriesId:            seriesId,
			SeriesUrlMd5:        seriesUrlMd5,
			ArticleLink:         articleLink,
			ArticleTitle:        articleTitle,
			ArticleLogo:         articleLogo,
			ArticleLogoLoc:      articleLogoLoc,
			ArticleAuthor:       articleAuthor,
			ArticlePubDate:      articlePubDate,
			ArticleViewCount:    articleViewCount,
			ArticleCommentCount: articleCommentCount,
			Urlmd5:              utils.MD5(articleLink),
			TotoalPages:         1,
			SourceURL:           sourceURL,
			UpdateTime:          time.Now().Format("2006-01-02 15:04:05"),
		}
		//fmt.Println(aritlce)

		//保存到数据库
		dao.SaveArticle(&aritlce)

		//详情页 http://www.autohome.com.cn/news/201512/883375.html
		page.AddTargetRequest(articleLink, "html")
	}

	//下一页链接
	nextpageUrl, _ := doc.Find(".page a.current").Next().Eq(0).Attr("href")

	if b, _ := regexp.MatchString("http://www.autohome.com.cn/\\d+/0/0\\-0\\-[0-9]{1}\\-0/", nextpageUrl); b {
		page.AddTargetRequest(nextpageUrl, "html")
	}
}
