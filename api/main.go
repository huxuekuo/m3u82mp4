package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// 119.29.226.140:13457 172.247.47.125
func main() {

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		other_kkk217 := "http://xdm530.com"

		// 第一次 URL 编码
		encodedStr := url.QueryEscape(other_kkk217)

		// 将编码后的字符串中的 '%' 替换为 '%25'
		encodedStr = strings.Replace(encodedStr, "%", "%25", -1)

		// 第二次 URL 编码
		doubleEncodedStr := url.QueryEscape(encodedStr)
		resp, err := http.Get(fmt.Sprintf("http://119.29.226.140:13457/ssszz.php?top=10&q=%s&other_kkk217=%s&dect=0", url.QueryEscape(key), doubleEncodedStr))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		bytedata, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		// var datamap map[string]any
		// err = json.Unmarshal(bytedata, &datamap)
		// if err != nil {
		// 	panic(err)
		// }
		w.Write(bytedata)
	})
	http.ListenAndServe(":8080", nil)
}
