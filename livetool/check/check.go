package check

import (
	"encoding/csv"
	"fmt"
	"livtool/settings"
	"log"
	"os"
)

func StartCheckIpAndPorts() {
	readCSV()
}

func readCSV() {
	// 打开CSV文件（替换为你的文件路径）
	file, err := os.Open(settings.CsvFileName)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 创建CSV读取器
	csvReader := csv.NewReader(file)

	// 读取所有记录（包含标题行）
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("读取CSV失败: %v", err)
	}

	// 输出结果
	fmt.Println("CSV内容:")
	for _, record := range records {
		fmt.Println(record)
	}

	// 如果需要跳过标题行，从第二行开始处理
	if len(records) > 0 {
		fmt.Println("\n跳过标题后的数据:")
		for _, record := range records[1:] {
			fmt.Println(record)
		}
	}
}

func check() {

}
