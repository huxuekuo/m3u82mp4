package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// 禁用chrome headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://xdm530.com/acg/6907/14.html`),
		chromedp.Sleep(3*time.Second),
		chromedp.WaitVisible("#playiframe"),
		chromedp.OuterHTML("#playiframe", &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	urls := strings.Split(example, "$$$")
	url := strings.Replace(strings.Split(urls[1], "|||")[1], "497", "", 1)
	host := strings.Replace(url, "index.m3u8", "", 1)
	download := downloadFile(url, "index.m3u8")
	if download {
		m3u8file, _ := readFile("index.m3u8")
		fmt.Println(host + m3u8file)
		mixedDow := downloadFile(host+m3u8file, "mixed.m3u8")
		if mixedDow {
			//https://vip.ffzy-play9.com/20230104/32931_64735b7f/2000k/hls/d079907d8d0c211c872be23029abfb01.ts
			_, ts := readFile("mixed.m3u8")
			fmt.Println(ts)
		}
	}

}

func downloadFile(url, name string) bool {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 创建一个文件用于保存
	out, err := os.Create(name)
	if err != nil {
		// panic(err)
		return false
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		// panic(err)
		return false
	}
	return true
}

func readFile(path string) (string, []string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", nil
	}
	defer file.Close() // 确保文件在函数结束时关闭

	// 创建 Scanner 来逐行读取文件
	tsArray := make([]string, 0)
	scanner := bufio.NewScanner(file)
	// 逐行读取，直到到达第三行
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "mixed.m3u8") {
			return line, nil
		}
		if strings.Contains(line, ".ts") {
			tsArray = append(tsArray, line)
		}

	}
	return "", tsArray

}
