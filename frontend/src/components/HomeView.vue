<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.file_path" autocomplete="off" class="input" type="text"/>
      <br>
      <button class="btn" @click="select_file">选择</button>
      <button class="btn" @click="conver">转换</button>
    </div>
    <br>
    <small><RouterLink to="/about">About</RouterLink></small>
  </main>
</template>

<script setup>
import {reactive} from 'vue'
import {Conver, SelectFile} from '../../wailsjs/go/main/App'

const data = reactive({
  file_path: "",
  resultText: "先选择 CSV 文件 👇",
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

function conver() {
  Conver(data.file_path).then(result => {
    data.resultText = result.info
  })
}

</script>
<style>
main{
    text-align: center;
}
</style>