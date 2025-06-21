<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

interface Schedule {
  id: number
  date: string
  start: string
  end: string
}
const schedules = ref<Schedule[]>([])

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

  schedules.value.push({
    id: Date.now(), // 一意のIDを生成するために現在のタイムスタンプを使用
    date: selectedDate.value,
    start: startTime.value,
    end: endTime.value,
  })
  selectedDate.value = ''
  startTime.value = ''
  endTime.value = ''
}

const handleNext = () => {
  if (schedules.value.length === 0) {
    alert('スケジュール候補を1つ以上追加してください。')
    return
  }
  router.push('/rooms/edit/place')
}
</script>

<template>
  <div class="flex w-full items-start justify-center py-8">
    <div
      class="flex h-[800px] w-[400px] flex-col overflow-hidden border border-gray-200 bg-white shadow-lg"
    >
      <header
        class="box-border flex h-[38px] w-full shrink-0 items-center justify-end border-b border-gray-200 bg-gray-50 px-5"
      >
        <svg
          class="h-6 w-6"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
          />
        </svg>
      </header>

      <div class="box-border w-full flex-grow overflow-y-auto px-[50px] pt-12">
        <div class="w-full">
          <h1 class="mb-4 h-[36px] text-[20px] font-medium leading-[36px]">日付を入力</h1>

          <div class="box-border mb-3 h-[68px] w-full">
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

          <div class="box-border mb-9 flex h-[69px] gap-3">
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
            <button
              @click="handleAdd"
              class="box-border flex h-12 w-[144px] items-center justify-center rounded-md border border-blue-500 bg-white p-0 text-base font-semibold text-blue-500 transition-all hover:bg-blue-50"
            >
              追加 ＋
            </button>
            <button
              @click="handleNext"
              class="box-border flex h-12 w-[144px] items-center justify-center rounded-md border border-transparent bg-blue-500 p-0 text-base font-semibold text-white transition-all hover:bg-blue-600"
            >
              次へ →
            </button>
          </div>

          <div class="mt-8" v-if="schedules.length > 0">
            <h2 class="mb-3 text-base font-semibold text-gray-600">追加された候補</h2>
            <ul class="m-0 max-h-[150px] list-none overflow-y-auto p-0">
              <li
                v-for="schedule in schedules"
                :key="schedule.id"
                class="mb-2 flex items-center justify-between rounded-md border border-gray-200 bg-gray-50 p-3 px-4 text-sm"
              >
                <span>日付: {{ schedule.date }}</span>
                <span>時間: {{ schedule.start }} 〜 {{ schedule.end }}</span>
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
