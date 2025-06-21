<script setup lang="ts">
import { useRouter } from 'vue-router'
const router = useRouter()

function goToeditpage() {
  router.push('/rooms/edit/date-and-time')
}

function goTovotepage() {
  router.push('/rooms/:room_id/places')
}
</script>

<template>
  <div class="py-6 max-w-[400px] max-h-[800px]">
    <div class="rounded-lg p-8">
      <div class="text-center">
        <h1 class="text-4xl font-bold text-gray-900 mb-4">これはアプリタイトル</h1>
        <p class="text-lg text-gray-600 mb-[196px]">これはアプリサブタイトル</p>
        <div class="flex flex-col items-center space-y-4">
          <button
            @click="goToeditpage"
            class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded transition-colors"
          >
            ルームを作成
          </button>
          <button
            @click="goTovotepage"
            class="bg-white hover:bg-gray-300 text-blue-500 font-bold py-2 px-4 rounded transition-colors"
          >
            ルームに参加
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
/* スタイル */
</style>
