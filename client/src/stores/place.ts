import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Place, PlaceVote } from '@/types'

export const usePlaceStore = defineStore('place', () => {
  // 状態
  const places = ref<Place[]>([])
  const votes = ref<PlaceVote[]>([])
  const selectedPlaces = ref<string[]>([])

  // 計算されたプロパティ
  const placeCount = computed(() => places.value.length)

  const getPlaceVotes = computed(() => (placeId: string) => {
    return votes.value.filter((vote) => vote.placeId === placeId)
  })

  const getPlaceAverageScore = computed(() => (placeId: string) => {
    const placeVotes = votes.value.filter((vote) => vote.placeId === placeId)
    if (placeVotes.length === 0) return 0
    const totalScore = placeVotes.reduce((sum, vote) => sum + vote.score, 0)
    return totalScore / placeVotes.length
  })

  // アクション
  const addPlace = (place: Place) => {
    places.value.push(place)
  }

  const updatePlace = (placeId: string, updates: Partial<Place>) => {
    const index = places.value.findIndex((place) => place.id === placeId)
    if (index !== -1) {
      places.value[index] = { ...places.value[index], ...updates }
    }
  }

  const deletePlace = (placeId: string) => {
    const index = places.value.findIndex((place) => place.id === placeId)
    if (index !== -1) {
      places.value.splice(index, 1)
    }
    // 関連する投票も削除
    votes.value = votes.value.filter((vote) => vote.placeId !== placeId)
  }

  const addVote = (vote: PlaceVote) => {
    // 既存の投票があれば更新、なければ追加
    const existingVoteIndex = votes.value.findIndex(
      (v) => v.placeId === vote.placeId && v.userId === vote.userId,
    )

    if (existingVoteIndex !== -1) {
      votes.value[existingVoteIndex] = vote
    } else {
      votes.value.push(vote)
    }
  }

  const togglePlaceSelection = (placeId: string) => {
    const index = selectedPlaces.value.indexOf(placeId)
    if (index !== -1) {
      selectedPlaces.value.splice(index, 1)
    } else {
      selectedPlaces.value.push(placeId)
    }
  }

  const clearPlaceSelection = () => {
    selectedPlaces.value = []
  }

  return {
    // 状態
    places,
    votes,
    selectedPlaces,
    // 計算されたプロパティ
    placeCount,
    getPlaceVotes,
    getPlaceAverageScore,
    // アクション
    addPlace,
    updatePlace,
    deletePlace,
    addVote,
    togglePlaceSelection,
    clearPlaceSelection,
  }
})
