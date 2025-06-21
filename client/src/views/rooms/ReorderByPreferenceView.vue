<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Header -->
    <div class="bg-white shadow-sm">
      <div class="flex items-center justify-between p-4">
        <div class="w-6 h-6 bg-blue-500 rounded"></div>
        <button class="p-2">
          <svg class="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h16M4 18h16"
            />
          </svg>
        </button>
      </div>
    </div>
    <!-- Main Content -->
    <div class="px-4 py-8">
      <h1 class="text-xl font-bold text-gray-800 mb-8 text-center">候補地を希望順に並び替え</h1>

      <!-- Priority Section -->
      <div class="mb-6">
        <div class="flex items-center mb-4">
          <div class="w-5 h-5 bg-green-500 rounded mr-2 flex items-center justify-center">
            <svg class="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
          <span class="text-gray-700 font-medium">あなたの優先順位</span>
        </div>

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

    <!-- Bottom Button -->
    <div class="fixed bottom-0 left-0 right-0 p-4 bg-white border-t border-gray-200">
      <button
        class="w-full bg-blue-500 text-white py-3 rounded-lg font-medium hover:bg-blue-600 transition-colors"
      >
        100人/200人
      </button>
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
