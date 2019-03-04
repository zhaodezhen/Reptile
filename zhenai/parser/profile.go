package parser

import (
	"crawler/engine"
	"crawler/zhenai/model"
	"log"
	"src/github.com/PuerkitoBio/goquery"
	"strings"
)

//获取用户详情
func ParseProfile(contents []byte) engine.ParserResult {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		log.Fatalln(err)
	}
	var user [10]string
	document.Find(".purple").Each(func(i int, selection *goquery.Selection) {
		user[i] = selection.Text()
	})

	profile := model.Profile{}
	profile.Name = user[0]
	profile.Marriage = user[1]
	profile.Age = user[2]
	profile.Constellation = user[3]
	profile.Height = user[4]
	profile.Weight = user[5]
	profile.Workplace = user[6]
	profile.Income = user[7]
	profile.Occupation = user[8]
	profile.Education = user[9]
	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result
}

