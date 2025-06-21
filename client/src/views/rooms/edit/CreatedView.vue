<script setup lang="ts">
// 作成完了ページのロジックをここに追加
import BasicButton from '@/components/BasicButton.vue'
import { ref } from 'vue'

const room_id = ref('sampleID')

function idcopy(id: string) {
  navigator.clipboard.writeText(id)
}

function backhome() {
  location.href = 'https://google.com'
}

function sharelink() {}
</script>

<template>
  <div class="container">
    <div class="flex flex-col gap-4 justify-center items-center">
      <h1 class="w-75 size-7">ルームを作成しました！</h1>
      <div class="roomid-box w-75 size-30 flex gap-4 flex flex-col justify-center items-center">
        <div class="box-header w-64 size-8">ルームID</div>
        <div class="size-6 w-64 flex">
          <div class="idtext flex-auto">{{ room_id }}</div>
          <button class="copybutton w-6" @click="idcopy(room_id)"><img src="@/assets/icons/content_copy.svg"></button>
        </div>
      </div>
      <div class="flex gap-3 flex-wrap justify-center items-center">
        <BasicButton text="ホームに戻る" left-icon="home_inv" @click="backhome" />
        <BasicButton
          text="リンクをシェア"
          variant="secondary"
          left-icon="link-variant"
          @click="sharelink"
        />
      </div>
    </div>
  </div>
</template>

<style>
.roomid-box {
  background-color: #ffffff;
  grid-row: 4/5;
  grid-column: 2/3;
  border-radius: 10px;
  box-shadow: 1px 1px 1px #eeeeee inset;
}

.box-header {
  text-align: center;
}

.idtext {
  color: #4a4a4a;
}

.copybutton {
  cursor: pointer;
  text-align: end;
}

.home-button {
  background-color: #38a3ee;
  cursor: pointer;
  border-radius: 10px;
}

.home-button-text {
  color: #ffffff;
}

.share-button {
  border: 2px solid #38a3ee;
  cursor: pointer;
  border-radius: 10px;
}

.share-button-text {
  color: #38a3ee;
}
</style>
