<template>
    <div class="goods-container">
      <h3>货物信息</h3>
      <div class="goods-table">
        <table>
          <thead>
            <tr>
              <th>序号</th>
              <th>货物号</th>
              <th>货物名</th>
              <th>规格</th>
              <th>型号</th>
              <th>说明</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(item, index) in goodsList" :key="item.id">
              <td>{{ index + 1 }}</td>
              <td>{{ item.goodsId }}</td>
              <td>{{ item.name }}</td>
              <td>{{ item.spec }}</td>
              <td>{{ item.model }}</td>
              <td>{{ item.desc }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </template>
  
  <script>
  import { vmApi } from '../vmapi/vmapi'
  export default {
    data() {
      return {
        goodsList: []
      }
    },
    created() {
      this.fetchGoodsList()
    },
    methods: {
      async fetchGoodsList() {
        // 实现获取货物列表逻辑
        const res = await vmApi.getGoodsList()
        this.goodsList = res.data.goods
        console.log(this.goodsList)
      },
    }
  }
  </script>
  
  <style>
  .goods-container {
    padding: 20px;
  }
  
  .goods-operations {
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
  
  .goods-table table {
    width: 100%;
    border-collapse: collapse;
    margin-bottom: 20px;
  }
  
  .goods-table th,
  .goods-table td {
    border: 1px solid #ddd;
    padding: 8px;
    text-align: left;
  }
  
  .goods-table th {
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
  
  .goods-count {
    margin-top: 20px;
    font-weight: bold;
  }
  </style>