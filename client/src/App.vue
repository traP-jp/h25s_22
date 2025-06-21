<script setup lang="ts">
import { RouterView } from 'vue-router'
import { ref, onMounted, onUnmounted } from 'vue'
import MenuIcon from '@/assets/icons/menu.svg'


const isMenuOpen = ref(false)
const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}
const handleMenuAction = (action: string) => {
  isMenuOpen.value = false

  switch (action) {
    case 'save':
      console.log('保存しました')
      break
    case 'reset':
      if (confirm('順序をリセットしますか？')) {
        console.log('リセットしました')
      }
      break
    case 'close':
      console.log('閉じました')
      break
  }
}

// メニューを閉じる関連の処理
const handleClickOutside = (event: Event) => {
  const target = event.target as HTMLElement
  if (!target.closest('.menu-container')) {
    isMenuOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})
onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})
</script>

<template>
  <div class="min-h-screen bg-gray-50 flex justify-center">
    <div class="max-w-md w-full relative">
      <div class="fixed top-0 left-1/2 transform -translate-x-1/2 w-full max-w-md z-50 bg-gray-100 h-9 flex items-center justify-end pr-3">
        <div class="menu-container relative">
          <button
            @click="toggleMenu"
            class="w-6 h-6 flex justify-center items-center hover:bg-gray-200 rounded transition-colors"
          >
            <img :src="MenuIcon" alt="Menu" class="w-6 h-6" />
          </button>

          <!-- ドロップダウンメニュー -->
          <div
            v-if="isMenuOpen"
            class="absolute top-6 right-0 bg-white border border-gray-200 rounded-lg shadow-lg py-2 min-w-32 z-10"
          >
            <button
              @click="handleMenuAction('save')"
              class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors"
            >
              保存
            </button>
            <button
              @click="handleMenuAction('reset')"
              class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors"
            >
              リセット
            </button>
            <button
              @click="handleMenuAction('close')"
              class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors"
            >
              閉じる
            </button>
          </div>
        </div>
      </div>

      <main class="pt-9">
        <RouterView />
      </main>
    </div>
  </div>
</template>

<style scoped></style>
