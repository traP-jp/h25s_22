<template>
  <div class="flex w-full items-start justify-center bg-white">
    <div class="flex h-screen w-[400px] max-w-md flex-col overflow-hidden rounded-xl">
      <main class="flex h-full w-full flex-col">
        <div class="shrink-0 px-12">
          <h1 class="mb-4 h-9 text-xl font-normal leading-9 text-gray-900">内容確認・決定</h1>

          <!-- Google Map -->
          <div class="h-72 w-full rounded-lg bg-gray-200">
            <GoogleMap
              :api-key="API_KEY"
              :center="mapCenter"
              :zoom="mapZoom"
              style="width: 100%; height: 100%"
              class="rounded-lg"
            >
              <Marker
                v-for="marker in markers"
                :key="marker.id"
                :options="{ position: marker.position, title: marker.title }"
                @click="() => onMarkerClick(marker.id)"
              />
            </GoogleMap>
          </div>
        </div>

        <!-- 候補場所リスト -->
        <div class="mt-8 space-y-4 overflow-y-auto px-12 flex-1 min-h-0">
          <div class="flex items-center justify-between mb-4">
            <h2 class="text-sm font-medium text-gray-700">候補場所</h2>
            <button
              @click="handleRetrySearch"
              class="px-2 py-1 text-xs text-blue-600 border border-blue-600 rounded-md hover:bg-blue-50 transition-colors"
              :disabled="roomCreationStore.isLoading"
            >
              再検索
            </button>
          </div>

          <div v-if="roomCreationStore.isLoading" class="flex items-center justify-center py-8">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
            <span class="ml-3 text-sm">検索中...</span>
          </div>

          <div
            v-else-if="roomCreationStore.suggestedPlaces.length === 0"
            class="text-center py-8 text-gray-500 text-sm"
          >
            候補場所が見つかりませんでした。
            <button @click="goBackToPlace" class="text-blue-600 underline ml-2">
              検索条件を変更する
            </button>
          </div>

          <div v-else class="space-y-4">
            <div
              v-for="place in roomCreationStore.suggestedPlaces"
              :key="place.placeID"
              @click="handlePlaceToggle(place)"
              class="flex min-h-20 w-full cursor-pointer items-start rounded-lg border bg-white p-3 shadow-sm transition-all"
              :class="{
                'border-blue-500 shadow-md ring-1 ring-blue-500':
                  roomCreationStore.selectedPlaces.find((p: any) => p.placeID === place.placeID),
                'border-gray-200': !roomCreationStore.selectedPlaces.find(
                  (p: any) => p.placeID === place.placeID,
                ),
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

          <div
            v-if="roomCreationStore.selectedPlaces.length > 0"
            class="mt-4 p-3 bg-green-50 rounded-md"
          >
            <p class="text-sm text-green-800">
              {{ roomCreationStore.selectedPlaces.length }}個の場所を選択しています。
              選択しない場合は、すべての候補が使用されます。
            </p>
          </div>
        </div>

        <!-- ルーム作成ボタン -->
        <div class="shrink-0 px-12 py-4 bg-white border-t">
          <div class="mb-4">
            <h2 class="text-sm font-medium text-gray-700 mb-2">ルーム名</h2>
            <p class="text-gray-900 font-semibold">{{ roomCreationStore.roomTitle }}</p>
          </div>

          <!-- アクションボタン -->
          <div class="flex gap-2">
            <button
              @click="goBackToPlace"
              class="flex-1 px-4 py-2 text-sm text-gray-600 border border-gray-300 rounded-md hover:bg-gray-50 transition-colors"
              :disabled="isCreating"
            >
              戻る
            </button>
            <button
              @click="handleCreateRoom"
              class="flex-1 px-4 py-2 text-sm bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              :disabled="isCreating || roomCreationStore.suggestedPlaces.length === 0"
            >
              <span v-if="isCreating">作成中...</span>
              <span v-else>ルームを作成</span>
            </button>
          </div>

          <!-- エラー表示 -->
          <div
            v-if="roomCreationStore.error"
            class="mt-2 p-2 bg-red-50 border border-red-200 rounded-md"
          >
            <p class="text-red-800 text-xs">{{ roomCreationStore.error }}</p>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { GoogleMap, Marker } from 'vue3-google-map'
import { useRoomCreationStore } from '@/stores'
import type { PlaceSearchResult } from '@/services/types'

const router = useRouter()
const roomCreationStore = useRoomCreationStore()

const isCreating = ref(false)

// Google Maps設定
const API_KEY =
  import.meta.env.VITE_GOOGLE_MAPS_API_KEY || 'AIzaSyBCTIIJoLo5Qj5pWNGFQ-UwP5F7Bq1qhYo'

// 地図の中心とズームレベルを候補場所から計算
const mapCenter = computed(() => {
  if (roomCreationStore.suggestedPlaces.length === 0) {
    return { lat: 35.68, lng: 139.73 } // デフォルト（東京）
  }

  const places = roomCreationStore.suggestedPlaces
  const avgLat = places.reduce((sum, place) => sum + place.location.lat, 0) / places.length
  const avgLng = places.reduce((sum, place) => sum + place.location.lng, 0) / places.length

  return { lat: avgLat, lng: avgLng }
})

const mapZoom = ref(12)

// 候補場所からマーカーを生成
const markers = computed(() => {
  return roomCreationStore.suggestedPlaces.map((place) => ({
    id: place.placeID,
    position: place.location,
    title: place.name,
  }))
})

// 価格レベルを金額範囲に変換
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

// マーカークリック時の処理
const onMarkerClick = (placeID: string) => {
  const place = roomCreationStore.suggestedPlaces.find((p) => p.placeID === placeID)
  if (place) {
    handlePlaceToggle(place)
  }
}

onMounted(() => {
  if (!roomCreationStore.canProceedToDecide) {
    alert('必要な情報が不足しています。最初からやり直してください。')
    router.push('/rooms/edit/date-and-time')
  }
})

const handlePlaceToggle = (place: PlaceSearchResult) => {
  roomCreationStore.togglePlaceSelection(place)
}

const handleRetrySearch = async () => {
  try {
    await roomCreationStore.searchPlaces()
  } catch (error) {
    alert('場所検索でエラーが発生しました。')
    console.error('場所検索エラー:', error)
  }
}

const handleCreateRoom = async () => {
  if (!roomCreationStore.roomTitle.trim()) {
    alert('ルーム名が設定されていません。')
    return
  }

  isCreating.value = true

  try {
    const result = await roomCreationStore.createNewRoom()

    if (result && result.success) {
      alert('ルームが作成されました！')
      router.push('/rooms/edit/created')
    } else {
      alert('ルームの作成に失敗しました。')
    }
  } catch (error) {
    alert('ルーム作成でエラーが発生しました。')
    console.error('ルーム作成エラー:', error)
  } finally {
    isCreating.value = false
  }
}

const goBackToPlace = () => {
  router.push('/rooms/edit/place')
}
</script>
