<!-- About.vue -->
<template>
  <div class="card">
    <router-link to="/">《 返回首页</router-link>
    <h3>CSV Tools，一个 CSV 文件转换小工具</h3>
    <details>
        <summary>功能列表</summary>
        <ol>
            <li>一键转换 CSV 编码，解决 Excel 打开 CSV 乱码问题</li>
            <li>解决数字科学计数法问题</li>
        </ol>
    </details>
      <details>
          <summary>更新记录</summary>
          <ol>
              <li>20241114 v2.0 增加科学计数法处理功能；解决文件格式错误闪退问题。</li>
              <li>20241113 v1.0 软件诞生</li>
          </ol>
      </details>
  <details>
      <summary>软件信息</summary>
      <ol>
          <li>Development Team: WangCheng</li>
          <li>Version: 2.0</li>
          <li>Github: https://github.com/TwoThreeWang/csv_tools</li>
      </ol>
  </details>
  </div>
</template>

<script>
export default {
  name: 'About'
}
</script>