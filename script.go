package main

import (
	"context"
	"encoding/csv"
	"errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"os"
	"strings"
)

// Conversion 转换函数
func Conversion(ctx context.Context, file_path string) error {
	runtime.LogPrint(ctx, file_path)
	if !strings.Contains(file_path, ".csv") {
		return errors.New("error opening CSV file: only CSV file can be converted")
	}
	// 读取UTF-8编码的CSV文件
	csvFile, err := os.Open(file_path)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	newPath := strings.Replace(file_path, ".csv", "_result.csv", -1)
	// 创建GBK编码的输出文件
	gbkFile, err := os.Create(newPath)
	if err != nil {
		log.Fatalf("Error creating GBK output file: %v", err)
		return err
	}
	defer gbkFile.Close()

	writer := csv.NewWriter(transform.NewWriter(gbkFile, simplifiedchinese.GB18030.NewEncoder()))
	defer writer.Flush()

	// 逐行读取CSV文件并写入到GBK编码的输出文件
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV record: %v", err)
			return err
		}

		err = writer.Write(record)
		if err != nil {
			log.Fatalf("Error writing CSV record: %v", err)
			return err
		}
	}
	return nil
}
