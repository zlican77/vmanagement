<template>
  <div class="inbound-container">
    <h3>入库登记</h3>
    <form class="inbound-form" @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="goodsId">货物号</label>
        <input
          type="text" 
          id="goodsId"
          v-model="formData.goodsId"
          required
        >
      </div>

      <div class="form-group">
        <label for="inboundTime">入库时间</label>
        <input
          type="datetime-local"
          id="inboundTime"
          v-model="formData.inboundTime"
          required
        >
      </div>

      <div class="form-group">
        <label for="quantity">入库数量</label>
        <input
          type="number"
          id="quantity"
          v-model="formData.quantity"
          required
        >
      </div>

      <div class="form-group">
        <label for="operator">经办人</label>
        <input
          type="text"
          id="operator"
          v-model="formData.operator"
          required
        >
      </div>

      <div class="button-group">
        <button type="submit" class="submit-btn">入库登记</button>
        <button type="button" class="back-btn" @click="goBack">返回</button>
      </div>
    </form>
  </div>
</template>

<script>
import { vmApi } from '../vmapi/vmapi'

export default {
  data() {
    return {
      formData: {
        goodsId: '',
        inboundTime: '',
        quantity: '',
        operator: ''
      },
      goodsInfo: null
    }
  },
  methods: {
    async handleSubmit() {
      try {
        await vmApi.inbound(this.formData)
        this.$router.push('/inbound')
      } catch (error) {
        console.error('入库登记失败:', error)
      }
    },
    goBack() {
      this.$router.go(-1)
    },
  },
  watch: {
    'formData.goodsId': {
      handler: 'fetchGoodsInfo',
      immediate: true
    }
  }
}
</script>

<style>
.inbound-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 20px;
}

.inbound-container h3 {
  text-align: center;
  color: #2c3e50;
  margin-bottom: 20px;
}

.inbound-form {
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  color: #606266;
}

.form-group input {
  width: 100%;
  padding: 8px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
}

.button-group {
  display: flex;
  justify-content: center;
  gap: 15px;
  margin-top: 20px;
}

.submit-btn,
.back-btn {
  padding: 8px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.submit-btn {
  background-color: #409eff;
  color: white;
}

.back-btn {
  background-color: #909399;
  color: white;
}

.goods-info {
  margin-top: 20px;
}

.goods-info table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 10px;
}

.goods-info th,
.goods-info td {
  border: 1px solid #dcdfe6;
  padding: 8px;
  text-align: left;
}

.goods-info th {
  background-color: #f5f7fa;
}
</style>