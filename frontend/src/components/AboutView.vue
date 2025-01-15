<!-- About.vue -->
<template>
    <!-- 关于页面 -->
    <div class="container about-page" id="about-page">
        <router-link to="/">《 返回首页</router-link>
        <h2>关于</h2>
        <p>CSV Tools，一个简单的 CSV 转换工具，支持将 CSV 文件转换为 Excel 编码格式。工具完全基于本地运行，无需上传文件到服务器，确保数据安全。</p>
        <p>功能特点：</p>
        <ul>
            <li>一键转换 CSV 编码，解决 Excel 打开 CSV 乱码问题</li>
            <li>解决数字科学计数法问题</li>
            <li>数据本地处理，无需上传</li>
        </ul>
        <p>软件信息：</p>
        <ul>
            <li>版本：3.0.0</li>
            <li>Development Team: WangCheng</li>
            <li>Version: 3.0.0</li>
            <li>Github: https://github.com/TwoThreeWang/csv_tools</li>
        </ul>
        <div class="update_info">
            <button class="btn" @click="check_update">点击检测更新</button>
            <p id="check_update_data"></p>
            <p id="check_update_info"></p>
        </div>
        <p>更新记录：</p>
        <ul>
            <li>20250115 v3.0 优化界面</li>
            <li>20241114 v2.0 增加科学计数法处理功能；解决文件格式错误闪退问题。</li>
            <li>20241113 v1.0 软件诞生</li>
        </ul>
    </div>
</template>

<script setup>
import {CheckUpdate} from '../../wailsjs/go/main/App'

function check_update() {
    console.log("check")
    var check_update_data = document.getElementById("check_update_data");
    var check_update_info = document.getElementById("check_update_info");
    check_update_data.innerHTML = "版本检测中..."
    CheckUpdate().then(result => {
        check_update_data.innerHTML = result.data
        var infostr = result.info
        if(result.data.startsWith("程序有更新啦")){
            infostr = "去更新："+result.info
        }
        check_update_info.innerHTML = infostr
    }).catch(err => {
            check_update_data.innerHTML = "检测报错"
            check_update_info.innerHTML = err
        }
    )
}
</script>