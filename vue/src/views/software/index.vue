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
      >
        <template slot-scope="scope">
          <el-button type="text" @click="handleInstall(scope.row)">安装</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>

import { softList, softOperation } from '@/api/soft'
import { Message } from 'element-ui'

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
    },
    handleInstall(row) {
      softOperation({ 'op': 'install', 'id': row.ID }).then(response => {
        Message({
          message: response.msg,
          type: 'success',
          duration: 3 * 1000
        })
      })
    }
  }
}
</script>

