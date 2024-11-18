<template>
  <div class="register-container">
    <h3>货物登记</h3>
    <form class="register-form" @submit.prevent="handleSubmit">
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
        <label for="name">货物名</label>
        <input
          type="text"
          id="name"
          v-model="formData.name"
          required
        >
      </div>

      <div class="form-group">
        <label for="spec">规格</label>
        <input
          type="text"
          id="spec"
          v-model="formData.spec"
          required
        >
      </div>

      <div class="form-group">
        <label for="model">型号</label>
        <input
          type="text"
          id="model"
          v-model="formData.model"
          required
        >
      </div>

      <div class="form-group">
        <label for="desc">说明</label>
        <textarea
          id="desc"
          v-model="formData.desc"
          rows="2"
        ></textarea>
      </div>

      <div class="button-group">
        <button type="submit" class="submit-btn">登记</button>
        <button type="button" class="back-btn" @click="goBack">返回</button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      formData: {
        goodsId: '',
        name: '',
        spec: '',
        model: '',
        desc: ''
      }
    }
  },
  methods: {
    handleSubmit() {
      // 处理表单提交
      console.log(this.formData)
    },
    goBack() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
.register-container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 10px;
}

.register-container h3 {
  text-align: center;
  color: #2c3e50;
  margin-bottom: 0px;
  font-size: 20px;
}

.register-form {
  background-color: white;
  padding: 2px 25px;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.05);
}

.form-group {
  margin-bottom: 12px;
}

.form-group label {
  display: block;
  margin-bottom: 3px;
  color: #606266;
  font-size: 13px;
  font-weight: 400;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 6px 10px;
  border: 1px solid #ebeef5;
  border-radius: 4px;
  font-size: 13px;
  transition: all 0.2s ease;
  background-color: #fcfcfc;
}

.form-group input:focus,
.form-group textarea:focus {
  border-color: #85b8ff;
  background-color: #fff;
  outline: none;
  box-shadow: 0 0 0 1px rgba(64,158,255,.15);
}

.form-group textarea {
  resize: vertical;
  min-height: 50px;
}

.button-group {
  display: flex;
  justify-content: center;
  gap: 12px;
  margin-top: 20px;
}

.submit-btn,
.back-btn {
  padding: 6px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 13px;
  font-weight: 400;
  transition: all 0.2s ease;
}

.submit-btn {
  background-color: #85b8ff;
  color: white;
}

.back-btn {
  background-color: #a3a6ad;
  color: white;
}

.submit-btn:hover {
  background-color: #99c4ff;
  transform: translateY(-1px);
  box-shadow: 0 1px 6px rgba(64,158,255,.2);
}

.back-btn:hover {
  background-color: #b1b3b8;
  transform: translateY(-1px);
  box-shadow: 0 1px 6px rgba(144,147,153,.2);
}
</style>