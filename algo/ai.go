package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 定义一个结构体来存储json数据
type ApiData struct {
	Method      string `json:"method"`
	Group       string `json:"group"`
	Path        string `json:"path"`
	Adapter     string `json:"adapter"`
	Transfer    string `json:"transfer"`
	Timeout     int    `json:"timeout"`
	Middleware  string `json:"middleware"`
	MercuryName string `json:"mercury_name"`
}

// 定义一个正则表达式来匹配文件中的目标代码
var re = regexp.MustCompile(`r\.(\w+)\.(\w+)\("([^"]+)",\s*([^)]+)\)`)

func main() {
	input := "/Users/lilinchong/gitCode/be_aio_bff/internal/ginrouter/app"
	output := "/Users/lilinchong/gitCode/tool-algorithm/algo/add"

	// 读取输入目录下的所有子目录
	inputDirs, err := ioutil.ReadDir(input)
	if err != nil {
		fmt.Println("读取输入目录失败:", err)
		return
	}
	for _, inputDir := range inputDirs {
		if inputDir.IsDir() {
			// 读取子目录下的所有文件
			inputFiles, err := ioutil.ReadDir(filepath.Join(input, inputDir.Name()))
			if err != nil {
				fmt.Println("读取子目录失败:", err)
				continue
			}
			for _, inputFile := range inputFiles {
				if !inputFile.IsDir() {
					// 读取文件内容
					inputData, err := ioutil.ReadFile(filepath.Join(input, inputDir.Name(), inputFile.Name()))
					if err != nil {
						fmt.Println("读取文件失败:", err)
						continue
					}
					// 查找文件中的目标代码
					matches := re.FindAllStringSubmatch(string(inputData), -1)
					if len(matches) > 0 {
						// 创建一个切片来存储json数据
						var apiDatas []ApiData
						for _, match := range matches {
							// 将匹配到的代码转换为json数据并添加到切片中
							apiData := ApiData{
								Method:   match[1],
								Group:    match[2],
								Path:     match[3],
								Adapter:  strings.Split(match[4], ".")[0],
								Transfer: strings.Split(match[4], ".")[1],
								Timeout:  1,
							}
							apiDatas = append(apiDatas, apiData)
						}
						// 将切片转换为json字符串
						outputData, err := json.MarshalIndent(apiDatas, "", "  ")
						if err != nil {
							fmt.Println("转换json失败:", err)
							continue
						}
						// 创建输出目录
						err = os.MkdirAll(output, 0755)
						if err != nil {
							fmt.Println("创建输出目录失败:", err)
							continue
						}
						// 创建输出文件并写入json字符串
						err = ioutil.WriteFile(filepath.Join(output, inputDir.Name()+".json"), outputData, 0644)
						if err != nil {
							fmt.Println("写入输出文件失败:", err)
							continue
						}
					}
				}
			}
		}
	}
}
