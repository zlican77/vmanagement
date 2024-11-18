<template>
  <div class="outbound-container">
    <h3>出库登记</h3>
    <form class="outbound-form" @submit.prevent="handleSubmit">
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
        <label for="outboundTime">出库时间</label>
        <input
          type="datetime-local"
          id="outboundTime"
          v-model="formData.outboundTime"
          required
        >
      </div>

      <div class="form-group">
        <label for="quantity">出库数量</label>
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
        <button type="submit" class="submit-btn">出库登记</button>
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
        outboundTime: '',
        quantity: '',
        operator: ''
      }
    }
  },
  methods: {
    async handleSubmit() {
      try {
        await vmApi.Outbound(this.formData)
        this.$router.push('/outbound')
      } catch (error) {
        console.error('出库登记失败:', error)
      }
    },
    goBack() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
.outbound-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 20px;
}

.outbound-container h3 {
  text-align: center;
  color: #2c3e50;
  margin-bottom: 20px;
}

.outbound-form {
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
  font-weight: 500;
}

.form-group input {
  width: 100%;
  padding: 10px;
  border: 1px solid #dcdfe6;
  border-radius: 4px;
  transition: border-color 0.2s;
}

.form-group input:focus {
  outline: none;
  border-color: #409eff;
}

.button-group {
  display: flex;
  justify-content: center;
  gap: 15px;
  margin-top: 25px;
}

.submit-btn,
.back-btn {
  padding: 10px 25px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  transition: all 0.3s;
}

.submit-btn {
  background-color: #409eff;
  color: white;
}

.submit-btn:hover {
  background-color: #66b1ff;
}

.back-btn {
  background-color: #909399;
  color: white;
}

.back-btn:hover {
  background-color: #a6a9ad;
}
</style>