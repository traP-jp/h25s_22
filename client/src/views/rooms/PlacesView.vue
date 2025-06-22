<script setup lang="ts">
import { ref } from 'vue'
import { GoogleMap, Marker } from 'vue3-google-map'
import BasicButton from '@/components/BasicButton.vue'

const API_KEY = 'MY_API_KEY' //本来は環境変数などから取得する?
// 地図関連のデータ,本来は取得したものによって設定
const mapCenter = ref({ lat: 35.68, lng: 139.73 })
const mapZoom = ref(12)

const markers = ref([
  { id: 'todai', position: { lat: 35.71267, lng: 139.76195 }, title: '東京大学' },
  { id: 'tokodai', position: { lat: 35.6062, lng: 139.683 }, title: '東京科学大学' },
  { id: 'waseda', position: { lat: 35.70902, lng: 139.71937 }, title: '早稲田大学' },
])
const nextPage = () => {
  // 次のページへの遷移処理
}
// ダミーのデータ
const places = ref([
  {
    id: 1,
    name: '候補地A：東大前のカフェ',
    address: '東京都文京区本郷',
    image: '/dummy.png',
  },
  {
    id: 2,
    name: '候補地B：早稲田の定食屋',
    address: '東京都新宿区早稲田',
    image: '/dummy.png',
  },
  {
    id: 3,
    name: '候補地C：東工大近くの公園',
    address: '東京都目黒区大岡山',
    image: '/dummy.png',
  },
])

const selectedPlaceId = ref(2)
</script>

<template>
  <div class="flex w-full items-start justify-center bg-white">
    <div class="flex h-[800px] w-[400px] flex-col overflow-hidden rounded-xl">
      <main class="flex h-full w-full flex-col">
        <div class="shrink-0 px-12">
          <h1 class="mb-4 h-9 text-xl font-normal leading-9 text-gray-900">候補地を確認</h1>

          <div class="h-72 w-full rounded-lg bg-gray-200">
            <GoogleMap
              :api-key="API_KEY"
              :center="mapCenter"
              :zoom="mapZoom"
              style="width: 100%; height: 100%"
              class="rounded-lg"
            >
              <Marker v-for="marker in markers" :key="marker.id" :options="marker" />
            </GoogleMap>
          </div>
        </div>

        <div class="mt-8 space-y-4 overflow-y-auto px-12 flex-12 min-h-0">
          <div
            v-for="place in places"
            :key="place.id"
            @click="selectedPlaceId = place.id"
            class="flex min-h-20 w-full cursor-pointer items-start rounded-lg border bg-white p-3 shadow-sm transition-all"
            :class="{
              'border-green-500 shadow-md ring-1 ring-green-500': selectedPlaceId === place.id,
              'border-gray-200': selectedPlaceId !== place.id,
            }"
          >
            <img
              :src="place.image || '/dummy.png'"
              :alt="place.name"
              class="h-16 w-16 shrink-0 rounded-md bg-gray-200 object-cover"
            />
            <div class="ml-3 flex-grow">
              <p class="font-semibold text-gray-800">{{ place.name }}</p>
              <p class="text-xs text-gray-500">{{ place.address }}</p>
              <div class="mt-1 flex items-center">
                <svg class="h-4 w-4 text-yellow-400" viewBox="0 0 20 20" fill="currentColor"></svg>
              </div>
            </div>
          </div>
        </div>

        <div class="flex-grow"></div>

        <footer class="shrink-0 bg-white px-12 pt-9 pb-20">
          <BasicButton
            text="投票する →"
            variant="primary"
            size="large"
            right-icon="arrow-right"
            @click="nextPage"
            class="h-12 w-full"
          />
        </footer>
      </main>
    </div>
  </div>
</template>
