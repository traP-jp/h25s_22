<script setup lang="ts">
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useRoomCreationStore } from '@/stores'
import type { TimeOption } from '@/types'
import BasicButton from '@/components/BasicButton.vue'

const router = useRouter()
const route = useRoute()
const roomCreationStore = useRoomCreationStore()

const selectedTimeIds = ref(new Set<string>())

// 選択状態を切り替える関数
const toggle = (slot: TimeOption) => {
  if (selectedTimeIds.value.has(slot.id)) {
    selectedTimeIds.value.delete(slot.id)
  } else {
    selectedTimeIds.value.add(slot.id)
  }
}

const handleNext = () => {
  if (selectedTimeIds.value.size === 0) {
    alert('参加可能な時間を1つ以上選択してください。')
    return
  }
  const roomId = route.params.room_id as string
  if (!roomId) {
    console.error('ルームIDが見つかりません。')
    return
  }
  const selectedIds = Array.from(selectedTimeIds.value)
  console.log('選択された時間帯のID:', selectedIds)

  router.push(`/rooms/${roomId}/reorder-by-preference`)
}
</script>

<template>
  <div class="flex h-full flex-col p-12">
    <div class="mb-6 shrink-0">
      <h1 class="text-xl text-gray-900">参加可能な時間を選択</h1>
    </div>

    <div class="flex-grow space-y-2 overflow-y-auto">
      <button
        v-for="slot in roomCreationStore.timeOptions"
        :key="slot.id"
        @click="toggle(slot)"
        class="flex w-full items-center justify-between rounded-lg p-4 py-3 text-left transition-colors"
        :class="
          selectedTimeIds.has(slot.id)
            ? 'bg-sky-400 text-white'
            : 'border border-gray-200 bg-white text-gray-900 shadow-sm hover:bg-gray-100'
        "
      >
        <div class="flex flex-col leading-tight">
          <p class="text-base font-light">
            {{ slot.date }}
          </p>
          <p class="text-base font-light">{{ slot.startTime }} - {{ slot.endTime }}</p>
        </div>
        <div
          class="flex h-5 w-5 items-center justify-center rounded-full border"
          :class="selectedTimeIds.has(slot.id) ? 'border-white bg-white' : 'border-gray-400'"
        >
          <div v-if="selectedTimeIds.has(slot.id)" class="h-3 w-3 rounded-full bg-sky-400"></div>
        </div>
      </button>

      <p v-if="!roomCreationStore.hasTimeOptions" class="pt-4 text-center text-gray-500">
        前のページで候補の日時を追加してください。
      </p>
    </div>

    <div class="mt-4 shrink-0">
      <BasicButton
        @click="handleNext"
        text="次へ"
        right-icon="arrow_foward_inv"
        variant="primary"
        size="large"
        class="h-12 w-full"
      />
    </div>
  </div>
</template>
