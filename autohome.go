package main

import (
	"github.com/hu17889/go_spider/core/common/mlog"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"spider/conf"
	"spider/process"
)

func main() {

	spider := spider.NewSpider(process.New("AutohomeSpider"), "autohome")
	spider.SetSleepTime("rand", 1000, 3000)
	spider.OpenFileLog(conf.LOGRoot)

	mlog.LogInst().LogInfo("spider start")

	for i := 'A'; i <= 'Z'; i++ {
		v := string([]rune{i})
		spider.AddUrl("http://www.autohome.com.cn/grade/carhtml/"+v+"_photo.html", "html")
	}

	spider.AddUrl("http://www.autohome.com.cn/tuning/201412/853818.html", "html")
	spider.AddPipeline(pipeline.NewPipelineConsole()).SetThreadnum(16).Run()

	mlog.LogInst().LogInfo("spider end")

	spider.CloseFileLog()
}
