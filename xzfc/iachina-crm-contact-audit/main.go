// package
package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

const (
	host = "http://member.iachina.cn" // 线上
	//host = "http://127.0.0.1:8000" // 本地
)

var header = map[string]string{
	"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	"Cookie": `__SESSION__="2|1:0|10:1722415704|11:__SESSION__|48:YTU4MmEwYjAtNGYxOS0xMWVmLTkwNWMtZmExNjNlNTg0NjE1|e5c3b8fa01ece66699a886024121700b83e1b5567c81f1b3ba2592f6422da94c"; next="/ad"`,
}
var logger = log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)

// Client 中保协网站客户端
type Client struct {
	http.Client
	Host string
}

// get 发送 get 请求
func (c *Client) get(path string, queryParam string, header map[string]string) (*http.Response, error) {
	var url string
	if queryParam == "" {
		url = c.Host + path
	} else {
		url = c.Host + path + "?" + queryParam
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := c.Do(req)
	logger.Printf("[request %s] status: %s err: %s", path, resp.Status, err)
	return resp, err
}

// post 发送 post 请求
func (c *Client) post(path string, queryParam string, header map[string]string) (*http.Response, error) {
	var url string
	if queryParam == "" {
		url = c.Host + path
	} else {
		url = c.Host + path + "?" + queryParam
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := c.Do(req)
	logger.Printf("[request %s] status: %s err: %s", path, resp.Status, err)
	return resp, err
}

// testConnect 验证网站是否能正常访问
func (c *Client) testConnect() bool {
	resp, err := c.get("/ad/login", "", map[string]string{})
	if err != nil {
		logger.Fatalln("[测试网站能否正常访问失败]", err)
	}
	if resp.StatusCode != 200 {
		return false
	}
	return true
}

// getAwaitReviewContactList 获取未审核的通讯录列表
func (c *Client) getAwaitReviewContactList() string {
	resp, err := c.get("/ad/contact/list", "company_name=&company_type_1=&iac_title_id=&title_type=&name=&mobile=&email=&status=3&page=1&size=10&search_from=&check_all_pages=false", header)
	if err != nil {
		logger.Fatalln("[获取待审核通讯录列表失败]", err)
	}
	if resp.StatusCode != 200 {
		logger.Fatalln("[获取待审核通讯录列表失败]", resp.Status)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Fatalln("[读取待审核通讯录列表 Body 失败]", err)
	}
	return string(body)
}

// getContactDetail 获取通讯录详情信息
func (c *Client) getContactDetail(path string) string {
	resp, err := c.get(path, "", header)
	if err != nil {
		logger.Fatalln("[获取通讯录详情异常]", err)
	}
	if resp.StatusCode != 200 {
		logger.Fatalln("[获取通讯录详情失败]", resp.Status)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Fatalln("[读取通讯录详情 Body 异常]", err)
	}
	return string(body)
}

func (c *Client) passContact(path string) {
	resp, err := c.post(path, "", header)
	if err != nil {
		logger.Fatalln("[审核通讯录异常]", err)
	}
	if resp.StatusCode != 200 {
		logger.Fatalln("[审核通讯录失败]", resp.Status)
	}
}

func main() {
	num := 1
	c := &Client{Host: host}
	isConnect := c.testConnect()
	if !isConnect {
		logger.Fatalln("[网站无法访问，程序退出]")
	}
	for {
		arcBody := c.getAwaitReviewContactList()
		matches := extractContactDetailUrlForArcBody(arcBody)
		if len(matches) == 0 {
			os.Exit(0)
		}
		for _, match := range matches {
			path := match[1]
			detailBody := c.getContactDetail(path)
			passPath := extractContactReviewUrlForDetailBody(detailBody)
			logger.Println(passPath, "      ", num)
			num += 1
			c.passContact(passPath)
		}
	}
}

// extractContactDetailUrlForArcBody 从待审批通讯录列表获取通讯录详情的 url
func extractContactDetailUrlForArcBody(body string) [][]string {
	regex := regexp.MustCompile(`<a class="btn_detail btn btn-warning btn-xs pull-left" href="([^"]+)"  style="background:#e27626;border:1px solid #e27626;">审核</a>`)
	matches := regex.FindAllStringSubmatch(body, -1)
	return matches
}

// extractContactReviewUrlForDetailBody 从通讯录详情页面提取审批通过的 url
func extractContactReviewUrlForDetailBody(body string) string {
	regex := regexp.MustCompile(`<button class="btn-succ btn-success btn" data-url="([^"]+)">审批通过</button>`)
	matches := regex.FindAllStringSubmatch(body, -1)
	return matches[0][1]
}
