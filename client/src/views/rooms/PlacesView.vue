<script setup lang="ts">
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { GoogleMap, Marker } from 'vue3-google-map'

const router = useRouter()
const route = useRoute()

// 地図関連のデータ,本来は取得したものによって設定
const mapCenter = ref({ lat: 35.68, lng: 139.73 })
const mapZoom = ref(12)

const markers = ref([
  { id: 'todai', position: { lat: 35.71267, lng: 139.76195 }, title: '東京大学' },
  { id: 'tokodai', position: { lat: 35.6062, lng: 139.683 }, title: '東京科学大学' },
  { id: 'waseda', position: { lat: 35.70902, lng: 139.71937 }, title: '早稲田大学' },
])

// ダミーのデータ
const places = ref([
  { id: 1, name: '候補地A：東大前のカフェ', address: '東京都文京区本郷', votes: 5, isVotedByMe: false },
  { id: 2, name: '候補地B：早稲田の定食屋', address: '東京都新宿区早稲田', votes: 12, isVotedByMe: true },
  { id: 3, name: '候補地C：東工大近くの公園', address: '東京都目黒区大岡山', votes: 8, isVotedByMe: false },
])

const selectedPlaceId = ref(2)

// --- ボタンのクリックイベント ---
const handleVote = () => {
  const roomId = route.params.room_id
  const selectedPlace = places.value.find((p) => p.id === selectedPlaceId.value)
  alert(`Room ID: ${roomId} の「${selectedPlace?.name}」に投票しました！`)
  // 次のページへ飛ぶ処理
}
</script>

<template>
  <div class="flex w-full items-start justify-center bg-gray-100 p-4">
    <div
      class="flex h-[800px] w-[400px] flex-col overflow-hidden rounded-xl border bg-white shadow-lg"
    >
      <header
        class="flex h-10 w-full shrink-0 items-center justify-end border-b border-gray-200 bg-gray-50 px-4"
      >
        <svg class="h-6 w-6 text-gray-700" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
        </svg>
      </header>

      <main class="flex h-full w-full flex-col">
        <div class="shrink-0 px-12 pt-12">
          <h1 class="mb-4 h-9 text-xl font-normal leading-9 text-gray-900">
            候補地を確認
          </h1>

          <div class="h-72 w-full rounded-lg bg-gray-200">
            <!-- 本来はGoogle Map APIキーを取得して設定 -->
            <GoogleMap
              :api-key="YOUR_Maps_API_KEY"
              :center="mapCenter"
              :zoom="mapZoom"
              style="width: 100%; height: 100%"
              class="rounded-lg"
            >
              <Marker
                v-for="marker in markers"
                :key="marker.id"
                :options="marker"
              />
            </GoogleMap>
          </div>
        </div>

        <div class="mt-8 space-y-4 overflow-y-auto px-12 h-44">
          <div
            v-for="place in places"
            :key="place.id"
            @click="selectedPlaceId = place.id"
            class="flex h-20 w-full cursor-pointer items-center rounded-lg border bg-white p-3 shadow-sm transition-all"
            :class="{
              'border-green-500 shadow-md ring-1 ring-green-500': selectedPlaceId === place.id,
              'border-gray-200': selectedPlaceId !== place.id,
            }"
          >
          <!-- 本来はGoogle Map のサムネイルなど？をとってきて設定する? -->
            <img src="/dummy.png" :alt="place.name" class="h-16 w-16 shrink-0 rounded-md bg-gray-200 object-cover" />
            <div class="ml-3 flex-grow">
              <p class="font-semibold text-gray-800">{{ place.name }}</p>
              <p class="text-xs text-gray-500">{{ place.address }}</p>
              <div class="mt-1 flex items-center">
                <svg class="h-4 w-4 text-yellow-400" viewBox="0 0 20 20" fill="currentColor">
                  <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                </svg>
                <span class="ml-1 text-sm font-medium text-gray-700">{{ place.votes }} 票</span>
              </div>
            </div>
          </div>
        </div>

        <div class="flex-grow"></div>

        <footer class="shrink-0 bg-white px-12 pt-9 pb-20">
          <button
            @click="handleVote"
            class="flex h-12 w-full items-center justify-center gap-2 rounded-lg bg-blue-500 text-base font-semibold text-white transition-colors hover:bg-blue-600"
          >
            <span>投票する</span>
            <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M13.5 4.5L21 12m0 0l-7.5 7.5M21 12H3" />
            </svg>
          </button>
        </footer>
      </main>
    </div>
  </div>
</template>
