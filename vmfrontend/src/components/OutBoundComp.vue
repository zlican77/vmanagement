<template>
  <div class="outbound-container">
    <h3>出库管理</h3>

    <div class="outbound-operations">
      <button class="register-btn" @click="registerOutbound">出库登记</button>
    </div>

    <div class="outbound-table">
      <table>
        <thead>
          <tr>
            <th>序号</th>
            <th>出库时间</th>
            <th>货物号</th>
            <th>货物名</th>
            <th>规格</th>
            <th>型号</th>
            <th>出库数量</th>
            <th>经办人</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in outboundList" :key="item.id">
            <td>{{ index + 1 }}</td>
            <td>{{ item.outboundTime }}</td>
            <td>{{ item.goodsId }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.spec }}</td>
            <td>{{ item.model }}</td>
            <td>{{ item.quantity }}</td>
            <td>{{ item.operator }}</td>
            <td>
              <button class="edit-btn" @click="editOutbound(item)">修改</button>
              <button class="delete-btn" @click="deleteOutbound(item)">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="outbound-count">
        当前货物数是：{{ outboundList.length }}
      </div>
    </div>
  </div>
</template>

<script>
import { vmApi } from '../vmapi/vmapi'
export default {
  data() {
    return {
      outboundList: []
    }
  },
  created() {
    this.fetchOutboundList()
  },
  methods: {
    async fetchOutboundList() {
      const res = await vmApi.outboundGoods()
      res.data.outboundGoods.forEach(item => {
        item.outboundTime = new Date(item.outboundTime).toLocaleString()
      })  // 将时间格式化
      this.outboundList = res.data.outboundGoods
    },  
    registerOutbound() {
      this.$router.push('/outgood')
    },
    editOutbound(item) {
      // 实现编辑出库记录逻辑
    },
    deleteOutbound(item) {
      // 实现删除出库记录逻辑
    }
  }
}
</script>

<style>
.outbound-container {
  padding: 20px;
}

.outbound-operations {
  margin: 20px 0;
}

.register-btn {
  padding: 8px 16px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.outbound-table table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
}

.outbound-table th,
.outbound-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

.outbound-table th {
  background-color: #f5f5f5;
}

.edit-btn,
.delete-btn {
  margin: 0 5px;
  padding: 4px 8px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.edit-btn {
  background-color: #2196F3;
  color: white;
}

.delete-btn {
  background-color: #f44336;
  color: white;
}

.outbound-count {
  margin-top: 20px;
  font-weight: bold;
}
</style>