<script setup lang="ts">
import { ref } from 'vue'

// --- ダミーデータ ---
const eventDetails = ref({
  title: '渋谷でランチ会',
  date: '2025年6月26日',
  participants: 4,
  responded: 4,
})

const timeSlots = ref([
  {
    id: 1,
    timeRange: '6/26 12:00-13:00',
    voters: ['田中さん', '佐藤さん', '山田さん'],
    voteCount: 3,
    percentage: 75,
  },
  {
    id: 2,
    timeRange: '6/26 13:00-14:00',
    voters: ['鈴木さん', '高橋さん'],
    voteCount: 2,
    percentage: 50,
  },
  { id: 3, timeRange: '6/27 14:00-15:00', voters: ['伊藤さん'], voteCount: 1, percentage: 25 },
  { id: 4, timeRange: '6/27 15:00-16:00', voters: ['渡辺さん'], voteCount: 1, percentage: 25 },
  { id: 5, timeRange: '6/27 16:00-17:00', voters: ['山本さん'], voteCount: 1, percentage: 25 },
])

const placeRankings = ref([
  { rank: 1, name: 'Italian Kitchen VANSAN 渋谷店' },
  { rank: 2, name: '焼肉ホルモン 炭火や' },
  { rank: 3, name: 'カフェ＆ダイニング BLUE' },
  { rank: 4, name: '寿司 まさる' },
  { rank: 5, name: '中華料理 泰興楼' },
  { rank: 6, name: 'とんかつ まい泉' },
  { rank: 7, name: 'ラーメン 一蘭 渋谷店' },
  { rank: 8, name: '居酒屋 鳥貴族 渋谷店' },
  { rank: 9, name: 'カレーの市民 アルバ 渋谷店' },
  { rank: 10, name: 'ベーカリーカフェ パンの田島 渋谷店' },
  { rank: 11, name: 'フレンチレストラン Le Pain Quotidien 渋谷店' },
  { rank: 12, name: 'ベトナム料理 フォーガー 渋谷店' },
  { rank: 13, name: '韓国料理 チーズタッカルビ 渋谷店' },
  { rank: 14, name: 'タイ料理 サイアム 渋谷店' },
  { rank: 15, name: 'インド料理 タージマハール 渋谷店' },
  { rank: 16, name: 'メキシコ料理 タコス 渋谷店' },
  { rank: 17, name: 'アメリカンダイナー ハンバーガー 渋谷店' },
  { rank: 18, name: 'スペイン料理 パエリア 渋谷店' },
  { rank: 19, name: 'ギリシャ料理 ムサカ 渋谷店' },
  { rank: 20, name: 'ロシア料理 ボルシチ 渋谷店' },
])

// カードの開閉状態
const isTimeCardOpen = ref(true)
const isPlaceCardOpen = ref(true)
</script>

<template>
  <div class="flex w-full items-start justify-center bg-gray-100 p-4">
    <div
      class="flex h-[800px] w-[400px] flex-col overflow-hidden rounded-xl border bg-gray-50 shadow-lg"
    >
      <main class="flex-grow p-6">
        <div class="mb-6">
          <h1 class="text-2xl font-bold text-gray-800">{{ eventDetails.title }}</h1>
          <div class="mt-2 flex items-center space-x-4 text-sm text-gray-500">
            <div class="flex items-center">
              <svg
                class="mr-1.5 h-4 w-4"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M5.75 3a.75.75 0 01.75.75v.5h7V3.75a.75.75 0 011.5 0v.5h.25a2.75 2.75 0 012.75 2.75v9.5A2.75 2.75 0 0115.25 19h-10.5A2.75 2.75 0 012 16.5v-9.5A2.75 2.75 0 014.75 4.5h.25v-.5a.75.75 0 01.75-.75zM3.5 8.75v7.75c0 .69.56 1.25 1.25 1.25h10.5c.69 0 1.25-.56 1.25-1.25V8.75h-13z"
                  clip-rule="evenodd"
                />
              </svg>
            </div>
            <div class="flex items-center">
              <svg
                class="mr-1.5 h-4 w-4"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  d="M7 8a3 3 0 100-6 3 3 0 000 6zM14.5 9a.75.75 0 01.75.75v.008a.75.75 0 01-.75.75h-4.5a.75.75 0 01-.75-.75V9.75A.75.75 0 0110 9h4.5zM8 12a2 2 0 100-4 2 2 0 000 4zM14.5 13.5a.75.75 0 01.75.75v.008a.75.75 0 01-.75.75h-4.5a.75.75 0 01-.75-.75V14.25a.75.75 0 01.75-.75h4.5z"
                />
                <path
                  fill-rule="evenodd"
                  d="M.99 10.25a.75.75 0 01.75-.75h16.5a.75.75 0 010 1.5H1.74a.75.75 0 01-.75-.75zM2.5 5.75A.75.75 0 013.25 5h13.5a.75.75 0 010 1.5H3.25A.75.75 0 012.5 5.75z"
                  clip-rule="evenodd"
                />
              </svg>
              <span>{{ eventDetails.participants }}人回答済み</span>
            </div>
          </div>
        </div>

        <div class="mb-6 rounded-xl border border-gray-200 bg-white">
          <div
            class="flex cursor-pointer items-center justify-between p-4"
            @click="isTimeCardOpen = !isTimeCardOpen"
          >
            <div class="flex items-center">
              <svg
                class="h-6 w-6 text-gray-500"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z"
                />
              </svg>
              <h2 class="ml-3 font-bold text-gray-800">時間帯の投票結果</h2>
            </div>
            <svg
              class="h-5 w-5 text-gray-500 transition-transform"
              :class="{ 'rotate-180': !isTimeCardOpen }"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M14.77 12.78a.75.75 0 01-1.06 0L10 9.06l-3.72 3.72a.75.75 0 11-1.06-1.06l4.25-4.25a.75.75 0 011.06 0l4.25 4.25a.75.75 0 010 1.06z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
          <div
            class="overflow-hidden transition-all duration-300 ease-in-out"
            :class="isTimeCardOpen ? 'max-h-[500px]' : 'max-h-0'"
          >
            <div class="border-t border-gray-200 p-4">
              <div class="max-h-48 space-y-4 overflow-y-auto">
                <div v-for="slot in timeSlots" :key="slot.id">
                  <div class="flex items-center justify-between text-sm">
                    <span class="font-semibold text-gray-700">{{ slot.timeRange }}</span>
                    <span class="font-medium text-gray-600"
                      >{{ slot.voteCount }}票 ({{ slot.percentage }}%)</span
                    >
                  </div>
                  <div class="mt-2 h-2 w-full rounded-full bg-gray-200">
                    <div
                      class="h-2 rounded-full bg-blue-500"
                      :style="{ width: slot.percentage + '%' }"
                    ></div>
                  </div>
                  <div class="mt-2 flex flex-wrap gap-1.5">
                    <span
                      v-for="voter in slot.voters"
                      :key="voter"
                      class="rounded-full bg-blue-100 px-2 py-0.5 text-xs font-medium text-blue-800"
                      >{{ voter }}</span
                    >
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="rounded-xl border border-gray-200 bg-white">
          <div
            class="flex cursor-pointer items-center justify-between p-4"
            @click="isPlaceCardOpen = !isPlaceCardOpen"
          >
            <div class="flex items-center">
              <svg
                class="h-6 w-6 text-gray-500"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M15 10.5a3 3 0 11-6 0 3 3 0 016 0z"
                />
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1115 0z"
                />
              </svg>
              <h2 class="ml-3 font-bold text-gray-800">店舗の人気ランキング</h2>
            </div>
            <svg
              class="h-5 w-5 text-gray-500 transition-transform"
              :class="{ 'rotate-180': !isPlaceCardOpen }"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M14.77 12.78a.75.75 0 01-1.06 0L10 9.06l-3.72 3.72a.75.75 0 11-1.06-1.06l4.25-4.25a.75.75 0 011.06 0l4.25 4.25a.75.75 0 010 1.06z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
          <div
            class="overflow-hidden transition-all duration-300 ease-in-out"
            :class="isPlaceCardOpen ? 'max-h-[500px]' : 'max-h-0'"
          >
            <div class="border-t border-gray-200 p-4">
              <div class="max-h-40 space-y-3 overflow-y-auto">
                <div v-for="place in placeRankings" :key="place.rank" class="flex items-center">
                  <div
                    class="flex h-7 w-7 shrink-0 items-center justify-center rounded-full text-sm font-bold text-white"
                    :class="{
                      'bg-yellow-500': place.rank === 1,
                      'bg-gray-400': place.rank === 2,
                      'bg-yellow-700': place.rank === 3,
                      'bg-gray-300': place.rank > 3,
                    }"
                  >
                    {{ place.rank }}
                  </div>
                  <span class="ml-4 font-medium text-gray-700">{{ place.name }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>
