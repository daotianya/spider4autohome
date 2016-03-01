package process

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"spider/dao"
	"spider/download"
	"spider/model"
	"spider/utils"
	"strings"
	"time"
)

//处理品牌列表页面
func processBrandPage(page *page.Page) {
	sourceURL := page.GetRequest().GetUrl()
	idxL := strings.Split(sourceURL, "/")
	last := idxL[len(idxL)-1]
	//品牌开头字母
	brandIndex := string([]byte(last)[0])

	doc := page.GetHtmlParser()
	doc.Find("dl").Each(func(i int, dl *goquery.Selection) {

		brandImg, _ := dl.Find("dt a img").Attr("src")
		brandName := dl.Find("dt div a").Text()
		brandLink, _ := dl.Find("dt div a").Attr("href")

		//下载图标
		imageFileName := download.SaveImage(brandImg, ImageRootDir+"brandlogo/")
		brandImgLoc := ""

		if imageFileName != "" {
			brandImgLoc = "brandlogo/" + imageFileName
		}

		brand := model.Brand{
			BrandName:   brandName,
			BrandLink:   brandLink,
			Urlmd5:      utils.MD5(brandLink),
			BrandImg:    brandImg,
			BrandImgLoc: brandImgLoc,
			BrandIndex:  brandIndex,
			SourceURL:   sourceURL,
			UpdateTime:  time.Now().Format("2006-01-02 15:04:05"),
		}

		//保存到数据库
		dao.SaveBrand(&brand)

		//fmt.Printf("品牌: %s, %s, %s, %s\n", brandName, brandLink, brandImg, brandIndex)

		//每个车品牌下又有不同的类型，比如宝马下有华晨宝，进口宝马，宝马M
		dl.Find("dd .h3-tit").Each(func(i int, tit *goquery.Selection) {
			producerName := tit.Text() //生产商

			p := model.Producer{
				ProducerName: producerName,
				UpdateTime:   time.Now().Format("2006-01-02 15:04:05"),
			}

			//保存到数据库
			dao.SaveProducer(&p)

			//每个类下面又有不同的车系
			tit.NextFiltered("ul").Find("li").Not(".dashline").Each(func(i int, li *goquery.Selection) {
				seriesName := li.Find("h4 a").Text()
				seriesLink, _ := li.Find("h4 a").Eq(0).Attr("href")
				seriesImg, _ := li.Find("div").Eq(0).Find("img").Eq(0).Attr("src")
				seriesPrice := li.Find("div").Eq(1).Find("a.red").Eq(0).Text()
				seriesId := strings.Split(seriesLink, "/")[3]
				//下载图标
				imageFileName := download.SaveImage(seriesImg, ImageRootDir+"serieslogo/")
				seriesImgLoc := ""

				if imageFileName != "" {
					seriesImgLoc = "serieslogo/" + imageFileName
				}

				series := model.Series{
					SeriesId:     seriesId,
					SeriesName:   seriesName,
					SeriesLink:   seriesLink,
					SeriesImg:    seriesImg,
					SeriesImgLoc: seriesImgLoc,
					SeriesPrice:  seriesPrice,
					Urlmd5:       utils.MD5(seriesLink),
					ProducerName: producerName,
					SourceURL:    sourceURL,
					UpdateTime:   time.Now().Format("2006-01-02 15:04:05"),
				}

				//保存到数据库
				dao.SaveSeries(&series)

				page.AddTargetRequest(seriesLink, "html")

				//fmt.Printf("分类: %s, %s, %s, %s, %s\n", producerName, seriesName, seriesLink, seriesImg, seriesPrice)
			})
		})
	})
}
