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
        prop="spec"
        label="spec"
      />
      <el-table-column
        prop="CreatedAt"
        label="创建时间"
      />
      <el-table-column
        label="操作"
      />
    </el-table>
    <div style="margin-top: 15px">
      <el-pagination
        :page-size="page_size"
        layout="total, prev, pager, next"
        :total="count"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script>

import { taskList } from '@/api/task'

export default {
  name: 'Task',
  data() {
    return {
      page_size: 10,
      page_index: 1,
      list: [],
      count: 0
    }
  },
  mounted() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      taskList({ index: this.page_index, size: this.page_size }).then(response => {
        this.list = response.data.list
        this.count = response.data.count
      })
    },
    handleCurrentChange(index) {
      this.page_index = index
      this.fetchData()
    }
  }
}
</script>

