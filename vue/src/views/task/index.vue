<template>
  <div class="app-container">
    <el-button type="success" @click="addTaskDialogVisible = true">新增任务</el-button>
    <el-dialog
      title="新增任务"
      :visible.sync="addTaskDialogVisible"
      width="55%"
    >
      <el-form ref="form" :model="addTaskForm" label-width="80px">
        <el-form-item label="任务名称" label-width="100px">
          <el-input v-model="addTaskForm.name" />
        </el-form-item>
        <el-form-item label="Spec" label-width="100px">
          <el-input v-model="addTaskForm.spec" />
        </el-form-item>
        <el-form-item label="脚本" label-width="100px">
          <el-input v-model="addTaskForm.data" type="textarea" rows="5" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="addTaskSubmit">确认</el-button>
          <el-button @click="addTaskDialogVisible = false">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <el-table
      :data="list"
      border
      style="width: 100%; margin-top: 15px;"
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

import { taskList, taskAdd } from '@/api/task'
import { Message } from 'element-ui'

export default {
  name: 'Task',
  data() {
    return {
      page_size: 10,
      page_index: 1,
      list: [],
      count: 0,
      addTaskDialogVisible: false,
      addTaskForm: {}
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
    },
    addTaskSubmit() {
      taskAdd(this.addTaskForm).then(response => {
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

