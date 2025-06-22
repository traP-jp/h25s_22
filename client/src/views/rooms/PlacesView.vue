<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { GoogleMap, Marker } from 'vue3-google-map'
import BasicButton from '@/components/BasicButton.vue'
import type { PlaceSearchResult } from '@/services/types'
import { useRoomStore } from '@/stores/room'

const route = useRoute()
const router = useRouter()
const roomStore = useRoomStore()

const API_KEY = import.meta.env.VITE_GOOGLE_MAP_API_KEY || 'MY_API_KEY'

// 地図関連のデータ
const mapCenter = computed(() => {
  if (
    !roomStore.currentRoomData?.place_options ||
    roomStore.currentRoomData.place_options.length === 0
  ) {
    return { lat: 35.68, lng: 139.73 }
  }

  const places = roomStore.currentRoomData.place_options
  const avgLat = places.reduce((sum, place) => sum + place.location.lat, 0) / places.length
  const avgLng = places.reduce((sum, place) => sum + place.location.lng, 0) / places.length

  return { lat: avgLat, lng: avgLng }
})

const mapZoom = ref(12)

// 候補場所からマーカーを生成
const markers = computed(() => {
  if (!roomStore.currentRoomData?.place_options) return []

  return roomStore.currentRoomData.place_options.map((place) => ({
    id: place.id,
    position: place.location,
    title: place.name,
  }))
})

// PlaceSearchResult形式に変換した候補地
const places = computed<PlaceSearchResult[]>(() => {
  if (!roomStore.currentRoomData?.place_options) return []

  return roomStore.currentRoomData.place_options.map(
    (place): PlaceSearchResult => ({
      name: place.name,
      placeID: place.google_place_id,
      location: place.location,
      photoReference: place.photoReference,
      priceLevel: place.priceLevel,
      rating: place.rating,
      address: place.address,
    }),
  )
})

const selectedPlaceId = ref('')

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

// ルームデータを取得
const fetchRoomData = async () => {
  const roomId = route.params.room_id as string
  console.log('PlacesView - ルートパラメータ:', route.params)
  console.log('PlacesView - roomId:', roomId)

  if (!roomId) {
    roomStore.error = 'ルームIDが指定されていません'
    return
  }

  try {
    const data = await roomStore.fetchRoomData(roomId)

    // デフォルトで最初の場所を選択
    if (data.place_options && data.place_options.length > 0) {
      selectedPlaceId.value = data.place_options[0].google_place_id
    }
  } catch (err) {
    console.error('ルーム情報の取得に失敗:', err)
    // エラーはstoreで管理されているので、ここでは何もしない
  }
}

const nextPage = () => {
  // 次のページ（投票画面）への遷移処理
  const roomId = route.params.room_id as string
  router.push(`/rooms/${roomId}/reorder-by-preference`)
}

onMounted(() => {
  fetchRoomData()
})
</script>

<template>
  <div class="flex w-full items-start justify-center bg-white">
    <div class="flex h-[800px] w-[400px] flex-col overflow-hidden rounded-xl">
      <main class="flex h-full w-full flex-col">
        <div class="shrink-0 px-12">
          <h1 class="mb-4 h-9 text-xl font-normal leading-9 text-gray-900">候補地を確認</h1>

          <!-- ルーム名表示 -->
          <div v-if="roomStore.currentRoomData" class="mb-4">
            <h2 class="text-sm font-medium text-gray-700 mb-2">ルーム名</h2>
            <p class="text-gray-900 font-semibold">{{ roomStore.currentRoomData.name }}</p>
          </div>

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

        <div class="mt-8 space-y-4 overflow-y-auto px-12 flex-5 min-h-0">
          <!-- ローディング表示 -->
          <div v-if="roomStore.isLoading" class="flex items-center justify-center py-8">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
            <span class="ml-3 text-sm">読み込み中...</span>
          </div>

          <!-- エラー表示 -->
          <div v-else-if="roomStore.error" class="text-center py-8 text-red-500 text-sm">
            {{ roomStore.error }}
          </div>

          <!-- 候補地が見つからない場合 -->
          <div v-else-if="places.length === 0" class="text-center py-8 text-gray-500 text-sm">
            候補地が見つかりませんでした。
          </div>

          <!-- 候補地リスト -->
          <div v-else>
            <div
              v-for="place in places"
              :key="place.placeID"
              @click="selectedPlaceId = place.placeID"
              class="flex min-h-20 w-full cursor-pointer items-start rounded-lg border bg-white p-3 shadow-sm transition-all"
              :class="{
                'border-green-500 shadow-md ring-1 ring-green-500':
                  selectedPlaceId === place.placeID,
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
        </div>

        <div class="flex-grow"></div>

        <footer class="shrink-0 bg-white px-12 pt-9 pb-20">
          <BasicButton
            :text="roomStore.isLoading ? '読み込み中...' : '投票する'"
            variant="primary"
            size="large"
            right-icon="arrow-right"
            @click="nextPage"
            class="h-12 w-full"
            :disabled="roomStore.isLoading || roomStore.error !== null || places.length === 0"
          />
        </footer>
      </main>
    </div>
  </div>
</template>
