<template>
  <div class="container">
    <header class="header">
            <nav class="nav-menu">
                <div class="logo">
                    <router-link to="/">仓库管理系统</router-link>
                </div>
                <div class="menu-items" v-if="auth.isAuthenticated">
                    <router-link to="/goods">货物管理</router-link>
                    <router-link to="/inbound">入库管理</router-link>
                    <router-link to="/outbound">出库管理</router-link>
                    <router-link to="/inventory">库存管理</router-link>
                    <a href="#" @click.prevent="handleLogout">退出系统</a>
                </div>
            </nav>
        </header>
        

    <main class="main-content">
      <div v-if="$route.path === '/ingood' || $route.path === '/checkin' || $route.path === '/outgood'" class="dual-component-container">
      <div class="component-left">
        <router-view />
      </div>
      <div class="component-right">
        <router-view name="GoodInfo" />
      </div>
    </div>

    <div v-else>
      <router-view v-slot="{ Component }">
        <component :is="Component" />
      </router-view>
    </div>
    </main>

    <footer class="footer">
      <p>© 2024 库存管理系统 | 保留所有权利</p>
    </footer>



  </div>
</template>

<script>
import { useAuth } from './config/auth'
import { useRouter } from 'vue-router'

export default {
    setup() {
        const auth = useAuth()
        const router = useRouter()

        const handleLogout = () => {
            auth.clearAuth()
            router.push('/')
        }

        return {
            auth,
            handleLogout
        }
    }
}
</script>

<style>
/* 全局样式重置 */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

/* 容器样式 */
.container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: #f5f7fa;
}

/* 头部样式 */
.header {
  background-color: #ffffff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  padding: 1rem 0;
  position: sticky;
  top: 0;
  z-index: 100;
}

.nav-menu {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 1rem;
}

.logo a {
  font-size: 1.5rem;
  font-weight: bold;
  color: #2c3e50;
  text-decoration: none;
  transition: color 0.3s ease;
}

.logo a:hover {
  color: #42b983;
}

.menu-items {
  display: flex;
  gap: 2rem;
}

.menu-items a {
  color: #606266;
  text-decoration: none;
  font-size: 1rem;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  transition: all 0.3s ease;
}

.menu-items a:hover {
  color: #42b983;
  background-color: #f0f9f6;
}

.menu-items a.router-link-active {
  color: #42b983;
  font-weight: bold;
}

/* 主要内容区域 */
.main-content {
  flex: 1;
  max-width: 1200px;
  margin: 2rem auto;
  padding: 0 1rem;
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

/* 页脚样式 */
.footer {
  background-color: #2c3e50;
  color: #ffffff;
  text-align: center;
  padding: 1.5rem 0;
  margin-top: auto;
}

.footer p {
  font-size: 0.9rem;
  opacity: 0.8;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .nav-menu {
    flex-direction: column;
    gap: 1rem;
  }

  .menu-items {
    flex-direction: column;
    gap: 1rem;
    width: 100%;
    text-align: center;
  }

  .menu-items a {
    display: block;
    padding: 0.8rem;
  }
}

.dual-component-container {
  display: flex;
  gap: 20px;
  margin-top: 20px;
}

.component-left {
  flex: 1;
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  padding: 20px;
}

.component-right {
  flex: 1;
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  padding: 20px;
}
</style>