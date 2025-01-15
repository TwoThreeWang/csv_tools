<template>
  <div class="container" id="home-page">
    <h1>CSV 转换工具</h1>
    <p>将 CSV 文件转换为 Excel 编码格式</p>
    <!-- 文件上传区域 -->
    <div class="upload-area" >
      <input id="name" v-model="data.file_path" autocomplete="off" class="input" type="text" placeholder="选择文件或填入文件路径"/>
      <p class="file-name" id="file-name" @click="select_file">{{ data.resultText }}</p>
    </div>

    <!-- 转换按钮 -->
    <button class="btn" @click="select_file">选择文件</button>
    <button id="convert-btn" class="btn" @click="escape_conver">转换</button>
    <!-- 跳转到关于页面按钮 -->
    <RouterLink to="/about"><button id="go-about-btn" class="btn">关于</button></RouterLink>
  </div>
</template>

<script setup>
import {reactive} from 'vue'
import {SelectFile, EscapeConver} from '../../wailsjs/go/main/App'

const data = reactive({
  file_path: "",
  resultText: "点击选择文件 或 填写文件路径⬆️",
})

function select_file() {
  SelectFile().then(result => {
  console.log(result);
    if(result.info!=""){
        data.resultText = result.info
    }
    data.file_path = result.data
  })
}

function escape_conver() {
  data.resultText = "文件转换中..."
  EscapeConver(data.file_path).then(result => {
    data.resultText = result.info
  })
}

</script>
<style>
main{
    text-align: center;
}
</style>