<!-- About.vue -->
<template>
  <div class="card">
    <router-link to="/">《 返回首页</router-link>
    <h3>CSV encode conver，一个 CSV 文件编码转 GBK 小工具</h3>
    <p>本工具用于将 CSV 文件转为 GBK 编码，以解决直接使用 Excel 打开 CSV 文件乱码的问题。</p>
    <div class="alert" role="alert">
        Development Team: WangCheng
        <br>
        Version: 1.0
        <br>
        Github: <a href="https://github.com/TwoThreeWang/csv_encode_conver">https://github.com/TwoThreeWang/csv_encode_conver</a>
    </div>
  </div>
</template>

<script>
export default {
  name: 'About'
}
</script>