<script setup lang="ts">
// 作成完了ページのロジックをここに追加
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
  <div id="container">
    <div class="flex gap-4 flex-wrap">
      <h1 class="w-75 size-7">ルームを作成しました！</h1>
      <div class="roomid-box w-75 size-30 flex gap-4 flex flex-col justify-center items-center">
        <div class="box-header w-64 size-8">ルームID</div>
        <div class="size-6 w-64 flex">
          <div class="idtext flex-auto">{{ room_id }}</div>
          <button class="copybutton w-6" @click="idcopy(room_id)">コ</button>
        </div>
      </div>

      <div class="flex gap-3 flex-wrap">
        <button class="home-button w-75 size-14 flex-none" @click="backhome">
          <div class="home-button-text">ホームに戻る</div>
        </button>
        <button class="share-button w-75 size-14 flex-none" @click="sharelink">
          <div class="share-button-text">リンクをシェア</div>
        </button>
      </div>
    </div>
  </div>
</template>

<style>
#container {
  display: grid;
  grid-template-columns: 50px 300px 50px;
}

.roomid-box {
  background-color: #d9d9d9;
  grid-row: 4/5;
  grid-column: 2/3;
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
