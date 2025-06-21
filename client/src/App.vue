<script setup lang="ts">
import { RouterView, useRouter } from 'vue-router'
import { ref, onMounted, onUnmounted } from 'vue'
import { useMainStore } from '@/stores'
import MenuIcon from '@/assets/icons/menu.svg'
import HomeIcon from '@/assets/icons/home.svg'
import EditCalendarIcon from '@/assets/icons/edit_calendar.svg'
import DoorOpenIcon from '@/assets/icons/door_open.svg'
import CloseIcon from '@/assets/icons/close.svg'

const router = useRouter()
const mainStore = useMainStore()
const isMenuOpen = ref(false)
const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}
const handleMenuAction = (action: string) => {
  isMenuOpen.value = false

  switch (action) {
    case 'home':
      router.push('/')
      break
    case 'create-room':
      router.push('/rooms/edit/date-and-time')
      break
    case 'join-room':
      // TODO: 不適切
      router.push('/rooms/123/places')
      break
    case 'close':
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
      <div
        class="fixed top-0 left-1/2 transform -translate-x-1/2 w-full max-w-md z-50 bg-gray-100 h-9 flex items-center justify-end pr-3"
      >
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
            class="absolute top-6 right-0 bg-white border border-gray-200 rounded-lg shadow-lg py-2 min-w-40 z-10"
          >
            <button
              @click="handleMenuAction('home')"
              class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors flex items-center gap-3"
            >
              <img :src="HomeIcon" alt="Home" class="w-4 h-4" />
              ホームへ
            </button>
            <button
              @click="handleMenuAction('create-room')"
              class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors flex items-center gap-3"
            >
              <img :src="EditCalendarIcon" alt="Create Room" class="w-4 h-4" />
              ルーム作成
            </button>
            <button
              @click="handleMenuAction('join-room')"
              class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors flex items-center gap-3"
            >
              <img :src="DoorOpenIcon" alt="Join Room" class="w-4 h-4" />
              ルーム参加
            </button>
            <hr class="my-1 border-gray-200" />
            <button
              @click="handleMenuAction('close')"
              class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors flex items-center gap-3"
            >
              <img :src="CloseIcon" alt="Close" class="w-4 h-4" />
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
