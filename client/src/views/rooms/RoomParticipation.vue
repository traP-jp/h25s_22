<script setup lang="ts">
import { ref } from 'vue'
import BasicInput from '@/components/BasicInput.vue'
import BasicButton from '@/components/BasicButton.vue'
const room_id = ref('')

function participate(roomid: string) {
  if (roomid !== '') {
    location.href = 'https://google.com'
  }
}
</script>

<template>
  <div class="flex flex-col gap-5 header justify-center items-center">
    <h1 class="w-75 size-9">ルーム参加</h1>
    <div class="w-75 flex flex-col gap-1">
      <div>ルームIDを入力</div>
      <BasicInput v-model="room_id" />
    </div>
    <BasicButton text="参加" @click="participate(room_id)" right-icon="login" />
  </div>
</template>

<style>
.header {
  font-weight: bold;
}
</style>
