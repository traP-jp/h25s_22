import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { User } from '@/types'

export const useMainStore = defineStore('main', () => {
  // アプリケーション全体の状態
  const isLoading = ref(false)
  const user = ref<User | null>(null)

  // アクション
  const setLoading = (loading: boolean) => {
    isLoading.value = loading
  }

  const setUser = (userData: User | null) => {
    user.value = userData
  }

  return {
    // 状態
    isLoading,
    user,
    // アクション
    setLoading,
    setUser,
  }
})
