import axios from 'axios'

// 创建 axios 实例
const request = axios.create({
  baseURL: 'http://127.0.0.1:8000/vmapi', // 基础URL
  timeout: 5000,  // 请求超时时间
  headers: {
    'Content-Type': 'application/json'
  }
})

export default request
