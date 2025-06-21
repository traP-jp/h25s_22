import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Room } from '@/types'

export const useRoomStore = defineStore('room', () => {
  // 状態
  const rooms = ref<Room[]>([])
  const currentRoom = ref<Room | null>(null)

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

  return {
    // 状態
    rooms,
    currentRoom,
    // 計算されたプロパティ
    roomCount,
    // アクション
    addRoom,
    updateRoom,
    deleteRoom,
    setCurrentRoom,
    getCurrentRoom,
  }
})
