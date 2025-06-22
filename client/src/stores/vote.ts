import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { PostVoteRequest, PlaceVoteData } from '@/services/types'
import { submitVote } from '@/services/api'

export const useVoteStore = defineStore('vote', () => {
  // 状態
  const userName = ref<string>('')
  const selectedTimeIds = ref<string[]>([])
  const placePreferences = ref<PlaceVoteData[]>([])
  const isSubmitting = ref(false)
  const error = ref<string | null>(null)

  // 計算されたプロパティ
  const canSubmitVote = computed(() => {
    return (
      userName.value.trim() !== '' &&
      selectedTimeIds.value.length > 0 &&
      placePreferences.value.length > 0
    )
  })

  // アクション
  const setUserName = (name: string) => {
    userName.value = name
  }

  const setSelectedTimeIds = (timeIds: string[]) => {
    selectedTimeIds.value = timeIds
  }

  const setPlacePreferences = (preferences: PlaceVoteData[]) => {
    placePreferences.value = preferences
  }

  const addPlacePreference = (placeId: string, rank: number) => {
    const existingIndex = placePreferences.value.findIndex((p) => p.id === placeId)
    if (existingIndex !== -1) {
      placePreferences.value[existingIndex].rank = rank
    } else {
      placePreferences.value.push({ id: placeId, rank })
    }
  }

  const removePlacePreference = (placeId: string) => {
    placePreferences.value = placePreferences.value.filter((p) => p.id !== placeId)
  }

  const clearVoteData = () => {
    userName.value = ''
    selectedTimeIds.value = []
    placePreferences.value = []
    error.value = null
  }

  // 投票送信
  const submitVoteData = async (roomId: string) => {
    if (!canSubmitVote.value) {
      error.value = '必要な情報が不足しています'
      return null
    }

    try {
      isSubmitting.value = true
      error.value = null

      const voteRequest: PostVoteRequest = {
        roomID: roomId,
        name: userName.value.trim(),
        timeId: selectedTimeIds.value,
        placeData: placePreferences.value,
      }

      console.log('投票データ送信:', voteRequest)
      const response = await submitVote(voteRequest)
      console.log('投票結果:', response)

      return response
    } catch (err) {
      console.error('投票送信エラー:', err)
      error.value = err instanceof Error ? err.message : '投票の送信に失敗しました'
      throw err
    } finally {
      isSubmitting.value = false
    }
  }

  return {
    // 状態
    userName,
    selectedTimeIds,
    placePreferences,
    isSubmitting,
    error,
    // 計算されたプロパティ
    canSubmitVote,
    // アクション
    setUserName,
    setSelectedTimeIds,
    setPlacePreferences,
    addPlacePreference,
    removePlacePreference,
    clearVoteData,
    submitVoteData,
  }
})
