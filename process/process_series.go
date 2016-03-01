package process

import (
	"github.com/hu17889/go_spider/core/common/page"
	"spider/dao"
	"spider/download"
	"spider/model"
	"spider/utils"
	"strings"
	"time"
)

//车系页处理
func processSeries(page *page.Page) {
	doc := page.GetHtmlParser()
	seriesInfoBox := doc.Find(".autoseries-info")

	seriesScore := doc.Find(".koubei-score .font-score").Eq(0).Text() //评分
	seriesBigImg, _ := doc.Find(".autoseries-pic-img1 a img").Eq(0).Attr("src")
	seriesEngine := seriesInfoBox.Find("dl dd").Eq(1).Find("a").Eq(0).Text()    //发动机
	seriesGearbox := seriesInfoBox.Find("dl dd").Eq(2).Find("a").Eq(0).Text()   //变速箱
	seriesCarStruct := seriesInfoBox.Find("dl dd").Eq(2).Find("a").Eq(1).Text() //车身结构
	seriesLink := page.GetRequest().GetUrl()
	seriesId := strings.Split(seriesLink, "/")[3]

	//fmt.Printf("seriesScore: %s\n", seriesScore)
	//fmt.Printf("seriesBigImg: %s\n", seriesBigImg)
	//fmt.Printf("seriesEngine: %s\n", seriesEngine)
	//fmt.Printf("seriesGearbox: %s\n", seriesGearbox)
	//fmt.Printf("seriesCarStruct: %s\n", seriesCarStruct)

	//下载图标
	imageFileName := download.SaveImage(seriesBigImg, ImageRootDir+"serieslogo/")
	seriesBigImgLoc := ""

	if imageFileName != "" {
		seriesBigImgLoc = "serieslogo/" + imageFileName
	}

	series := model.Series{
		SeriesScore:     seriesScore,
		SeriesBigImg:    seriesBigImg,
		SeriesBigImgLoc: seriesBigImgLoc,
		SeriesEngine:    seriesEngine,
		SeriesGearbox:   seriesGearbox,
		SeriesCarStruct: seriesCarStruct,
		Urlmd5:          utils.MD5(seriesLink),
		UpdateTime:      time.Now().Format("2006-01-02 15:04:05"),
	}

	//fmt.Println(series)

	//保存到数据库
	dao.UpdateSeries(&series)

	page.AddTargetRequest("http://www.autohome.com.cn/"+seriesId+"/0/0-0-1-0/", "html")
}
