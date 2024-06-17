package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	PROXY     = "http://192.168.50.217:8888"
	TestPROXY = "https://www.baidu.com"
)

type Client struct {
	Proxy      string
	HttpClient http.Client
}

// newClient 创建 client，根据 hasProxy 判断是否设置代理
func (c *Client) newClient() {
	fmt.Println("[init client]")
	t := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	httpClient := http.Client{
		Timeout:   30 * time.Second,
		Transport: t,
	}
	if c.hasProxy() {
		fmt.Println("[settings proxy]", c.Proxy)
		proxyUrl, _ := url.Parse(c.Proxy)
		t.Proxy = http.ProxyURL(proxyUrl)
	}
	c.HttpClient = httpClient
}

// hasProxy 判断是否设置代理
func (c *Client) hasProxy() bool {
	return c.Proxy != ""
}

// TestProxy 如果设置了代理的话，测试代理是否可以连接
func (c *Client) TestProxy() {
	if c.hasProxy() {
		_, err := c.HttpClient.Get(TestPROXY)
		if err != nil {
			panic(err)
		}
	}
}

// request 发送请求，封装起来主要是统一处理 headers
func (c *Client) request(url string, method string, body string, headers map[string]string) (resp *http.Response, err error) {
	var req *http.Request
	if method == "" {
		method = "GET"
	}
	method = strings.ToUpper(method)

	if method == "GET" {
		req, err = http.NewRequest("GET", url, nil)
	} else if method == "POST" {
		req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	} else {
		panic("不支持的 Method")
	}

	// 添加 headers
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err = c.HttpClient.Do(req)
	return resp, err
}

func (c *Client) Get(url string, headers map[string]string) (resp *http.Response, err error) {
	return c.request(url, "GET", "", headers)
}

func (c *Client) Post(url string, body string, headers map[string]string) (resp *http.Response, err error) {
	return c.request(url, "POST", body, headers)
}

func main() {
	header := map[string]string{
		"Connection":         "keep-alive",
		"sec-ch-ua":          `"Microsoft Edge";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`,
		"Accept":             "application/json, text/plain, */*",
		"sec-ch-ua-mobile":   "?0",
		"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
		"sec-ch-ua-platform": "Windows",
		"Sec-Fetch-Site":     "same-origin",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Dest":     "empty",
		"Referer":            "https://www.cyouzai.com/subao-h5/passport/index.html",
		"Accept-Encoding":    "gzip, deflate, br, zstd",
		"Accept-Language":    "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
	}
	c := &Client{Proxy: PROXY}
	c.newClient()
	c.TestProxy()
	r, err := c.Get("https://www.cyouzai.com/subao-message/wechat/js/getSignature?token=&url=https:%2F%2Fwww.cyouzai.com%2Fsubao-h5%2Fpassport%2Findex.html", header)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	r, err = c.Get("https://www.cyouzai.com/subao-proxy-supplier/passport/channel/baseInfo?token=&channelCode=BAOBIAO_CAR", header)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// 获取当前时间
	currentTime := time.Now()

	// 获取 Unix 时间戳（毫秒级）
	unixTimeMilli := currentTime.UnixNano() / int64(time.Millisecond)

	header["Accept"] = "*/*"
	header["Sec-Fetch-Site"] = "cross-site"
	header["Sec-Fetch-Mode"] = "no-cors"
	header["Sec-Fetch-Dest"] = "script"
	header["Referer"] = "https://www.cyouzai.com/"
	captchaImgUrl := fmt.Sprintf("https://vpc-af.zhongan.io/api/v1/captcha.jsonp?callback=fn_%d&d=H4sIADauXmYA_w3L60-CQAAAcO6OA-J8LBws3exTHxp-a63mo9U8LbOXwGnYoQQLNzXYYNKGbf3t9gf8ms9d16mVjTAK4P7tD922DZREsZtzZxTLwbxndueFl73m3ONJyuSWVxnGq9gP-foxSdOPXltfFHm8CLnrFSmTzk0NFjz1fqII7cNw0DJPn36_MrpFq3yV0cVF61IIds7n9zrbBlK0w_B6VKDBSxIJJaN-_HCEO75OdAkKgGAIRXuWAAljLBDK7hnlw0anTzXOfXdJKta4eXVWtesElwFQgYglmbkbGwCEqlNFef934nIDHWChMpoxy1FV6CdEAPpJo68Rg6lTTG4USxGlknZHJ4Y9wcD0aW18ALetI90XAQAA&dt=dt&label_lang=zh", unixTimeMilli)
	r, err = c.Get(captchaImgUrl, header)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	if r.StatusCode != http.StatusOK {
		fmt.Println(string(body))
	}

}
