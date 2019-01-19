package parser

import (
	"imooc.com/mongo/learngo1/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	/*resp, err := http.Get("http://album.zhenai.com/u/1426975040")
	if err != nil {
		panic(err)
	}
	contents, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", contents)*/
	contents, err := ioutil.ReadFile(
		"profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "不想")

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1 "+
			"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expectedProfile := model.Profile{
		Name:       "不想",
		Gender:     "女",
		Age:        37,
		Height:     160,
		Weight:     55,
		Income:     "3-5千",
		Marriage:   "离异",
		Education:  "大学本科",
		Occupation: "阿坝马尔康",
		Hokou:      "四川阿坝",
		Xinzuo:     "天蝎座",
		House:      "已购房",
		Car:        "未买车",
	}

	if profile != expectedProfile {
		t.Errorf("expected %v; but was %v",
			expectedProfile, profile)
	}
}
