package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	ProfileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUreRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	//cityUreRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
)

func ParseCity(contents []byte) engine.ParserResult {
	matches := ProfileRe.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, "User:"+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseProfile,
		})
	}
	matches = cityUreRe.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:  string(m[1]),
			ParserFunc:ParseCity,
		})

	}
	return result
}
