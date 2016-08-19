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

type Result struct {
	Article []Post `xml:"post"`
}
type Post struct {
	Title       string `xml:"id,attr"`
	Description string `xml:"tags,attr"`
	PicURL      string `xml:"preview_url,attr"`
	URL         string `xml:"sample_url,attr"`
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
