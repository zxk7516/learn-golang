package models

import (
	"zxk/go-fiber-demo/common"
)

type Page struct {
	ID            int    `gorm:"column:id;primary_key" json:"id"`
	URL           string `gorm:"column:url" json:"url"`
	Title         string `gorm:"column:title" json:"title"`
	ContentDigest string `gorm:"column:content_digest" json:"content_digest"`
	Content       string `gorm:"column:content" json:"content"`
	Keywords      string `gorm:"column:keywords" json:"keywords"`
	MetaContent   string `gorm:"column:metacontent" json:"metacontent"`
	NeedKey       bool   `gorm:"column:need_key" json:"need_key"`
	Password      string `gorm:"column:password" json:"password"`
	Tags          string `gorm:"column:tags" json:"tags"`
	Editor        string `gorm:"column:editor" json:"editor"`
	AllowVisit    bool   `gorm:"column:allow_visit" json:"allow_visit"`
	AllowComment  bool   `gorm:"column:allow_comment" json:"allow_comment"`
	IsOriginal    bool   `gorm:"column:is_original" json:"is_original"`
	NumLookup     int    `gorm:"column:num_lookup" json:"num_lookup"`
	HTML          string `gorm:"column:html" json:"html"`
	TimeStamps
}

func (p *Page) TableName() string {
	return "page"
}

func (p *Page) GetFullURL(schema, host string) string {
	return schema + "://" + host + "/" + p.URL
}

func GetPage(condition interface{}) (Page, error) {
	var page Page
	err := common.DB.Where(condition).First(&page).Error
	return page, err
}
func GetPagesByIds(ids []int) ([]Page, error) {
	var page []Page
	err := common.DB.Where("id in ?", ids).First(&page).Error
	return page, err
}

func GetPages(condition interface{}, p *common.Pagination) ([]Page, error) {
	var pages []Page
	err := common.DB.Where(condition).Order("id desc").Offset(p.GetOffset()).Limit(p.GetLimit()).Find(&pages).Error
	return pages, err
}

func SearchPage(tagname string, p *common.Pagination) ([]Page, error) {
	var pages []Page
	err := common.DB.Where("tags like ?", "%"+tagname+"%").Order("id desc").Offset(p.GetOffset()).Limit(p.GetLimit()).Find(&pages).Error
	return pages, err
}
