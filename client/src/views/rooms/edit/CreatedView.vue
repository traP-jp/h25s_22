<script setup lang="ts">
// 作成完了ページのロジックをここに追加
import { isJSDocOptionalType } from 'typescript';
import{ref} from 'vue'

const room_id = ref('sampleID')

function idcopy(id: string){
  navigator.clipboard.writeText(id)
}

function backhome(){
 location.href='https://google.com'
}

function sharelink(){

}
</script>

<template>
  <div class="header">ルームを作成しました！</div>
  <br>
  <div class="roomid-box">
    <div class="box-header">ルームID</div>
    <p class="idtext">{{ room_id }}
      <button class="copybutton" @click="idcopy(room_id)">コピーボタン</button>
    </p>
  </div>
  <button class="home-button" @click="backhome">
    <div class="home-button-text">
      ホームに戻る
    </div>
  </button>
  <br>
  <button class="share-button" @click="sharelink">
    <div class="share-button-text">
      リンクをシェア
    </div>
  </button>
</template>

<style>
.header{
  font-size: 28px;
}

.roomid-box{
  background-color: #d9d9d9;
}

.box-header{
  text-align: center;
  font-size: 32px;
}

.idtext{
  color: #4a4a4a;
  font-size: 25px;
}

.copybutton{
  cursor: pointer;
  text-align:end;
}

.home-button{
  background-color: #38a3ee;
  cursor: pointer;
}

.home-button-text{
  font-size: 24px;
  color:#ffffff;
}

.share-button{
  outline-color: #38a3ee;
  outline-width: 2px;
}

.share-button-text{
  font-size: 24px;
  color:#38a3ee
}
</style>