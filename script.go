package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"regexp"
	"strings"
)

// Conversion 编码转换函数
func Conversion(file_path string) error {
	if !strings.Contains(file_path, ".csv") {
		return errors.New("Error: 仅支持 CSV 文件")
	}
	// 读取UTF-8编码的CSV文件
	csvFile, err := os.Open(file_path)
	if err != nil {
		return errors.New("Error: 文件打开失败")
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	newPath := strings.Replace(file_path, ".csv", "_result.csv", -1)
	// 创建GBK编码的输出文件
	gbkFile, err := os.Create(newPath)
	if err != nil {
		return errors.New("创建新文件失败: " + err.Error())
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
			return errors.New("文件格式错误: " + err.Error())
		}

		err = writer.Write(record)
		if err != nil {
			return errors.New("文件写入失败: " + err.Error())
		}
	}
	return nil
}

// EscapeConversion 科学计数法和编码转换
func EscapeConversion(file_path string) error {
	if !strings.Contains(file_path, ".csv") {
		return errors.New("Error: 仅支持 CSV 文件转换")
	}
	// 打开 CSV 文件
	file, err := os.Open(file_path)
	if err != nil {
		return errors.New("Error: 文件打开失败")
	}
	defer file.Close()

	// 创建输出文件
	newPath := strings.Replace(file_path, ".csv", "_result.csv", -1)
	outputFile, err := os.Create(newPath)
	if err != nil {
		return errors.New("创建新文件失败: " + err.Error())
	}
	defer outputFile.Close()
	writer := bufio.NewWriter(transform.NewWriter(outputFile, simplifiedchinese.GB18030.NewEncoder()))
	defer writer.Flush()
	// 使用 bufio.Scanner 逐行读取文件
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`^\d+$`)
	for scanner.Scan() {
		line := scanner.Text()
		// 分割行，得到每个字段
		fields := strings.Split(line, ",")
		// 处理每个字段
		for i, field := range fields {
			// 如果字段没有双引号，就加上双引号
			if strings.Contains(field, "\"") {
				// 如果两头都有引号，则不处理
				if strings.HasPrefix(field, "\"") && strings.HasSuffix(field, "\"") {
					num := strings.ReplaceAll(field, "\"", "")
					if re.MatchString(num) {
						fields[i] = "\"" + num + "\t\""
					}
					continue
				} else {
					// 如果中间有引号或者两头只有一个引号，则转义
					field = strings.ReplaceAll(field, "\"", "\\\"")
				}
			}
			if re.MatchString(field) {
				fields[i] = field + "\t"
			}
		}
		// 将处理后的字段拼接成一行
		newLine := strings.Join(fields, ",")

		// 将处理后的行写入输出文件，指定编码为 GBK
		_, err = writer.WriteString(newLine + "\n")
		if err != nil {
			return errors.New("文件写入失败: " + err.Error())
		}
	}

	if err := scanner.Err(); err != nil {
		return errors.New("文件打开失败: " + err.Error())
	}
	return nil
}
