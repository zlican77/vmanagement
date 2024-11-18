<template>
    <div class="login-container">
        <div class="login-box">
            <h3>欢迎使用仓库管理系统</h3>
            <form @submit.prevent="handleLogin" class="login-form">
                <div class="form-group">
                    <label for="username">用户名</label>
                    <input 
                        type="text"
                        id="username"
                        v-model="username"
                        placeholder="请输入用户名"
                        required
                    >
                </div>
                <div class="form-group">
                    <label for="password">密码</label>
                    <input
                        type="password"
                        id="password" 
                        v-model="password"
                        placeholder="请输入密码"
                        required
                    >
                </div>
                <button type="submit" class="login-btn">登录</button>
            </form>
        </div>
    </div>
</template>

<script>
import { vmApi } from '../vmapi/vmapi'
import { useAuth } from '../config/auth'
import { useRouter } from 'vue-router'

export default {
    setup() {
        const auth = useAuth()
        const router = useRouter()
        return { auth, router }
    },
    data() {
        return {
            username: '',
            password: ''
        }
    },
    methods: {
        async handleLogin() {
            try {
                const loginData = {
                    username: this.username,
                    password: this.password
                }
                const res = await vmApi.login(loginData)
                this.auth.setAuth(res.data.token)
                this.$router.push('/goods')
            } catch (error) {
                alert("登入失败")
            }
        }
    }
}
</script>
<style>
.login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 70vh;
    background-color: #f5f7fa;
    padding: 20px;
}

.login-box {
    background-color: white;
    padding: 30px;
    border-radius: 8px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
    width: 400px;
}

.login-box h3 {
    text-align: center;
    color: #2c3e50;
    margin-bottom: 30px;
}

.login-form {
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.form-group {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.form-group label {
    color: #606266;
    font-size: 14px;
}

.form-group input {
    padding: 10px;
    border: 1px solid #dcdfe6;
    border-radius: 4px;
    font-size: 14px;
    transition: border-color 0.3s;
}

.form-group input:focus {
    outline: none;
    border-color: #42b983;
}

.login-btn {
    background-color: #42b983;
    color: white;
    padding: 10px;
    border: none;
    border-radius: 4px;
    font-size: 14px;
    cursor: pointer;
    transition: background-color 0.3s;
}

.login-btn:hover {
    background-color: #3aa876;
}
</style>