<!-- About.vue -->
<template>
  <div class="card">
    <router-link to="/">《 返回首页</router-link>
    <h3>CSV Tools，一个 CSV 文件转换小工具</h3>
    <p>本工具用于将 CSV 文件转为 GBK 编码，以解决直接使用 Excel 打开 CSV 文件乱码的问题。</p>
    <li>一键转换 CSV 编码</li>
    <li>一键转换 CSV 编码</li>
    <details>
        <summary>常见的水果</summary>
        <ol>
            <li>草莓</li>
            <li>香蕉</li>
            <li>苹果</li>
        </ol>
    </details>
    <div class="alert" role="alert">
        Development Team: WangCheng
        <br>
        Version: 2.0
        <br>
        Github: <a href="https://github.com/TwoThreeWang/csv_encode_conver">https://github.com/TwoThreeWang/csv_encode_conver</a>
    </div>
  </div>
  <br>
  <div class="card">
      <h3>版本更新记录</h3>
      <div class="alert" role="alert">
          <li>20241114 v2.0 增加科学计数法处理功能；解决文件格式错误闪退问题。</li>
          <li>20241113 v1.0 软件诞生</li>
      </div>
    </div>
</template>

<script>
export default {
  name: 'About'
}
</script>