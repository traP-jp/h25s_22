<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useRoomStore, useVoteStore } from '@/stores'
import BasicButton from '@/components/BasicButton.vue'

const router = useRouter()
const route = useRoute()
const roomStore = useRoomStore()
const voteStore = useVoteStore()

const selectedTimeIds = ref(new Set<string>())

// roomStoreから時間オプションを取得
const timeOptions = computed(() => {
  return roomStore.currentRoomData?.time_options || []
})

// 選択状態を切り替える関数
const toggle = (timeOption: any) => {
  if (selectedTimeIds.value.has(timeOption.id)) {
    selectedTimeIds.value.delete(timeOption.id)
  } else {
    selectedTimeIds.value.add(timeOption.id)
  }

  // voteStoreに保存
  voteStore.setSelectedTimeIds(Array.from(selectedTimeIds.value))
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

  // voteStoreに最終的な選択を保存
  voteStore.setSelectedTimeIds(selectedIds)

  router.push(`/rooms/${roomId}/reorder-by-preference`)
}

// 時間を表示用にフォーマット
const formatTime = (datetime: string) => {
  const date = new Date(datetime)
  return date.toLocaleString('ja-JP', {
    month: 'numeric',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })
}

onMounted(() => {
  // ページ読み込み時にroomDataがない場合は取得
  if (!roomStore.currentRoomData) {
    const roomId = route.params.room_id as string
    if (roomId) {
      roomStore.fetchRoomData(roomId)
    }
  }
})
</script>

<template>
  <div class="flex h-full flex-col p-12">
    <div class="mb-6 shrink-0">
      <h1 class="text-xl text-gray-900">参加可能な時間を選択</h1>
    </div>

    <div class="flex-grow space-y-2 overflow-y-auto">
      <!-- ローディング表示 -->
      <div v-if="roomStore.isLoading" class="flex items-center justify-center py-8">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        <span class="ml-3 text-sm">読み込み中...</span>
      </div>

      <!-- エラー表示 -->
      <div v-else-if="roomStore.error" class="text-center py-8 text-red-500 text-sm">
        {{ roomStore.error }}
      </div>

      <!-- 時間オプション表示 -->
      <button
        v-else
        v-for="timeOption in timeOptions"
        :key="timeOption.id"
        @click="toggle(timeOption)"
        class="flex w-full items-center justify-between rounded-lg p-4 py-3 text-left transition-colors"
        :class="
          selectedTimeIds.has(timeOption.id)
            ? 'bg-sky-400 text-white'
            : 'border border-gray-200 bg-white text-gray-900 shadow-sm hover:bg-gray-100'
        "
      >
        <div class="flex flex-col leading-tight">
          <p class="text-base font-light">
            {{ formatTime(timeOption.start_time) }} - {{ formatTime(timeOption.end_time) }}
          </p>
        </div>
        <div
          class="flex h-5 w-5 items-center justify-center rounded-full border"
          :class="selectedTimeIds.has(timeOption.id) ? 'border-white bg-white' : 'border-gray-400'"
        >
          <div
            v-if="selectedTimeIds.has(timeOption.id)"
            class="h-3 w-3 rounded-full bg-sky-400"
          ></div>
        </div>
      </button>

      <p
        v-if="timeOptions.length === 0 && !roomStore.isLoading && !roomStore.error"
        class="pt-4 text-center text-gray-500"
      >
        時間候補が見つかりません。
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
