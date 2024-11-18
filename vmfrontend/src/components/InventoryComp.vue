<template>
  <div class="inventory-container">
    <h3>库存管理</h3>

    <div class="inventory-table">
      <table>
        <thead>
          <tr>
            <th>序号</th>
            <th>货物号</th>
            <th>货物名</th>
            <th>规格</th>
            <th>型号</th>
            <th>更新日期</th>
            <th>库存量</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in inventoryList" :key="item.id">
            <td>{{ index + 1 }}</td>
            <td>{{ item.goodsId }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.spec }}</td>
            <td>{{ item.model }}</td>
            <td>{{ item.updateTime }}</td>
            <td>{{ item.quantity }}</td>
          </tr>
        </tbody>
      </table>
      <div class="inventory-count">
        当前货物数是：{{ inventoryList.length }}
      </div>
    </div>
  </div>
</template>

<script>
import { vmApi } from '../vmapi/vmapi'
export default {
  created() {
    this.fetchInventoryList()
  },
  data() {
    return {
      inventoryList: []
    }
  },
  methods: {
    async fetchInventoryList() {
      const res = await vmApi.inventoryGoods()
      res.data.inventoryGoods.forEach(item => {
        item.updateTime = new Date(item.updateTime).toLocaleString()
      })  // 将时间格式化
      this.inventoryList = res.data.inventoryGoods
    }
  }
}
</script>

<style>
.inventory-container {
  padding: 20px;
}

.inventory-table table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
}

.inventory-table th,
.inventory-table td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

.inventory-table th {
  background-color: #f5f5f5;
}

.inventory-count {
  margin-top: 20px;
  font-weight: bold;
}
</style>