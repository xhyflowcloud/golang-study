package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"imooc.com/mongo/learngo1/crawler/engine"
	"imooc.com/mongo/learngo1/crawler/zhenai/parser"
	"io"
	"regexp"
)

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {

	re, err := regexp.Compile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	if err != nil {
		panic(err)
	}

	/*matchs := re.FindAll(contents, -1)

	for _, m := range matchs {
		fmt.Printf("%s\n", m)
	}*/

	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n",
			m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n",
		len(matches))
}

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

	/*resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",
			resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", all)
	printCityList(all)*/
}
