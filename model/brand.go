package model

import "github.com/astaxie/beego/orm"

type Brand struct {
	Id          int    `orm:"column(id)"`
	BrandName   string `orm:"column(brandName)"`
	BrandLink   string `orm:"column(brandLink)"`
	Urlmd5      string `rom:coloum(urlmd5)`
	BrandImg    string `orm:"column(brandImg)"`
	BrandImgLoc string `orm:"column(brandImgLoc)"`
	BrandIndex  string `orm:"column(brandIndex)"`
	SourceURL   string `orm:"column(sourceURL)"`
	UpdateTime  string `orm:"column(updateTime)"`
}

func init() {
	orm.RegisterModel(new(Brand))
}
