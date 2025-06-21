<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useRoomCreationStore } from '@/stores'
import BasicButton from '@/components/BasicButton.vue'
import BasicInput from '@/components/BasicInput.vue'

const router = useRouter()
const roomCreationStore = useRoomCreationStore()

// フォームの状態
const genre = ref('')
const standardPlace = ref('')
const useCurrentLocation = ref(false)
const distance = ref(1000)
const candidatesNumber = ref(5)

const handleLocationTypeChange = (value: string) => {
  useCurrentLocation.value = value === 'currentPlace'
  if (useCurrentLocation.value) {
    standardPlace.value = ''
  }
}

const goNext = async () => {
  // 設定をストアに保存
  roomCreationStore.updatePlaceSearchSettings({
    genre: genre.value,
    baseLocation: standardPlace.value,
    useCurrentLocation: useCurrentLocation.value,
    radius: distance.value,
    maxCandidates: candidatesNumber.value,
  })

  // 場所検索を実行
  try {
    await roomCreationStore.searchPlaces()

    if (roomCreationStore.suggestedPlaces.length > 0) {
      router.push('/rooms/edit/decide')
    } else {
      alert('該当する場所が見つかりませんでした。条件を変更してください。')
    }
  } catch (error) {
    alert('場所検索でエラーが発生しました。')
    console.error('場所検索エラー:', error)
  }
}
</script>

<template>
  <div class="flex flex-col gap-4 justify-center items-center">
    <h1 class="header w-75 size-9">場所を設定</h1>
    <div class="w-75 flex flex-col gap-9">
      <div class="w-75 flex flex-col gap-3">
        <div class="w-75 size-18 flex flex-col gap-2">
          <div>ジャンル</div>
          <select v-model="genre" name="genre" class="inputform size-10 w-75">
            <option value="">選択してください</option>
            <option value="restaurant">レストラン</option>
            <option value="cafe">カフェ</option>
            <option value="bar">バー</option>
            <option value="entertainment">エンターテイメント</option>
            <option value="shopping">ショッピング</option>
          </select>
        </div>

        <div class="w-75 flex flex-col gap-2">
          <div class="w-15 size-6">基準地</div>
          <div class="w-75 flex flex-col">
            <div class="w-42 size-10">
              <input
                type="radio"
                :value="true"
                v-model="useCurrentLocation"
                @change="handleLocationTypeChange('currentPlace')"
                class="radio-text"
              />現在地
              <input
                type="radio"
                :value="false"
                v-model="useCurrentLocation"
                @change="handleLocationTypeChange('other')"
                class="radio-text"
              />その他
            </div>
            <BasicInput
              v-if="!useCurrentLocation"
              v-model="standardPlace"
              placeholder="基準地を入力してください"
            />
          </div>
        </div>

        <div class="w-75 size-17 flex flex-col gap-1">
          <div>距離(m)</div>
          <BasicInput v-model="distance" type="number" />
        </div>

        <div class="w-75 size-17 flex flex-col gap-1">
          <div>候補数</div>
          <BasicInput v-model="candidatesNumber" type="number" />
        </div>
      </div>

      <div class="flex items-center gap-4" v-if="roomCreationStore.isLoading">
        <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
        <span>場所を検索中...</span>
      </div>

      <div v-if="roomCreationStore.error" class="text-red-600 bg-red-50 p-3 rounded-md">
        {{ roomCreationStore.error }}
      </div>

      <BasicButton
        text="次へ"
        right-icon="arrow_foward_inv"
        @click="goNext"
        :disabled="roomCreationStore.isLoading || !genre || (!useCurrentLocation && !standardPlace)"
      />
    </div>
  </div>
</template>

<style>
.header {
  font-weight: bold;
}

.inputform {
  border: 1px solid #000000;
  border-radius: 4px;
}

.radio-text {
  font-size: 10px;
}
</style>
