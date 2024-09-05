package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

var wg sync.WaitGroup

func main() {
	path := flag.String("PATH", "", "下载地址")
	url := flag.String("URL", "", "网址切换页面使用{index}替换")
	var startIndex int
	flag.IntVar(&startIndex, "start", 0, "起始集数")
	var index int
	flag.IntVar(&index, "index", 0, "最大集数")
	flag.Parse()
	f := 0
	for f != index+1 {
		if startIndex <= f {
			urlNew := strings.Replace(*url, "{index}", fmt.Sprint(f), 1)
			chromedpTest(fmt.Sprintf("%s/%s/", *path, fmt.Sprint(f)), urlNew, strconv.Itoa(f))
		}
		f++
	}
}

func chromedpTest(path, urlsss, index string) {
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
		chromedp.Navigate(urlsss),
		chromedp.Sleep(3*time.Second),
		chromedp.WaitVisible("#playiframe"),
		chromedp.OuterHTML("#playiframe", &example),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 匹配任意两个字母开头，后面跟着三个竖线和URL，最后是三个美元符号
	linkRegexp := regexp.MustCompile(`([A-Za-z]{2})\|\|\|(https://[^\|\$]+)`)
	// 在字符串中查找所有匹配的链接
	links := linkRegexp.FindAllStringSubmatch(example, -1)
	url := strings.Replace(links[0][2], "497", "", 1)
	host := strings.Replace(url, "index.m3u8", "", 1)
	download := downloadFile(url, "index.m3u8", path)
	if download {
		m3u8file, _ := readFile(path + "index.m3u8")
		fmt.Println(host + m3u8file)
		mixedDow := downloadFile(host+m3u8file, "mixed.m3u8", path)
		tsUrl := strings.Replace(host+m3u8file, "mixed.m3u8", "", 1)
		if mixedDow {
			_, ts := readFile(path + "mixed.m3u8")
			tsDownload := true
			sem := make(chan struct{}, 20)
			for i, v := range ts {
				wg.Add(1)
				sem <- struct{}{} // 进入并发段
				go func(fileName string, index int) {
					defer wg.Done()
					defer func() { <-sem }()
					tsDownload = downloadFile(tsUrl+fileName, fileName, path+"ts/")
					if !tsDownload {
						tsDownload = !tsDownload
					}
					fmt.Printf("%s-%v==", fmt.Sprint(index), tsDownload)
				}(v, i)
				if !tsDownload {
					break
				}
			}
			wg.Wait() // 等待所有goroutine完成
			if tsDownload {
				url := ""
				for _, v := range ts {
					url += "file '" + path + "ts/" + v + "' \n"
				}
				err := os.WriteFile(path+"ts.txt", []byte(url), 0644)
				if err != nil {
					log.Fatal(err)
				}
			}

			// 合并成视频
			cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", path+"ts.txt", "-c", "copy", path+index+".mp4")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error executing command:", err)
				return
			}

			// 打印输出
			fmt.Println(string(output))

		}
	}

}

func downloadFile(url, name, path string) bool {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 创建一个文件用于保存
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return false
	}
	out, err := os.Create(path + name)
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
	fmt.Println("ts文件数:" + fmt.Sprint(len(tsArray)))
	return "", tsArray

}
