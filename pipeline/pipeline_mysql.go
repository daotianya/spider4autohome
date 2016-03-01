package pipeline

import (
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page_items"
)

type PipelineMySql struct {
}

func NewPipelineMySql() *PipelineMySql {
	return &PipelineMySql{}
}

func (this *PipelineMySql) Process(items *page_items.PageItems, t com_interfaces.Task) {
	println("----------------------------------------------------------------------------------------------")
	println("Crawled url :\t" + items.GetRequest().GetUrl() + "\n")
	println("Crawled result : ")
	for key, value := range items.GetAll() {
		println(key + "\t:\t" + value)
	}


}
