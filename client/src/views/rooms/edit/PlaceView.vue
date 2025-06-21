<script setup lang="ts">
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { GoogleMap, Marker } from 'vue3-google-map'

// 作成したカスタムコンポーネントをインポート
import BasicButton from '@/components/BasicButton.vue'
import BasicInput from '@/components/BasicInput.vue'

const router = useRouter()
const route = useRoute()

// --- 地図関連のデータ ---
const mapCenter = ref({ lat: 35.68, lng: 139.73 }) // 中心
const mapZoom = ref(12) // 縮尺

// 表示する3つの大学のピン情報
const markers = ref([
  { id: 'todai', position: { lat: 35.71267, lng: 139.76195 }, title: '東京大学' },
  { id: 'tokodai', position: { lat: 35.6062, lng: 139.683 }, title: '東京科学大学' },
  { id: 'waseda', position: { lat: 35.70902, lng: 139.71937 }, title: '早稲田大学' },
])

// --- 候補地リストのダミーデータ ---
const places = ref([
  {
    id: 1,
    name: '候補地A：東大前のカフェ',
    votes: 5,
    isVoted: false,
  },
  {
    id: 2,
    name: '候補地B：早稲田の定食屋',
    votes: 12,
    isVoted: true,
  },
  {
    id: 3,
    name: '候補地C：東工大近くの公園',
    votes: 8,
    isVoted: false,
  },
])
const selectedPlaceId = ref(2) // ID:2の場所を選択状態に

// --- 検索バーの入力値 ---
const searchQuery = ref('')

// --- ボタンのクリックイベント ---
const handleConfirm = () => {
  const roomId = route.params.room_id
  const selectedPlace = places.value.find((p) => p.id === selectedPlaceId.value)
  alert(`Room ID: ${roomId} の場所を「${selectedPlace?.name}」に決定します！`)
}
</script>

<template>
  <div class="flex w-full items-start justify-center bg-gray-100 py-8">
    <div
      class="flex h-[800px] w-[400px] flex-col overflow-hidden rounded-xl border bg-white shadow-lg"
    >
      <header
        class="box-border flex h-9 w-full shrink-0 items-center justify-end border-b border-gray-200 bg-gray-50 px-5"
      >
        <svg class="h-6 w-6 text-gray-800" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
        </svg>
      </header>

      <div class="flex flex-grow flex-col overflow-y-auto">
        <div class="shrink-0 px-12 pt-12">
          <h1 class="mb-4 h-9 text-2xl font-bold text-gray-900">
            候補地を選択
          </h1>

          <BasicInput
            v-model="searchQuery"
            label="場所を検索"
            size="large"
            placeholder="例: 東京大学"
            left-icon="search"
          />

          <div class="mt-4 h-72 w-full rounded-lg bg-gray-200">
            <GoogleMap
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

        <div class="mt-8 flex-grow space-y-4 overflow-y-auto px-12 pb-4">
          <div
            v-for="place in places"
            :key="place.id"
            @click="selectedPlaceId = place.id"
            class="flex h-20 w-full cursor-pointer items-center rounded-lg border bg-white p-3 shadow-sm transition-all"
            :class="{
              'border-green-400 shadow-md ring-2 ring-green-200':
                selectedPlaceId === place.id,
              'border-gray-200': selectedPlaceId !== place.id,
            }"
          >
            <img
              src="/dummy.png"
              :alt="place.name"
              class="h-full w-20 shrink-0 rounded-md bg-gray-200 object-cover"
            />
            <div class="ml-4 flex-grow">
              <p class="font-semibold text-gray-800">{{ place.name }}</p>
              <p class="text-sm text-gray-500">
                <span class="font-bold">{{ place.votes }} 票</span>
              </p>
            </div>
          </div>
        </div>

        <footer class="mt-9 shrink-0 bg-white p-6 pt-0">
          <BasicButton
            text="次へ"
            variant="primary"
            size="large"
            right-icon="arrow-right"
            @click="handleNext"
            class="h-12 w-full"
          />
        </footer>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* スタイルはTailwind CSSとカスタムコンポーネントで管理 */
</style>