package settings

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	CurrentDir  string
	CsvFileName string
)

func InitSetting() {
	// 获取当前路径
	CurrentDir, _ = os.Getwd()

	// 获取csv 文件名称
	getCsvFileName()
}

func getCsvFileName() {
	// 收集所有CSV文件
	csvFiles := []string{}
	err := filepath.Walk(CurrentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".csv" {
			csvFiles = append(csvFiles, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("遍历目录失败: %v", err)
	}

	// 处理结果
	if len(csvFiles) == 0 {
		fmt.Println("当前目录没有CSV文件")
	} else {
		fmt.Println("找到的CSV文件:")
		for _, file := range csvFiles {
			fmt.Println(file)
			csvFileName = file
			break
		}
	}
}
