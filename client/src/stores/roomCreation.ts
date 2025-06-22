import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { TimeOption, PlaceSearchSettings, DateTimeFormData } from '@/types'
import type {
  PlaceSearchResult,
  NewCreateRoomRequest,
  TimeOption as ApiTimeOption,
} from '@/services/types'
import { searchNearbyPlaces, createRoom, createNewRoom as apiCreateNewRoom } from '@/services/api'

export const useRoomCreationStore = defineStore('roomCreation', () => {
  // 状態
  const timeOptions = ref<TimeOption[]>([])
  const roomTitle = ref<string>('')
  const placeSearchSettings = ref<PlaceSearchSettings>({
    genre: '',
    baseLocation: '',
    useCurrentLocation: false,
    radius: 1000,
    maxCandidates: 5,
  })
  const suggestedPlaces = ref<PlaceSearchResult[]>([])
  const selectedPlaces = ref<PlaceSearchResult[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  // 計算されたプロパティ
  const hasTimeOptions = computed(() => timeOptions.value.length > 0)
  const hasRoomTitle = computed(() => roomTitle.value.trim() !== '')
  const hasPlaceSettings = computed(() => {
    const settings = placeSearchSettings.value
    return (
      settings.genre &&
      (settings.useCurrentLocation || settings.baseLocation) &&
      settings.radius > 0
    )
  })
  const canProceedToDecide = computed(
    () => hasTimeOptions.value && hasRoomTitle.value && suggestedPlaces.value.length > 0,
  )

  // アクション - 日時関連
  const addTimeOption = (formData: DateTimeFormData) => {
    const newTimeOption: TimeOption = {
      id: `time_${Date.now()}`,
      date: formData.date,
      startTime: formData.startTime,
      endTime: formData.endTime,
    }
    timeOptions.value.push(newTimeOption)
  }

  const removeTimeOption = (timeOptionId: string) => {
    const index = timeOptions.value.findIndex((option) => option.id === timeOptionId)
    if (index !== -1) {
      timeOptions.value.splice(index, 1)
    }
  }

  const clearTimeOptions = () => {
    timeOptions.value = []
  }

  // アクション - ルームタイトル関連
  const setRoomTitle = (title: string) => {
    roomTitle.value = title
  }

  const clearRoomTitle = () => {
    roomTitle.value = ''
  }

  // アクション - 場所設定関連
  const updatePlaceSearchSettings = (settings: Partial<PlaceSearchSettings>) => {
    placeSearchSettings.value = { ...placeSearchSettings.value, ...settings }
  }

  const clearPlaceSearchSettings = () => {
    placeSearchSettings.value = {
      genre: '',
      baseLocation: '',
      useCurrentLocation: false,
      radius: 1000,
      maxCandidates: 5,
    }
  }

  // アクション - 場所検索
  const searchPlaces = async () => {
    if (!hasPlaceSettings.value) {
      error.value = '場所検索の設定が不完全です'
      return
    }

    isLoading.value = true
    error.value = null

    try {
      const searchParams = {
        name: placeSearchSettings.value.useCurrentLocation
          ? '現在地'
          : placeSearchSettings.value.baseLocation,
        radius: placeSearchSettings.value.radius,
        type: placeSearchSettings.value.genre,
        maxSize: placeSearchSettings.value.maxCandidates,
        language: 'ja',
      }
      console.log('検索パラメータ:', searchParams)
      const results = await searchNearbyPlaces(searchParams)
      suggestedPlaces.value = results
    } catch (err) {
      error.value = err instanceof Error ? err.message : '場所検索でエラーが発生しました'
      suggestedPlaces.value = []
    } finally {
      isLoading.value = false
    }
  }

  // アクション - 選択された場所の管理
  const togglePlaceSelection = (place: PlaceSearchResult) => {
    const index = selectedPlaces.value.findIndex((p) => p.placeID === place.placeID)
    if (index !== -1) {
      selectedPlaces.value.splice(index, 1)
    } else {
      selectedPlaces.value.push(place)
    }
  }

  const clearSelectedPlaces = () => {
    selectedPlaces.value = []
  }

  // アクション - ルーム作成
  const createNewRoom = async (description: string = '') => {
    if (!canProceedToDecide.value) {
      error.value = 'ルーム作成に必要な情報が不足しています'
      return null
    }

    isLoading.value = true
    error.value = null

    try {
      // 新しいAPI仕様に合わせてデータを構築
      const timeOptionsForApi: ApiTimeOption[] = timeOptions.value.map((option) => ({
        start_time: `${option.date}T${option.startTime}:00Z`,
        end_time: `${option.date}T${option.endTime}:00Z`,
      }))

      // 使用する場所のPlace IDを取得
      const placesToUse =
        selectedPlaces.value.length > 0 ? selectedPlaces.value : suggestedPlaces.value
      const placeIds = placesToUse.map((place) => place.placeID)

      // 中心位置として最初の場所を使用（または実際の中心位置を計算）
      const centerPlaceId = placesToUse.length > 0 ? placesToUse[0].placeID : ''

      const roomData: NewCreateRoomRequest = {
        name: roomTitle.value,
        time_options: timeOptionsForApi,
        place_options: placeIds,
        center_place_id: centerPlaceId,
        radius: placeSearchSettings.value.radius,
        place_max: placeSearchSettings.value.maxCandidates,
      }

      console.log('新しいAPI用ルームデータ:', roomData)
      const result = await apiCreateNewRoom(roomData)

      if (result.room_id) {
        // 成功時はデータをクリア
        resetAllData()
        return { success: true, id: result.room_id, message: 'ルームが正常に作成されました' }
      }

      return { success: false, message: 'ルーム作成に失敗しました' }
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'ルーム作成でエラーが発生しました'
      return null
    } finally {
      isLoading.value = false
    }
  }

  // アクション - データリセット
  const resetAllData = () => {
    clearTimeOptions()
    clearRoomTitle()
    clearPlaceSearchSettings()
    suggestedPlaces.value = []
    clearSelectedPlaces()
    error.value = null
  }

  return {
    // 状態
    timeOptions,
    roomTitle,
    placeSearchSettings,
    suggestedPlaces,
    selectedPlaces,
    isLoading,
    error,
    // 計算されたプロパティ
    hasTimeOptions,
    hasRoomTitle,
    hasPlaceSettings,
    canProceedToDecide,
    // アクション
    addTimeOption,
    removeTimeOption,
    clearTimeOptions,
    setRoomTitle,
    clearRoomTitle,
    updatePlaceSearchSettings,
    clearPlaceSearchSettings,
    searchPlaces,
    togglePlaceSelection,
    clearSelectedPlaces,
    createNewRoom,
    resetAllData,
  }
})
