package model

import "github.com/astaxie/beego/orm"

type Series struct {
	Id              int    `orm:"column(id)"`
	SeriesId        string `orm:"column(seriesId)"`
	SeriesName      string `orm:"column(seriesName)"`
	SeriesLink      string `orm:"column(seriesLink)"`
	SeriesImg       string `orm:"column(seriesImg)"`
	SeriesImgLoc    string `orm:"column(seriesImgLoc)"`
	SeriesPrice     string `orm:"column(seriesPrice)"`
	Urlmd5          string `orm:"column(urlmd5)"`
	SeriesScore     string `orm:"column(seriesScore)"`
	ProducerName    string `orm:"column(producerName)"`
	SourceURL       string `orm:"column(sourceURL)"`
	SeriesBigImg    string `orm:"column(seriesBigImg)"`
	SeriesBigImgLoc string `orm:"column(seriesBigImgLoc)"`
	SeriesEngine    string `orm:"column(seriesEngine)"`
	SeriesGearbox   string `orm:"column(seriesGearbox)"`
	SeriesCarStruct string `orm:"column(seriesCarStruct)"`
	UpdateTime      string `orm:"column(updateTime)"`
}

func init() {
	orm.RegisterModel(new(Series))
}
