import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Room } from '@/types'
import type { GetRoomResponse } from '@/services/types'
import { getRoomDetails } from '@/services/api'

export const useRoomStore = defineStore('room', () => {
  // 状態
  const rooms = ref<Room[]>([])
  const currentRoom = ref<Room | null>(null)
  const currentRoomData = ref<GetRoomResponse | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  // 計算されたプロパティ
  const roomCount = computed(() => rooms.value.length)

  // アクション
  const addRoom = (room: Room) => {
    rooms.value.push(room)
  }

  const updateRoom = (roomId: string, updates: Partial<Room>) => {
    const index = rooms.value.findIndex((room) => room.id === roomId)
    if (index !== -1) {
      rooms.value[index] = { ...rooms.value[index], ...updates }
    }
  }

  const deleteRoom = (roomId: string) => {
    const index = rooms.value.findIndex((room) => room.id === roomId)
    if (index !== -1) {
      rooms.value.splice(index, 1)
    }
  }

  const setCurrentRoom = (room: Room | null) => {
    currentRoom.value = room
  }

  const getCurrentRoom = () => {
    return currentRoom.value
  }

  // 新しいAPI用メソッド
  const fetchRoomData = async (roomId: string) => {
    try {
      isLoading.value = true
      error.value = null
      console.log('RoomStore - API呼び出し開始:', roomId)
      
      const data = await getRoomDetails(roomId)
      console.log('RoomStore - API レスポンス:', data)
      
      currentRoomData.value = data
      return data
    } catch (err) {
      console.error('RoomStore - API エラー:', err)
      error.value = err instanceof Error ? err.message : 'ルーム情報の取得に失敗しました'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const clearCurrentRoomData = () => {
    currentRoomData.value = null
    error.value = null
  }

  return {
    // 状態
    rooms,
    currentRoom,
    currentRoomData,
    isLoading,
    error,
    // 計算されたプロパティ
    roomCount,
    // アクション
    addRoom,
    updateRoom,
    deleteRoom,
    setCurrentRoom,
    getCurrentRoom,
    fetchRoomData,
    clearCurrentRoomData,
  }
})
