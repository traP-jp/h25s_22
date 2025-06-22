<script setup lang="ts">
// 作成完了ページのロジックをここに追加
import BasicButton from '@/components/BasicButton.vue'
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()
const room_id = ref('')

onMounted(() => {
  const roomId = route.query.roomId as string
  if (roomId) {
    room_id.value = roomId
  } else {
    alert('ルームIDが見つかりません。')
    router.push('/')
  }
})

function idcopy(id: string) {
  navigator.clipboard.writeText(id)
}

function backhome() {
  router.push('/')
}

function sharelink(room_id: string) {
  navigator.share({ title: 'ルームをシェア', url: `/rooms/${room_id}/room-participation` })
}
</script>

<template>
  <div class="container">
    <div class="flex flex-col gap-4 justify-center items-center">
      <h1 class="w-75 size-7 header">ルームを作成しました！</h1>
      <div class="roomid-box w-75 size-30 flex gap-4 flex flex-col justify-center items-center">
        <div class="box-header w-64 size-8">ルームID</div>
        <div class="size-6 w-64 flex">
          <div class="idtext flex-auto">{{ room_id }}</div>
          <button class="copybutton w-6" @click="idcopy(room_id)">
            <img src="@/assets/icons/content_copy.svg" />
          </button>
        </div>
      </div>
      <div class="flex gap-3 flex-wrap justify-center items-center">
        <BasicButton text="ホームに戻る" left-icon="home_inv" @click="backhome" />
        <BasicButton
          text="リンクをシェア"
          variant="secondary"
          left-icon="link-variant"
          @click="sharelink(room_id)"
        />
      </div>
    </div>
  </div>
</template>

<style>
.header {
  font-weight: bold;
  font-size: large;
}

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
