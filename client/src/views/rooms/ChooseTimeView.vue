<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
const router = useRouter()

function goTokouhotipage() {
  router.push('/rooms/:room_id/reorder-by-preference')
}

import BasicButton from '@/components/BasicButton.vue'
const isActiveList = ref([false, false, false])

const toggle = (index: number) => {
  isActiveList.value[index] = !isActiveList.value[index]
}
</script>

<template>
  <div class="py-6 max-w-[400px] max-h-[800px]">
    <div class="p-4">
      <h1 class="w-[200px] h-[28px] text-gray-900 mb-4">参加可能な時間を選択</h1>
    </div>
    <div class="flex-col space-y-2 rounded-none px-4">
      <div v-for="(active, index) in isActiveList" :key="index" class="flex items-center">
        <button
          @click="toggle(index)"
          :class="active ? 'bg-sky-400 text-white' : 'bg-gray-200 text-gray-900'"
          class="py-4 px-6 min-w-[245px] text-left"
        >
          日付<br />時間帯
        </button>

        <button
          @click="toggle(index)"
          :class="active ? 'bg-sky-400 text-white' : 'bg-gray-200 text-gray-900'"
          class="flex items-center justify-center w-13.5 h-20 pr-4"
        >
          <div class="flex items-center justify-center w-5 h-5 rounded-full bg-white border">
            <div v-if="active" class="w-3 h-3 rounded-full bg-sky-400"></div>
          </div>
        </button>
      </div>
    </div>

    <div class="flex rounded-lg p-4 items-center">
      <BasicButton @click="goTokouhotipage" text="次へ  →" variant="primary" size="large" />
    </div>
  </div>
</template>
