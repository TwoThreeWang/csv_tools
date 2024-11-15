# CSV encode conver

CSV Tools，一个 CSV 文件转换小工具

本工具用于将 CSV 文件转为 GBK 编码，以解决直接使用 Excel 打开 CSV 文件乱码的问题。

更新记录：
- v2.0 增加科学计数法处理功能；解决字段带引号转码报错问题。
- v1.0 第一个版本诞生

使用 Walis + VUE 开发

- 本地测试命令：wails dev 
- 打包命令：wails build

Wails文档：https://wails.io/zh-Hans/docs/


```自动打包推送
# 1. 提交代码
git add .
git commit -m "feat: new features"

# 2. 创建新版本标签
git tag v1.0.1

# 3. 推送代码和标签
git push origin main
git push origin v1.0.1
```