import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'

export const useAuthStore = defineStore('auth', () => {
  const router = useRouter()
  const user = ref<string | null>(localStorage.getItem('user'))
  const isAuthenticated = computed(() => user.value !== null)

  const login = (username: string, password: string) => {
    if (username === 'admin' && password === 'password') {
      user.value = username
      localStorage.setItem('user', username)
      router.push({ name: 'projectList' })
    } else {
      alert('Incorrect credentials')
    }
  }

  const logout = () => {
    user.value = null
    localStorage.removeItem('user')
    router.push({ name: 'login' })
  }

  return { user, isAuthenticated, login, logout }
})
