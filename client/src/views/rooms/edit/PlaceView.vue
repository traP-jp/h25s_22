<script setup lang="ts">
// 場所設定ページのロジックをここに追加
import { ref } from 'vue'
import BasicButton from '@/components/BasicButton.vue'
import BasicInput from '@/components/BasicInput.vue'

const distance = ref(0)
const standardPlace = ref('')
const candidatesNumber = ref(0)

function goNext() {
  location.href = 'https://google.com'
}
</script>

<template>
  <div class="flex flex-col gap-4 justify-center items-center">
    <h1 class="header w-75 size-9">場所を設定</h1>
    <div class="w-75 flex flex-col gap-9">
      <div class="w-75 flex flex-col gap-3">
        <div class="w-75 size-18 flex flex-col gap-2">
          <div>ジャンル</div>
          <select name="genre" class="inputform size-10 w-75">
            <option value=""></option>
            <option value="ジャンル1">ジャンル1</option>
            <option value="ジャンル2">ジャンル2</option>
          </select>
        </div>
        <div class="w-75 flex flex-col gap-2">
          <div class="w-15 size-6">基準地</div>
          <div class="w-75 flex flex-col">
            <div class="w-42 size-10">
              <input
                type="radio"
                v-model="standardPlace"
                value="currentPlace"
                class="radio-text"
              />現在地
              <input type="radio" v-model="standardPlace" value="" class="radio-text" />その他
            </div>
            <BasicInput v-if="standardPlace !== 'currentPlace'" v-model="standardPlace" />
          </div>
        </div>
        <div class="w-75 size-17 flex flex-col gap-1">
          <div>距離(m)</div>
          <BasicInput v-model="distance" />
        </div>
        <div class="w-75 size-17 flex flex-col gap-1">
          <div>候補数</div>
          <BasicInput v-model="candidatesNumber" />
        </div>
      </div>
      <BasicButton text="次へ" right-icon="arrow_foward_inv" @click="goNext"></BasicButton>
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
