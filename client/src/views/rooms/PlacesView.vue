<script setup lang="ts">
import { ref } from 'vue'
import { GoogleMap, Marker } from 'vue3-google-map'
import BasicButton from '@/components/BasicButton.vue'
import type { PlaceSearchResult } from '@/services/types'

const API_KEY = 'MY_API_KEY' //本来は環境変数などから取得する?
// 地図関連のデータ,本来は取得したものによって設定
const mapCenter = ref({ lat: 35.68, lng: 139.73 })
const mapZoom = ref(12)

const markers = ref([
  { id: 'todai', position: { lat: 35.71267, lng: 139.76195 }, title: '東京大学' },
  { id: 'tokodai', position: { lat: 35.6062, lng: 139.683 }, title: '東京科学大学' },
  { id: 'waseda', position: { lat: 35.70902, lng: 139.71937 }, title: '早稲田大学' },
])

const getPriceRange = (priceLevel: number | undefined): string => {
  if (priceLevel === undefined) return ''

  switch (priceLevel) {
    case 1:
      return '～ ¥1,000'
    case 2:
      return '¥1,000 ～ ¥3,000'
    case 3:
      return '¥3,000 ～ ¥8,000'
    case 4:
      return '¥8,000 以上'
    default:
      return ''
  }
}

const nextPage = () => {
  // 次のページへの遷移処理
}
// PlaceSearchResult型のダミーデータ
const places = ref<PlaceSearchResult[]>([
  {
    name: '候補地A：東大前のカフェ',
    placeID: 'ChIJ1abc123def',
    location: { lat: 35.71267, lng: 139.76195 },
    photoReference: 'AXQCQNS4aDjR_w686vtMDT-_-2a39aS9tFZU-aqdCL9qRy7yujneEi2P',
    priceLevel: 2,
    rating: 4.3,
    address: '東京都文京区本郷7-3-1',
  },
  {
    name: '候補地B：早稲田の定食屋',
    placeID: 'ChIJ2def456ghi',
    location: { lat: 35.70902, lng: 139.71937 },
    photoReference: 'AXQCQNS4aDjR_w686vtMDT-_-2a39aS9tFZU-aqdCL9qRy7yujneEi2P',
    priceLevel: 1,
    rating: 4.1,
    address: '東京都新宿区早稲田鶴巻町513',
  },
  {
    name: '候補地C：東工大近くの公園',
    placeID: 'ChIJ3ghi789jkl',
    location: { lat: 35.6062, lng: 139.683 },
    photoReference: 'AXQCQNS4aDjR_w686vtMDT-_-2a39aS9tFZU-aqdCL9qRy7yujneEi2P',
    priceLevel: 0,
    rating: 3.8,
    address: '東京都目黒区大岡山2-12-1',
  },
])

const selectedPlaceId = ref('ChIJ2def456ghi')
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
            :key="place.placeID"
            @click="selectedPlaceId = place.placeID"
            class="flex min-h-20 w-full cursor-pointer items-start rounded-lg border bg-white p-3 shadow-sm transition-all"
            :class="{
              'border-green-500 shadow-md ring-1 ring-green-500': selectedPlaceId === place.placeID,
              'border-gray-200': selectedPlaceId !== place.placeID,
            }"
          >
            <img
              src="/dummy.png"
              :alt="place.name"
              class="h-16 w-16 shrink-0 rounded-md bg-gray-200 object-cover"
            />
            <div class="ml-3 flex-grow">
              <p class="font-semibold text-gray-800">{{ place.name }}</p>
              <p class="text-xs text-gray-500">{{ place.address }}</p>
              <div class="mt-1 flex items-center space-x-2">
                <!-- 評価を表示 -->
                <div v-if="place.rating" class="flex items-center">
                  <svg class="h-4 w-4 text-yellow-400" viewBox="0 0 20 20" fill="currentColor">
                    <path
                      d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
                    />
                  </svg>
                  <span class="text-sm text-gray-600 ml-1">{{ place.rating }}</span>
                </div>
                <!-- 価格レベルを表示 -->
                <div v-if="place.priceLevel !== undefined" class="flex items-center">
                  <span class="text-xs text-gray-500">{{ getPriceRange(place.priceLevel) }}</span>
                </div>
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
