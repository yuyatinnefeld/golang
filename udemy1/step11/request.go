package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	resp, _ := http.Get("http://example.com")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll((resp.Body))
	fmt.Println(body)

	fmt.Println(string(body))

	fmt.Println("####### GET #########")
	base, _ := url.Parse("http://example.com")
	fmt.Println(base)
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	q := req.URL.Query()
	fmt.Println(q)
	q.Add("c", "3&%%%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resp1, _ := client.Do(req)
	body2, _ := ioutil.ReadAll(resp1.Body)
	fmt.Println(string(body2))

	fmt.Println("####### POST #########")
	endpoint = "http://example.com"
	req3, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
	resp3, _ := client.Do(req3)
	body3, _ := ioutil.ReadAll(resp3.Body)
	fmt.Println(string(body3))

}
