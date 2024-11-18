import { ref } from 'vue'

export const useAuth = () => {
    const isAuthenticated = ref(!!localStorage.getItem('token'))
    const token = ref(localStorage.getItem('token') || '')

    const setAuth = (newToken) => {
        token.value = newToken
        isAuthenticated.value = true
        localStorage.setItem('token', newToken)
    }

    const clearAuth = () => {
        token.value = ''
        isAuthenticated.value = false
        localStorage.removeItem('token')
    }

    return {
        isAuthenticated,
        token,
        setAuth,
        clearAuth
    }
}