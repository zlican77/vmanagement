import { createRouter, createWebHistory } from 'vue-router';
import HomeComp from '../components/HomeComp.vue';
import GoodsComp from '../components/GoodsComp.vue';
import InboundComp from '../components/InboundComp.vue';
import OutboundComp from '../components/OutboundComp.vue';
import InventoryComp from '../components/InventoryComp.vue';
import CheckIn from '../components/CheckIn.vue';
import InGood from '../components/InGood.vue';
import GoodInfo from '../components/GoodInfo.vue';
import OutGood from '../components/OutGood.vue';
import { useAuth } from '../config/auth'

const routes = [
    { 
        path: '/', 
        name: 'HomeComp',
        component: HomeComp 
    },
    {
        path: '/goods',
        name: 'GoodsComp',
        component: GoodsComp,
        meta: { requiresAuth: true }
    },
    {
        path: '/checkin',
        name: 'CheckIn',
        components: {
            default: CheckIn,
            GoodInfo: GoodInfo
        },
        meta: { requiresAuth: true }
    },
    {
        path: '/inbound',
        name: 'InboundComp',
        component: InboundComp,
        meta: { requiresAuth: true }
    },
    {   // 双组件页面
        path: '/ingood',
        name: 'InGood',
        components: {
            default: InGood,
            GoodInfo: GoodInfo
        },
        meta: { requiresAuth: true }
    },
    {
        path: '/outbound',
        name: 'OutboundComp',
        component: OutboundComp,
        meta: { requiresAuth: true }
    },
    {
        path: '/outgood',
        name: 'OutGood',
        components: {
            default: OutGood,
            GoodInfo: GoodInfo
        },
        meta: { requiresAuth: true }
    },
    {
        path: '/inventory',
        name: 'InventoryComp',
        component: InventoryComp,
        meta: { requiresAuth: true }
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})


// 路由守卫
router.beforeEach((to, from, next) => {
    const auth = useAuth()
    const token = localStorage.getItem('token')
    
    if (token && !auth.isAuthenticated.value) {
        auth.setAuth(token)
    }
    
    if (to.meta.requiresAuth && !auth.isAuthenticated.value) {
        alert('请先完成登录')
        next('/')
    } else if (to.path === '/' && auth.isAuthenticated.value) {
        next('/goods')
    } else {
        next()
    }
})

export default router;
