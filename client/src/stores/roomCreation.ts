import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { TimeOption, Place, PlaceSearchSettings, DateTimeFormData } from '@/types'
import type { PlaceSearchResult } from '@/services/types'
import { searchNearbyPlaces, createRoom } from '@/services/api'

export const useRoomCreationStore = defineStore('roomCreation', () => {
  // 状態
  const timeOptions = ref<TimeOption[]>([])
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
  const hasPlaceSettings = computed(() => {
    const settings = placeSearchSettings.value
    return (
      settings.genre &&
      (settings.useCurrentLocation || settings.baseLocation) &&
      settings.radius > 0
    )
  })
  const canProceedToDecide = computed(
    () => hasTimeOptions.value && suggestedPlaces.value.length > 0,
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
    const index = selectedPlaces.value.findIndex((p) => p.id === place.id)
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
  const createNewRoom = async (roomName: string, description: string = '') => {
    if (!canProceedToDecide.value) {
      error.value = 'ルーム作成に必要な情報が不足しています'
      return null
    }

    isLoading.value = true
    error.value = null

    try {
      const roomData = {
        name: roomName,
        description,
        timeOptions: timeOptions.value.map((option) => ({
          date: option.date,
          startTime: option.startTime,
          endTime: option.endTime,
        })),
        places: selectedPlaces.value.length > 0 ? selectedPlaces.value : suggestedPlaces.value,
      }

      const result = await createRoom(roomData)

      if (result.success) {
        // 成功時はデータをクリア
        resetAllData()
      }

      return result
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
    clearPlaceSearchSettings()
    suggestedPlaces.value = []
    clearSelectedPlaces()
    error.value = null
  }

  return {
    // 状態
    timeOptions,
    placeSearchSettings,
    suggestedPlaces,
    selectedPlaces,
    isLoading,
    error,
    // 計算されたプロパティ
    hasTimeOptions,
    hasPlaceSettings,
    canProceedToDecide,
    // アクション
    addTimeOption,
    removeTimeOption,
    clearTimeOptions,
    updatePlaceSearchSettings,
    clearPlaceSearchSettings,
    searchPlaces,
    togglePlaceSelection,
    clearSelectedPlaces,
    createNewRoom,
    resetAllData,
  }
})
