<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useRoomCreationStore } from '@/stores'
import type { DateTimeFormData } from '@/types'
import BasicButton from '@/components/BasicButton.vue'

const router = useRouter()
const roomCreationStore = useRoomCreationStore()

// フォームの入力値を保持する変数
const selectedDate = ref('')
const startTime = ref('')
const endTime = ref('')

// ボタンのクリックイベント用の関数
const handleAdd = () => {
  if (!selectedDate.value || !startTime.value || !endTime.value) {
    alert('すべての日付と時刻を入力してください。')
    return
  }

  const formData: DateTimeFormData = {
    date: selectedDate.value,
    startTime: startTime.value,
    endTime: endTime.value,
  }

  // ストアに追加
  roomCreationStore.addTimeOption(formData)

  // フォームをクリア
  selectedDate.value = ''
  startTime.value = ''
  endTime.value = ''
}

const handleNext = () => {
  if (!roomCreationStore.hasTimeOptions) {
    alert('スケジュール候補を1つ以上追加してください。')
    return
  }
  router.push('/rooms/edit/place')
}

const removeSchedule = (timeOptionId: string) => {
  roomCreationStore.removeTimeOption(timeOptionId)
}
</script>

<template>
  <div class="flex w-full items-start justify-center py-8">
    <div
      class="flex h-[800px] w-[400px] flex-col overflow-hidden border border-gray-200 bg-white shadow-lg"
    >

      <div class="box-border w-full flex-grow overflow-y-auto px-12 pt-12">
        <div class="w-full">
          <h1 class="mb-4 h-9 text-xl font-medium leading-9">日付を入力</h1>

          <div class="mb-3 w-full">
            <label for="date-input" class="mb-2 block h-5 text-sm font-medium leading-5"
              >日付</label
            >
            <input
              type="date"
              id="date-input"
              v-model="selectedDate"
              required
              class="box-border h-10 w-full rounded-md border border-gray-300 px-3 text-base"
            />
          </div>

          <div class="mb-9 flex gap-3">
            <div class="w-full">
              <label for="start-time" class="mb-2 block h-5 text-sm font-medium leading-5"
                >開始時刻</label
              >
              <input
                type="time"
                id="start-time"
                v-model="startTime"
                required
                class="box-border h-10 w-full rounded-md border border-gray-300 px-3 text-base"
              />
            </div>
            <div class="w-full">
              <label for="end-time" class="mb-2 block h-5 text-sm font-medium leading-5"
                >終了時刻</label
              >
              <input
                type="time"
                id="end-time"
                v-model="endTime"
                required
                class="box-border h-10 w-full rounded-md border border-gray-300 px-3 text-base"
              />
            </div>
          </div>

          <div class="flex h-12 gap-3">
            <BasicButton
              text="追加"
              rightIcon="plus"
              variant="secondary"
              size="medium"
              @click="handleAdd"
            />
            <BasicButton
              text="次へ"
              rightIcon="arrow_foward_inv"
              variant="primary"
              size="medium"
              @click="handleNext"
            />
          </div>

          <div class="mt-8" v-if="roomCreationStore.timeOptions.length > 0">
            <h2 class="mb-3 text-base font-semibold text-gray-600">追加された候補</h2>
            <ul class="m-0 max-h-[150px] list-none overflow-y-auto p-0">
              <li
                v-for="timeOption in roomCreationStore.timeOptions"
                :key="timeOption.date + timeOption.startTime"
                class="mb-2 flex items-center justify-between rounded-md border border-gray-200 bg-gray-50 p-3 px-4 text-sm"
              >
                <span>日付: {{ timeOption.date }}</span>
                <span>時間: {{ timeOption.startTime }} 〜 {{ timeOption.endTime }}</span>
                <button
                  @click="removeSchedule(timeOption.date + timeOption.startTime)"
                  class="ml-2 text-red-500 transition-all hover:text-red-700"
                  aria-label="Remove schedule"
                >
                  <svg
                    class="h-5 w-5"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                  >
                    <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
input[type='date'],
input[type='time'] {
  color: transparent;
}

input[type='date']:focus,
input[type='time']:focus,
input[type='date']:valid,
input[type='time']:valid {
  color: #374151;
}
</style>
