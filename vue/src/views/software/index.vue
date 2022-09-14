<template>
  <div class="app-container">
    <el-table
      :data="list"
      border
      style="width: 100%"
    >
      <el-table-column
        prop="name"
        label="名称"
      />
      <el-table-column
        prop="desc"
        label="描述"
      />
      <el-table-column
        label="操作"
      />
    </el-table>
  </div>
</template>

<script>

import { softList } from '@/api/soft'

export default {
  name: 'Software',
  data() {
    return {
      list: []
    }
  },
  mounted() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      softList().then(response => {
        this.list = response.data.list
      })
    }
  }
}
</script>

