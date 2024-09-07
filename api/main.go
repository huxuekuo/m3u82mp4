package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
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
		w.Write(bytedata)

	})
	http.HandleFunc("/getInfo", func(w http.ResponseWriter, r *http.Request) {
		urls := r.FormValue("url")
		host := "http://xdm530.com"
		newHostUrl := host + urls
		// 禁用chrome headless
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", true),
		)
		allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
		defer cancel()
		ctx, cancel := chromedp.NewContext(
			allocCtx,
		)
		defer cancel()
		ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
		var onClick []*cdp.Node
		var aNodes []*cdp.Node
		err := chromedp.Run(ctx, chromedp.Navigate(newHostUrl),
			chromedp.WaitVisible("#tab12"),
			chromedp.Nodes("body > div.wrap > div.taba-down.mb.clearfix > div.pfrom.tab1.clearfix > #play_list_34", &onClick),
			chromedp.Nodes(":is(a)", &aNodes, chromedp.ByQueryAll))

		if err != nil {
			log.Fatal(err)
		}
		line := make(map[string]string, len(onClick[0].Children))
		for _, v := range onClick[0].Children {
			onClick, exits := v.Attribute("onclick")
			if exits {
				line[v.Children[0].NodeValue] = strings.Split(onClick, ",")[2]
			}
		}
		fmt.Printf("%+v\n", line)
		episodes := make(map[string]string, 0)
		for _, node := range aNodes {
			href, b := node.Attribute("href")
			if b && strings.Contains(href, "html") {
				if strings.Contains(href, "kb") {
					fmt.Println("debug")
				}
				queryParse, err := url.Parse(href)
				if err != nil {
					continue
				}
				qp := queryParse.Query().Get("qp")
				if qp == "" {
					episodes["default"] = href
				} else {
					episodes[qp] = href
				}
			}
		}

	})
	http.ListenAndServe(":8080", nil)
}
