<template>
  <div class="dashboard-editor-container">
    <el-row :gutter="32">
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <gauge-chart :value="cpu" />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <gauge-chart :value="mem" />
        </div>
      </el-col>
      <el-col :xs="24" :sm="24" :lg="8">
        <div class="chart-wrapper">
          <gauge-chart :value="disk" />
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import GaugeChart from '@/views/dashboard/admin/components/GaugeChart'
import { state } from '@/api'

export default {
  name: 'DashboardAdmin',
  components: {
    GaugeChart
  },
  data() {
    return {
      cpu: {
        value: 0,
        name: 'CPU'
      },
      mem: {
        value: 0,
        name: 'MEM'
      },
      disk: {
        value: 0,
        name: 'DISK'
      },
      currentTimer: null
    }
  },
  mounted() {
    this.fetchData()
  },
  beforeDestroy() {
    clearInterval(this.currentTimer)
  },
  methods: {
    fetchData() {
      const _this = this
      this.currentTimer = setInterval(function() {
        state().then(response => {
          _this.cpu.value = response.data.cpu_used_percent
          _this.mem.value = response.data.mem_used_percent
          _this.disk.value = response.data.disk_used_percent
        })
      }, 2000)
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .github-corner {
    position: absolute;
    top: 0px;
    border: 0;
    right: 0;
  }

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

@media (max-width:1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
