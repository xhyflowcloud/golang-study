package parser

import (
	"imooc.com/mongo/learngo1/crawler/engine"
	"imooc.com/mongo/learngo1/crawler/model"
	"regexp"
	"strconv"
)

/*
    Name string
	Gender string
	Age int
	Height int
	Weight int
	Income string
	Marriage string
	Education string
	Occupation string
	Hokou string
	Xinzuo string
	House string
	Car string
*/
var genderRe = regexp.MustCompile(`([他|她])的动态`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([0-9]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([0-9]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([0-9]+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入.([^<]+)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([离异|未婚|丧偶]+)</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([高中及一下|中专|大专|大学本科|硕士|博士]+)</div>`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>工作地.([^<]+)</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>籍贯.([^<]+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+座)[^<]+</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([^<]*房[^<]*)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([^<]*车[^<]*)</div>`)

func ParseProfile(
	contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Name = name
	gender := extractString(
		contents, genderRe)
	if gender != "" {
		if gender == "她" {
			profile.Gender = "女"
		} else {
			profile.Gender = "男"
		}
	}

	age, err := strconv.Atoi(
		extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(
		extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(
		extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(
		contents, incomeRe)

	profile.Marriage = extractString(
		contents, marriageRe)

	profile.Education = extractString(
		contents, educationRe)

	profile.Occupation = extractString(
		contents, occupationRe)

	profile.Hokou = extractString(
		contents, hokouRe)

	profile.Xinzuo = extractString(
		contents, xinzuoRe)

	profile.House = extractString(
		contents, houseRe)

	profile.Car = extractString(
		contents, carRe)

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(
	contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
