<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { GoogleMap, Marker } from 'vue3-google-map'
import BasicButton from '@/components/BasicButton.vue'
import { useRoomCreationStore } from '@/stores'

interface MapPosition {
  lat: number
  lng: number
}
interface MarkerData {
  id: string
  position: MapPosition
  title: string
}

//APIキーを.envから取得する
const API_KEY = import.meta.env.VITE_GOOGLE_MAP_API_KEY
// 地図関連のデータ,本来は取得したものによって設定
const router = useRouter()
const roomCreationStore = useRoomCreationStore()

const placesData = roomCreationStore.suggestedPlaces.map((place) => ({
  id: place.placeID, 
  position: {
    lat: place.location.lat,
    lng: place.location.lng,
  },
  address: place.address,
  title: place.name,
}));
// 開発環境のため、ダミーデータをplacesDataに設定
if (placesData.length === 0) {
  placesData.push({
    id: '1',
    position: { lat: 35.681236, lng: 139.767125
    },
    address: '東京都千代田区丸の内1丁目9-1',
    title: '東京駅',
  })
  placesData.push({
    id: '2',
    position: { lat: 35.689487, lng: 139.691706
    },
    address: '東京都新宿区西新宿1丁目1-1',
    title: '新宿駅',
  })
  placesData.push({
    id: '3',
    position: { lat: 35.658581, lng: 139.745433
    },
    address: '東京都港区六本木6丁目10-1',
    title: '渋谷駅',
  })
}
// 地図の中心座標をplacesDataの最初の要素から取得
const mapCenter = ref<MapPosition>({
  lat: placesData.length > 0 ? placesData[0].position.lat :

  35.681236,
  lng: placesData.length > 0 ? placesData[0].position.lng : 139.767125, 
})
// 地図のズームレベル
const mapZoom = ref(10)

// マーカーのデータをplacesDataから取得
const markers = ref<MarkerData[]>(placesData.map((place) => ({
  id: place.id,
  position: place.position,
  title: place.title,
})))


const nextPage = () => {
  if (roomCreationStore.selectedPlaces.length === 0) {
    alert('候補地を1つ以上選択してください。')
    return
  }
  router.push({ name: 'rooms-edit-decide' }) 
}

</script>

<template>
  <div class="flex w-full items-start justify-center bg-gray-100 p-4">
    <div
      class="flex h-[800px] w-[400px] flex-col overflow-hidden rounded-xl border bg-white shadow-lg"
    >
      <main class="flex h-full w-full flex-col">
        <div class="shrink-0 px-12 pt-12">
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
            v-for="place in placesData"
            :key="place.id"
            @click="selectedPlaceId = place.id"
            class="flex min-h-20 w-full cursor-pointer items-start rounded-lg border bg-white p-3 shadow-sm transition-all"
            :class="{
              'border-green-500 shadow-md ring-1 ring-green-500': selectedPlaceId === place.id,

              'border-gray-200': selectedPlaceId !== place.id,
            }"
          >


            <div class="ml-3 flex-grow">
              <p class="font-semibold text-gray-800">{{ place.title }}</p>

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
