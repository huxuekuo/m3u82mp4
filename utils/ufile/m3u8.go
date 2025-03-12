package ufile

import (
	"bufio"
	"fmt"
	"io"
	"m3u82mp4/library"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

const WorkDir = "/Users/huxuekuo/Downloads/"

type M3U8 struct {
	sourcePath  string
	targetPath  string
	callBack    func(node, total int)
	lock        *sync.Mutex
	nodeProcess int
}

func NewM3U8(sourcePath, targetPath string, callBack func(node, total int)) *M3U8 {
	return &M3U8{sourcePath: sourcePath, targetPath: targetPath, callBack: callBack, lock: &sync.Mutex{}}
}

func (m *M3U8) SetSourcePath(sourcePath string) {
	m.sourcePath = sourcePath
}

func (m M3U8) CheckMixed() string {
	resp, err := http.Get(m.sourcePath)
	if err != nil {
		library.Logger.Warn("请求资源错误", zap.Error(err))
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		library.Logger.Warn("请求资源状态码错误", zap.Int("status", resp.StatusCode))
		return ""
	}
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "mixed.m3u8") {
			return strings.ReplaceAll(m.sourcePath, filepath.Base(m.sourcePath), "") + text
		}
	}
	return ""
}

func (m M3U8) dowloadFile() (string, string) {
	resp, err := http.Get(m.sourcePath)
	if err != nil {
		library.Logger.Warn("请求资源错误", zap.Error(err))
		return "", ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		library.Logger.Warn("请求资源状态码错误", zap.Int("status", resp.StatusCode))
		return "", ""
	}
	// 创建单文件操作目录
	optDir := strconv.FormatInt(time.Now().UnixNano(), 10)
	optDir = WorkDir + optDir
	err = os.Mkdir(optDir, 0755)
	if err != nil {
		library.Logger.Warn("创建工作目录失败", zap.Error(err))
		return "", ""
	}

	tsPath := optDir + "/" + "ts.txt"
	file, err := os.Create(tsPath)
	scanner := bufio.NewScanner(resp.Body)
	tsUrls := []string{}
	// 匹配ts地址
	for scanner.Scan() {
		text := scanner.Text()
		b, line := matchMode(m.sourcePath, text, file)
		if !b {
			return "", ""
		}
		if line != "" {
			tsUrls = append(tsUrls, line)
		}
	}
	if len(tsUrls) <= 0 {
		library.Logger.Error("无ts数据")
		// 没有数据退出
		return "", ""
	}
	var gogroup sync.WaitGroup
	sem := make(chan struct{}, 50)
	for _, url := range tsUrls {
		gogroup.Add(1)
		go func(urls string) {
			defer func() {
				m.lock.Lock()
				m.nodeProcess++
				m.callBack(m.nodeProcess, len(tsUrls))
				m.lock.Unlock()
				gogroup.Done()

			}()
			sem <- struct{}{}
			defer func() { <-sem }() // 防止无法释放问题
			resp, err := http.Get(urls)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			file, err := os.Create(optDir + "/" + filepath.Base(urls))
			_, err = io.Copy(file, resp.Body)
		}(url)
	}
	gogroup.Wait()
	return tsPath, optDir
}

func (m M3U8) ToMP4() (bool, string) {
	tsPath, optDir := m.dowloadFile()
	if tsPath == "" || optDir == "" {
		return false, ""
	}
	if m.targetPath == "" {
		fileName := strconv.FormatInt(time.Now().UnixNano(), 10)
		m.targetPath = optDir + "/" + fileName + ".mp4"
	}
	cmd := exec.Command("/Users/huxuekuo/Downloads/ffmpeg", "-f", "concat", "-safe", "0", "-i", tsPath, "-c", "copy", m.targetPath)
	err := cmd.Run()
	if err != nil {
		library.Logger.Warn("合并文件失败", zap.Error(err))
		return false, ""
	}
	library.Logger.Info("合并成功", zap.String("targetPath", m.targetPath))
	return true, m.targetPath
}

func HttpMatchMode(line string) (bool, string) {
	if strings.Contains(line, "https://") || strings.Contains(line, "http://") {
		return true, line
	}
	return false, ""
}

func TsMatchMode(url, line string) (bool, string) {
	if strings.Contains(line, ".ts") {
		line = strings.ReplaceAll(url, filepath.Base(url), "") + line
		return true, line
	}
	return false, ""
}

func matchMode(url, line string, file *os.File) (bool, string) {
	txt := ""
	if b, l := HttpMatchMode(line); b {
		txt = l
	}
	if b, l := TsMatchMode(url, line); b {
		txt = l
	}
	if txt == "" {
		return true, ""
	}
	_, err := file.WriteString("file '" + filepath.Base(line) + "'\n")
	if err != nil {
		library.Logger.Warn("ts写入文件失败", zap.String("fileName", filepath.Base(line)), zap.Error(err))
		return false, ""
	}
	return true, txt
}
