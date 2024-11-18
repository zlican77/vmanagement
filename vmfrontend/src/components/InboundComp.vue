<template>
  <div class="inbound-container">
    <h3>入库管理</h3>

    <div class="inbound-operations">
      <button class="register-btn" @click="registerInbound">入库登记</button>
    </div>

    <div class="inbound-table">
      <table>
        <thead>
          <tr>
            <th>序号</th>
            <th>入库时间</th>
            <th>货物号</th>
            <th>货物名</th>
            <th>规格</th>
            <th>型号</th>
            <th>入库数量</th>
            <th>经办人</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in inboundList" :key="item.id">
            <td>{{ index + 1 }}</td>
            <td>{{ item.inboundTime }}</td>
            <td>{{ item.goodsId }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.spec }}</td>
            <td>{{ item.model }}</td>
            <td>{{ item.quantity }}</td>
            <td>{{ item.operator }}</td>
            <td>
              <button class="edit-btn" @click="editInbound(item)">修改</button>
              <button class="delete-btn" @click="deleteInbound(item)">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div class="inbound-count">
        当前货物数是：{{ inboundList.length }}
      </div>
    </div>
  </div>
</template>

<script>
import { vmApi } from '../vmapi/vmapi'
export default {
  data() {
    return {
      inboundList: []
    }
  },
  created() {
    this.fetchInboundList()
  },
  methods: {
    async fetchInboundList() {
      const res = await vmApi.inboundGoods()
      res.data.inboundGoods.forEach(item => {
        item.inboundTime = new Date(item.inboundTime).toLocaleString()
      })  // 将时间格式化
      this.inboundList = res.data.inboundGoods
    },
    registerInbound() {
      // 实现入库登记逻辑
      this.$router.push('/ingood')
    },
    editInbound(item) {
      // 实现编辑入库记录逻辑
    },
    deleteInbound(item) {
      // 实现删除入库记录逻辑
    }
  }
}
</script>

<style>
.inbound-container {
  padding: 20px;
}

.inbound-operations {
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

.inbound-table table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
}

.inbound-table th,
.inbound-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

.inbound-table th {
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

.inbound-count {
  margin-top: 20px;
  font-weight: bold;
}
</style>