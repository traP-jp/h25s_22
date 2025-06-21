<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Main Content -->
    <div class="px-4 py-8">
      <h1 class="text-xl font-bold text-gray-800 mb-8 text-center">候補地を希望順に並び替え</h1>

      <!-- Priority Section -->
      <div class="mb-6">

        <!-- Location List -->
        <div class="space-y-3">
          <div
            v-for="(location, index) in locations"
            :key="location.id"
            class="bg-white rounded-lg shadow-sm border border-gray-200 p-4"
          >
            <div class="flex items-center justify-between">
              <div class="flex items-center">
                <div
                  :class="`w-8 h-8 ${location.color} rounded-full flex items-center justify-center text-white font-bold text-sm mr-4`"
                >
                  {{ index + 1 }}
                </div>
                <div>
                  <h3 class="font-medium text-gray-800">{{ location.name }}</h3>
                  <p class="text-sm text-gray-500">{{ location.distance }}</p>
                </div>
              </div>

              <div class="flex flex-col space-y-1">
                <button
                  @click="moveItem(index, 'up')"
                  :disabled="index === 0"
                  :class="[
                    'p-1 rounded',
                    index === 0
                      ? 'text-gray-300 cursor-not-allowed'
                      : 'text-blue-500 hover:bg-blue-50',
                  ]"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M5 15l7-7 7 7"
                    />
                  </svg>
                </button>
                <button
                  @click="moveItem(index, 'down')"
                  :disabled="index === locations.length - 1"
                  :class="[
                    'p-1 rounded',
                    index === locations.length - 1
                      ? 'text-gray-300 cursor-not-allowed'
                      : 'text-blue-500 hover:bg-blue-50',
                  ]"
                >
                  <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M19 9l-7 7-7-7"
                    />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

interface Location {
  id: number
  name: string
  distance: string
  color: string
}

const locations = ref<Location[]>([
  { id: 1, name: 'カフェ SUNNY', distance: '渋谷駅徒歩5分', color: 'bg-yellow-500' },
  { id: 2, name: 'イタリアン Bella Vista', distance: '渋谷駅徒歩7分', color: 'bg-gray-400' },
  { id: 3, name: 'ゲームセンター JOYPOLIS', distance: '渋谷駅徒歩2分', color: 'bg-orange-500' },
  { id: 4, name: 'ボウリング場 STRIKE', distance: '渋谷駅徒歩10分', color: 'bg-blue-500' },
  { id: 5, name: '居酒屋 みんなの家', distance: '渋谷駅徒歩3分', color: 'bg-purple-500' },
])

const moveItem = (index: number, direction: 'up' | 'down') => {
  const newLocations = [...locations.value]
  const targetIndex = direction === 'up' ? index - 1 : index + 1

  if (targetIndex >= 0 && targetIndex < locations.value.length) {
    ;[newLocations[index], newLocations[targetIndex]] = [
      newLocations[targetIndex],
      newLocations[index],
    ]
    locations.value = newLocations
  }
}
</script>
