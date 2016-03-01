package process

import (
	"fmt"
	"github.com/hu17889/go_spider/core/common/page"
	"regexp"
	"spider/conf"
)

var ImageRootDir string = conf.IMGRoot

type AutohomeProcess struct {
	Name string
}

//构造方法
func New(name string) *AutohomeProcess {
	return &AutohomeProcess{Name: name}
}

func (p *AutohomeProcess) Process(page *page.Page) {

	sourceURL := page.GetRequest().GetUrl()

	if b, _ := regexp.MatchString(".+[A-Z]{1}_photo\\.html", sourceURL); b {
		processBrandPage(page)
	} else if b, _ := regexp.MatchString("http://www.autohome.com.cn/\\d+/#levelsource\\.*", sourceURL); b {
		processSeries(page)
	} else if b, _ := regexp.MatchString("http://www.autohome.com.cn/\\d+/0/0\\-0\\-[0-9]{1}\\-0/", sourceURL); b {
		processAricle(page)
	} else if b, _ := regexp.MatchString("http://www.autohome.com.cn/(\\w+)/(\\d+)/(\\d+).html", sourceURL); b {
		processAricleDetail(page)
	} else if b, _ := regexp.MatchString("http://www.autohome.com.cn/(\\w+)/(\\d+)/(\\d+)-(\\d+).html", sourceURL); b {
		processAriclePages(page)
	}

}

func (p *AutohomeProcess) Finish() {
	fmt.Printf("\n%s Finish!\n", p.Name)
}
