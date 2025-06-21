<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="bg-white rounded-lg shadow p-6">
        <h1 class="text-2xl font-bold text-gray-900 mb-6">内容確認・決定</h1>

        <!-- 日時情報の確認 -->
        <div class="mb-8">
          <h2 class="text-lg font-semibold text-gray-800 mb-4">選択された日時</h2>
          <div class="space-y-2">
            <div
              v-for="timeOption in roomCreationStore.timeOptions"
              :key="timeOption.id"
              class="bg-blue-50 p-3 rounded-md border border-blue-200"
            >
              <span class="font-medium">{{ timeOption.date }}</span>
              <span class="ml-4 text-gray-600"
                >{{ timeOption.startTime }} 〜 {{ timeOption.endTime }}</span
              >
            </div>
          </div>
        </div>

        <!-- 場所検索結果 -->
        <div class="mb-8">
          <div class="flex items-center justify-between mb-4">
            <h2 class="text-lg font-semibold text-gray-800">候補場所</h2>
            <button
              @click="handleRetrySearch"
              class="px-4 py-2 text-sm text-blue-600 border border-blue-600 rounded-md hover:bg-blue-50 transition-colors"
              :disabled="roomCreationStore.isLoading"
            >
              再検索
            </button>
          </div>

          <div v-if="roomCreationStore.isLoading" class="flex items-center justify-center py-8">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
            <span class="ml-3">検索中...</span>
          </div>

          <div
            v-else-if="roomCreationStore.suggestedPlaces.length === 0"
            class="text-center py-8 text-gray-500"
          >
            候補場所が見つかりませんでした。
            <button @click="goBackToPlace" class="text-blue-600 underline ml-2">
              検索条件を変更する
            </button>
          </div>

          <div v-else class="grid gap-4 md:grid-cols-2">
            <div
              v-for="place in roomCreationStore.suggestedPlaces"
              :key="place.id"
              class="border rounded-lg p-4 hover:shadow-md transition-shadow cursor-pointer"
              :class="{
                'border-blue-500 bg-blue-50': roomCreationStore.selectedPlaces.find(
                  (p: any) => p.id === place.id,
                ),
                'border-gray-200': !roomCreationStore.selectedPlaces.find(
                  (p: any) => p.id === place.id,
                ),
              }"
              @click="handlePlaceToggle(place)"
            >
              <h3 class="font-semibold text-gray-900">{{ place.name }}</h3>
              <p class="text-sm text-gray-600 mt-1">{{ place.address }}</p>
              <div v-if="place.rating" class="flex items-center mt-2">
                <span class="text-yellow-500">★</span>
                <span class="text-sm text-gray-600 ml-1">{{ place.rating }}</span>
              </div>
              <div v-if="place.types && place.types.length > 0" class="mt-2">
                <span
                  v-for="type in place.types.slice(0, 2)"
                  :key="type"
                  class="inline-block bg-gray-100 text-xs px-2 py-1 rounded-full mr-1"
                >
                  {{ type }}
                </span>
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

        <!-- ルーム作成フォーム -->
        <div class="mb-8">
          <h2 class="text-lg font-semibold text-gray-800 mb-4">ルーム情報</h2>
          <div class="space-y-4">
            <div>
              <label for="room-name" class="block text-sm font-medium text-gray-700 mb-2">
                ルーム名 <span class="text-red-500">*</span>
              </label>
              <input
                id="room-name"
                v-model="roomName"
                type="text"
                required
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="ルーム名を入力してください"
              />
            </div>
            <div>
              <label for="room-description" class="block text-sm font-medium text-gray-700 mb-2">
                説明（任意）
              </label>
              <textarea
                id="room-description"
                v-model="roomDescription"
                rows="3"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="ルームの説明を入力してください"
              ></textarea>
            </div>
          </div>
        </div>

        <!-- アクションボタン -->
        <div class="flex gap-4 justify-end">
          <button
            @click="goBackToPlace"
            class="px-6 py-2 text-gray-600 border border-gray-300 rounded-md hover:bg-gray-50 transition-colors"
            :disabled="isCreating"
          >
            戻る
          </button>
          <button
            @click="handleCreateRoom"
            class="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            :disabled="
              isCreating || !roomName.trim() || roomCreationStore.suggestedPlaces.length === 0
            "
          >
            <span v-if="isCreating">作成中...</span>
            <span v-else>ルームを作成</span>
          </button>
        </div>

        <!-- エラー表示 -->
        <div
          v-if="roomCreationStore.error"
          class="mt-4 p-3 bg-red-50 border border-red-200 rounded-md"
        >
          <p class="text-red-800 text-sm">{{ roomCreationStore.error }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useRoomCreationStore } from '@/stores'

const router = useRouter()
const roomCreationStore = useRoomCreationStore()

const roomName = ref('')
const roomDescription = ref('')
const isCreating = ref(false)

onMounted(() => {
  // 必要なデータが揃っていない場合は前のページに戻す
  if (!roomCreationStore.canProceedToDecide) {
    alert('必要な情報が不足しています。最初からやり直してください。')
    router.push('/rooms/edit/date-and-time')
  }
})

const handlePlaceToggle = (place: any) => {
  roomCreationStore.togglePlaceSelection(place)
}

const handleRetrySearch = async () => {
  try {
    await roomCreationStore.searchPlaces()
  } catch (error) {
    alert('場所検索でエラーが発生しました。')
  }
}

const handleCreateRoom = async () => {
  if (!roomName.value.trim()) {
    alert('ルーム名を入力してください。')
    return
  }

  isCreating.value = true

  try {
    const result = await roomCreationStore.createNewRoom(roomName.value, roomDescription.value)

    if (result && result.success) {
      alert('ルームが作成されました！')
      router.push('/rooms/edit/created')
    } else {
      alert('ルームの作成に失敗しました。')
    }
  } catch (error) {
    alert('ルーム作成でエラーが発生しました。')
  } finally {
    isCreating.value = false
  }
}

const goBackToPlace = () => {
  router.push('/rooms/edit/place')
}
</script>
