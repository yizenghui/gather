package yande

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// 查询器结构体
type Querier struct {
	Tag   string
	Page  int
	Limit int
}

// 文章集
type Result struct {
	Article []Post `xml:"post"`
}

// 文章
type Post struct {
	Title       string `xml:"id,attr"`
	Description string `xml:"tags,attr"`
	PicURL      string `xml:"preview_url,attr"`
	URL         string `xml:"sample_url,attr"`
}

// 标签集
type Tags struct {
	Tags []Tag `xml:"tag"`
}

// 标签
type Tag struct {
	Id    string `xml:"id,attr"`
	Name  string `xml:"name,attr"`
	Count string `xml:"count,attr"`
}

// 获取目标数据
func GetPost(que Querier) Result {

	resp, err := http.Get(GetUrl(que))
	if err != nil {

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	var result Result
	err = xml.Unmarshal(body, &result)
	if err != nil {

	}

	return result
}

// 获取目标分类标签
func GetTags() Tags {
	// 获取 100 个热门标签
	resp, err := http.Get("https://yande.re/tag.xml?order=count&limit=100")
	if err != nil {

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	var tags Tags
	err = xml.Unmarshal(body, &tags)
	if err != nil {

	}

	return tags
}

// 获取资源地址
func GetUrl(que Querier) string {
	var buf bytes.Buffer

	buf.WriteString("https://yande.re/post.xml")
	buf.WriteByte('?')

	v := url.Values{}
	//	if que.Tag != nil {
	//	}
	v.Set("tags", que.Tag)
	v.Set("page", strconv.Itoa(que.Page))
	v.Set("limit", strconv.Itoa(que.Limit))
	//	v.Encode()
	buf.WriteString(v.Encode())
	return buf.String()
}
