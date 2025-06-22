<script setup lang="ts">
import { useRouter } from 'vue-router'
const router = useRouter()

function goToeditpage() {
  router.push('/rooms/edit/date-and-time')
}

function goTovotepage() {
  router.push('/rooms/:room_id/places')
}

import BasicButton from '@/components/BasicButton.vue'
</script>

<template>
  <div class="py-6 max-w-[400px] max-h-[800px]">
    <div class="rounded-lg p-8">
      <div class="text-center">
        <h1 class="text-3xl font-bold text-gray-900 mb-4 mt-18">これはアプリタイトル</h1>
        <p class="text-lg text-gray-600 mb-[196px]">これはアプリサブタイトル</p>
        <div class="flex flex-col items-center space-y-4">
          <BasicButton
            @click="goToeditpage"
            text="ルームを作成"
            variant="primary"
            size="large"
            leftIcon="Edit_Calender"
          />
          <BasicButton
            @click="goTovotepage"
            text="ルームに参加"
            variant="secondary"
            size="large"
            leftIcon="door_open"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<style>
/* スタイル */
</style>
