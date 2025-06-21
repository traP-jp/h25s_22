<template>
  <div class="flex w-full items-start justify-center">
    <div
      class="flex w-[400px] flex-col"
    >
      <div class="box-border w-full flex-grow overflow-y-auto px-12 pt-12">
        <div class="w-full">
          <h1 class="mb-8 h-9 text-xl font-medium leading-9">候補地を希望順に並び替え</h1>

          <div class="space-y-3">
            <div
              v-for="(location, index) in displayLocations"
              :key="location.id"
              draggable="true"
              @dragstart="handleDragStart(index, $event)"
              @dragover="handleDragOver(index, $event)"
              @dragenter="handleDragEnter(index)"
              @drop="handleDrop(index, $event)"
              @dragend="handleDragEnd"
              @touchstart="handleTouchStart(index, $event)"
              @touchmove="handleTouchMove($event)"
              @touchend="handleTouchEnd($event)"
              :class="[
                'bg-gray-50 rounded-lg border border-gray-200 p-4 cursor-move transition-all duration-200 ease-out',
                draggedIndex === index && !isDragging ? 'opacity-60 scale-105 shadow-lg z-10' : '',
                draggedIndex === index && isDragging ? 'opacity-20' : '',
                dragOverIndex === index && draggedIndex !== index ? 'border-blue-400 bg-blue-50' : '',
                'transform select-none'
              ]"
            >
              <div class="flex items-center justify-between">
                <div class="flex items-center">
                  <div
                    class="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center text-white font-bold text-sm mr-4 transition-colors duration-200"
                  >
                    {{ getDisplayRank(location, index) }}
                  </div>
                  <div>
                    <h3 class="font-medium text-gray-800">{{ location.name }}</h3>
                    <p class="text-sm text-gray-500">{{ location.distance }}</p>
                  </div>
                </div>

                <div class="flex items-center space-x-2">

                  <div class="flex flex-col space-y-1">
                    <button
                      @click="moveItem(index, 'up')"
                      :disabled="index === 0"
                      :class="[
                        'p-1 rounded transition-colors',
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
                        'p-1 rounded transition-colors',
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
    </div>
    <div
      v-if="isDragging && draggedIndex !== null && touchPosition.x !== null && touchPosition.y !== null"
      :style="{
        position: 'fixed',
        top: touchPosition.y - 40 + 'px',
        width: '400px',
        zIndex: 1000,
        pointerEvents: 'none'
      }"
      class="bg-gray-50 rounded-lg border border-gray-200 p-4 opacity-80 scale-105 shadow-2xl transform rotate-2"
    >
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <div
            class="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center text-white font-bold text-sm mr-4"
          >
            {{ draggedIndex !== null ? draggedIndex + 1 : 1 }}
          </div>
          <div>
            <h3 class="font-medium text-gray-800">{{ locations[draggedIndex]?.name }}</h3>
            <p class="text-sm text-gray-500">{{ locations[draggedIndex]?.distance }}</p>
          </div>
        </div>

        <div class="flex items-center space-x-2">
          <div class="text-gray-400">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4 8h16M4 16h16"
              />
            </svg>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface Location {
  id: number
  name: string
  distance: string
}

const locations = ref<Location[]>([
  { id: 1, name: 'カフェ SUNNY', distance: '渋谷駅徒歩5分' },
  { id: 2, name: 'イタリアン Bella Vista', distance: '渋谷駅徒歩7分' },
  { id: 3, name: 'ゲームセンター JOYPOLIS', distance: '渋谷駅徒歩2分' },
  { id: 4, name: 'ボウリング場 STRIKE', distance: '渋谷駅徒歩10分' },
  { id: 5, name: 'くそなげーーーーーーーなまえ', distance: '渋谷駅徒歩3分' },
])

// ドラッグアンドドロップの状態管理
const draggedIndex = ref<number | null>(null)
const dragOverIndex = ref<number | null>(null)

// タッチイベント用の状態管理
const touchStartY = ref<number | null>(null)
const touchStartIndex = ref<number | null>(null)
const isDragging = ref(false)
const touchPosition = ref<{ x: number | null; y: number | null }>({ x: null, y: null })

// リアルタイムでドラッグ中の状態を反映する計算プロパティ
const displayLocations = computed(() => {
  if (draggedIndex.value === null || dragOverIndex.value === null) {
    return locations.value
  }

  const result = [...locations.value]
  const draggedItem = result[draggedIndex.value]

  // ドラッグ中のアイテムを削除
  result.splice(draggedIndex.value, 1)

  // 新しい位置に挿入
  const insertIndex = draggedIndex.value < dragOverIndex.value ? dragOverIndex.value - 1 : dragOverIndex.value
  result.splice(insertIndex, 0, draggedItem)

  return result
})

// ドラッグ中は元の順位を表示し、ドラッグ後にのみ新しい順位を表示する関数
const getDisplayRank = (location: Location, currentIndex: number): number => {
  // ドラッグ中でない場合は、現在の表示順位を返す
  if (draggedIndex.value === null || !isDragging.value) {
    return currentIndex + 1
  }

  // ドラッグ中の場合は、元の配列での位置を返す
  const originalIndex = locations.value.findIndex(loc => loc.id === location.id)
  return originalIndex + 1
}

const moveItem = (index: number, direction: 'up' | 'down') => {
  const newLocations = [...locations.value]
  const targetIndex = direction === 'up' ? index - 1 : index + 1

  if (targetIndex >= 0 && targetIndex < locations.value.length) {
    ;[newLocations[index], newLocations[targetIndex]] = [
      newLocations[targetIndex],
      newLocations[index],
    ]
    locations.value = newLocations
    console.log(`Moved item from index ${index} to ${targetIndex}`)
  }
}

// ドラッグアンドドロップイベントハンドラ
const handleDragStart = (index: number, event: DragEvent) => {
  draggedIndex.value = index
  if (event.dataTransfer) {
    event.dataTransfer.effectAllowed = 'move'
    event.dataTransfer.setData('text/plain', index.toString())
  }
}

const handleDragOver = (index: number, event: DragEvent) => {
  event.preventDefault()
  if (event.dataTransfer) {
    event.dataTransfer.dropEffect = 'move'
  }

  if (draggedIndex.value === null || draggedIndex.value === index) {
    return
  }

  // マウスの位置を取得
  const mouseY = event.clientY

  // 対象要素の境界を取得
  const targetElement = event.currentTarget as HTMLElement
  const rect = targetElement.getBoundingClientRect()
  const elementTop = rect.top
  const elementBottom = rect.bottom

  // ドラッグしている要素の中心が、対象要素の上端または下端を越えたかチェック
  let shouldUpdatePosition = false

  if (draggedIndex.value < index) {
    // 下方向にドラッグしている場合：要素の上端を越えたか
    shouldUpdatePosition = mouseY > elementTop
  } else {
    // 上方向にドラッグしている場合：要素の下端を越えたか
    shouldUpdatePosition = mouseY < elementBottom
  }

  if (shouldUpdatePosition) {
    dragOverIndex.value = index
  }
}

const handleDrop = (dropIndex: number, event: DragEvent) => {
  event.preventDefault()

  if (draggedIndex.value !== null && draggedIndex.value !== dropIndex) {
    locations.value = [...displayLocations.value]
  }

  draggedIndex.value = null
  dragOverIndex.value = null
}

const handleDragEnd = () => {
  draggedIndex.value = null
  dragOverIndex.value = null
}

const handleDragEnter = (index: number) => {
  if (draggedIndex.value !== null && index !== draggedIndex.value) {
    dragOverIndex.value = index
  }
}

// タッチイベントハンドラ（モバイル対応）
const handleTouchStart = (index: number, event: TouchEvent) => {
  touchStartIndex.value = index
  touchStartY.value = event.touches[0].clientY
  draggedIndex.value = index
  isDragging.value = true

  // 初期タッチ位置を設定
  touchPosition.value = {
    x: event.touches[0].clientX,
    y: event.touches[0].clientY
  }

  // bodyのスクロールを無効化
  document.body.classList.add('drag-active')

  // スクロールを無効化
  event.preventDefault()
}

const handleTouchMove = (event: TouchEvent) => {
  if (!isDragging.value || touchStartY.value === null || touchStartIndex.value === null) {
    return
  }

  event.preventDefault()

  const currentY = event.touches[0].clientY
  const currentX = event.touches[0].clientX

  // タッチ位置を更新（フローティング要素用）
  touchPosition.value = {
    x: currentX,
    y: currentY
  }

  // 新しいインデックスを直接計算
  const deltaY = currentY - touchStartY.value
  const itemHeight = 92
  const movedItems = Math.round(deltaY / itemHeight)
  let newIndex = touchStartIndex.value + movedItems

  // 境界値チェック
  newIndex = Math.max(0, Math.min(locations.value.length - 1, newIndex))

  // dragOverIndexを更新
  if (newIndex !== dragOverIndex.value) {
    console.log("新しいインデックス:", newIndex)
    dragOverIndex.value = newIndex
  }
}

const handleTouchEnd = (event: TouchEvent) => {
  if (!isDragging.value || draggedIndex.value === null) {
    return
  }

  event.preventDefault()

  // ドロップ処理
  if (dragOverIndex.value !== null && draggedIndex.value !== dragOverIndex.value) {
    locations.value = [...displayLocations.value]
  }

  // bodyのスクロールを復元
  document.body.classList.remove('drag-active')

  // 状態リセット
  isDragging.value = false
  draggedIndex.value = null
  dragOverIndex.value = null
  touchStartY.value = null
  touchStartIndex.value = null
  touchPosition.value = { x: null, y: null }
}
</script>

<style scoped>
/* モバイルでのタッチ操作を改善 */
.cursor-move {
  touch-action: none; /* タッチ操作でのスクロールを無効化 */
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

/* タッチデバイスでのホバー効果を調整 */
@media (hover: hover) {
  .hover\:bg-blue-50:hover {
    background-color: rgb(239 246 255);
  }

  .hover\:text-gray-600:hover {
    color: rgb(75 85 99);
  }
}

/* タッチ時の視覚的フィードバックを改善 */
.transform {
  will-change: transform;
}

/* モバイルでのスクロール防止 */
body.drag-active {
  overflow: hidden;
  touch-action: none;
}
</style>
