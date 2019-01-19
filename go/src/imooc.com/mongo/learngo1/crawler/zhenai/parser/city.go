package parser

import (
	"imooc.com/mongo/learngo1/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(
	contents []byte) engine.ParserResult {
	re, err := regexp.Compile(cityRe)
	if err != nil {
		panic(err)
	}
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(
			result.Items, "User "+string(m[2]))
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParserFunc: func(
					contents []byte) engine.ParserResult {
					return ParseProfile(
						contents, string(m[2]))
				},
			})
	}
	return result
}
