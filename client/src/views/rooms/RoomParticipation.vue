<script setup lang="ts">
import { ref, onMounted } from 'vue'
import BasicInput from '@/components/BasicInput.vue'
import BasicButton from '@/components/BasicButton.vue'
import { useRoute, useRouter } from 'vue-router'
import { getRoomDetails } from '@/services/api'

const route = useRoute()
const router = useRouter()
const path = ref(route.path)
const room_id = ref(path)
const isLoading = ref(false)

onMounted(() => {
  room_id.value = room_id.value.replace('/room-participation', '')
  room_id.value = room_id.value.replace('/rooms/', '')
  if (room_id.value === 'Default') {
    room_id.value = ''
  }
})

async function participate(roomid: string) {
  if (roomid.trim() === '') {
    alert('ルームIDを入力してください。')
    return
  }

  isLoading.value = true

  try {
    // GetRoom APIを叩いてルームの存在を確認
    const roomData = await getRoomDetails(roomid.trim())

    // APIが成功した場合、場所選択画面に遷移
    router.push(`/rooms/${roomid.trim()}/places`)
  } catch (error) {
    console.error('ルーム取得エラー:', error)
    alert('ルームが見つかりません。ルームIDを確認してください。')
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="flex flex-col gap-5 header justify-center items-center">
    <h1 class="w-75 size-9">ルーム参加</h1>
    <div class="w-75 flex flex-col gap-1">
      <div>ルームIDを入力</div>
      <BasicInput v-model="room_id" :disabled="isLoading" />
    </div>
    <BasicButton
      :text="isLoading ? '確認中...' : '参加'"
      @click="participate(room_id)"
      right-icon="login"
      :disabled="isLoading"
    />
  </div>
</template>

<style>
.header {
  font-weight: bold;
}
</style>
