package parser

import (
	"imooc.com/mongo/learngo1/crawler/model"
	"net/http"
	"net/http/httputil"
	"testing"
)

/*
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9
Cache-Control: max-age=0
Connection: keep-alive
Cookie: sid=48a31e37-676b-49d8-8c37-ae40cbfb9b3b; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1547900298; isSignOut=%5E%7ElastLoginActionTime%3D1547911808799%5E%7E; mid=%5E%7Emid%3D110776738%5E%7E; loginactiontime=%5E%7Eloginactiontime%3D1547911808799%5E%7E; live800=%5E%7EinfoValue%3DuserId%253D110776738%2526name%253D110776738%2526memo%253D%5E%7EisOfflineCity%3Dtrue%5E%7E; preLG_110776738=2019-01-19+23%3A30%3A09; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1547917624
Host: album.zhenai.com
Referer: http://www.zhenai.com/zhenghun/suzhou
Upgrade-Insecure-Requests: 1

*/
func TestParseProfile(t *testing.T) {
	newRequest, err := http.NewRequest(
		http.MethodGet,
		"http://album.zhenai.com/u/1426975040", nil)
	if err != nil {
		panic(err)
	}

	//newRequest.Header.Add("Accept",
	//	"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	//newRequest.Header.Add("Accept-Encoding", "gzip, deflate")
	//newRequest.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	//newRequest.Header.Add("Cache-Control", "max-age=0")
	//newRequest.Header.Add("Cookie", "sid=48a31e37-676b-49d8-8c37-ae40cbfb9b3b; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1547900298; isSignOut=%5E%7ElastLoginActionTime%3D1547911808799%5E%7E; mid=%5E%7Emid%3D110776738%5E%7E; loginactiontime=%5E%7Eloginactiontime%3D1547911808799%5E%7E; live800=%5E%7EinfoValue%3DuserId%253D110776738%2526name%253D110776738%2526memo%253D%5E%7EisOfflineCity%3Dtrue%5E%7E; preLG_110776738=2019-01-19+23%3A30%3A09; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1547918740")
	//newRequest.Header.Add("host", "album.zhenai.com")
	//newRequest.Header.Add("Referer", "http://www.zhenai.com/zhenghun/aba")
	//newRequest.Header.Add("Upgrade-Insecure-Requests", "1")
	//newRequest.Header.Set("UserAgent",
	//"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	newRequest.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	client := http.Client{}

	resp, err := client.Do(newRequest)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	contents, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", contents)
	/*contents, err = ioutil.ReadFile(
		"profile_test_data.html")
	if err != nil {
		panic(err)
	}
	*/
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
