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
	Tags  string
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
	var buf bytes.Buffer

	buf.WriteString("https://yande.re/post.xml")
	buf.WriteByte('?')

	v := url.Values{}
	//	if que.Tags != nil {
	//	}
	v.Set("tags", que.Tags)
	v.Set("page", strconv.Itoa(que.Page))
	v.Set("limit", strconv.Itoa(que.Limit))
	//	v.Encode()
	buf.WriteString(v.Encode())
	resp, err := http.Get(buf.String())
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
