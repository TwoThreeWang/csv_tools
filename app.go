package main

import (
	"context"
	"encoding/json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/ioutil"
	"net/http"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type Response struct {
	Data string `json:"data"`
	Info string `json:"info"`
}

// EscapeConver 科学计数法+编码转换
func (a *App) EscapeConver(filePath string) Response {
	err := EscapeConversion(filePath)
	response := Response{
		Data: "",
		Info: "",
	}
	if err != nil {
		response.Info = err.Error()
		return response
	}
	response.Info = "转换成功，结果文件跟原文件在同一个目录下！"
	return response
}

func (a *App) SelectFile() Response {
	dialogOption := runtime.OpenDialogOptions{
		Title: "选择要转换的CSV文件",
	}
	response := Response{
		Data: "",
		Info: "",
	}
	filePath, err := runtime.OpenFileDialog(a.ctx, dialogOption)
	if err != nil {
		response.Info = err.Error()
		return response
	}
	response.Data = filePath
	response.Info = "点击转换按钮开始转换"
	if filePath == "" {
		response.Info = "未选择文件，先选择文件 或 填写文件路径⬆"
	}
	return response
}

type TagResponse struct {
	TagName string `json:"tag_name"` // 将 JSON 中的 tag_name 映射到 Go 结构体的 TagName 字段
}

// Version 当前程序版本号
const Version = "3.0.0"

func (a *App) CheckUpdate() Response {
	response := Response{
		Data: "更新检测失败",
		Info: "",
	}
	// 发送 GET 请求获取最新的版本号
	url := "https://api.github.com/repos/TwoThreeWang/csv_tools/releases/latest"
	resp, err := http.Get(url)
	if err != nil {
		response.Info = err.Error()
		return response
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.Info = err.Error()
		return response
	}

	// 解析 JSON 数据
	var responseData TagResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		response.Info = err.Error()
		return response
	}

	if responseData.TagName != Version {
		response.Data = "程序有更新啦！当前版本：" + Version + "，最新版本：" + responseData.TagName
		response.Info = "https://github.com/TwoThreeWang/csv_tools/releases/latest"
	} else {
		response.Data = "已经是最新版啦！"
	}
	return response
}
