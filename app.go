package main

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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

// Conver returns a greeting for the given name
func (a *App) Conver(filePath string) Response {
	err := Conversion(a.ctx, filePath)
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
	return response
}
