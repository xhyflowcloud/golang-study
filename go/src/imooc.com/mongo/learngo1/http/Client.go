package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	newRequest, err := http.NewRequest(
		http.MethodGet,
		"http://www.imooc.com", nil)

	newRequest.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//resp, err := http.DefaultClient.Do(newRequest)
	client := http.Client{
		CheckRedirect: func(
			req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(newRequest)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	dumpResponse, e := httputil.DumpResponse(resp, true)
	if e != nil {
		panic(e)
	}

	fmt.Printf("%s\n", dumpResponse)

}
