<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

// フォームの入力値を保持する変数
const selectedDate = ref('')
const startTime = ref('')
const endTime = ref('')

// ボタンのクリックイベント用の関数
const handleAdd = () => {
  if (!selectedDate.value || !startTime.value || !endTime.value) {
    alert('すべての日付と時刻を入力してください。')
    return
  }
  console.log('追加されたデータ:', {
    date: selectedDate.value,
    start: startTime.value,
    end: endTime.value,
  })
  alert('スケジュール候補が追加されました！')
}

const handleNext = () => {
  router.push('/rooms/edit/place')
}
</script>

<template>
  <div class="page-wrapper">
    <div class="smartphone-view">
      <header class="form-header">
        <svg
          class="hamburger-menu"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
          />
        </svg>
      </header>
      <div class="form-body">
        <div class="form-container">
          <h1 class="title">日付を入力</h1>
          <div class="form-group date-group">
            <label for="date-input">日付</label>
            <input type="date" id="date-input" v-model="selectedDate" required />
          </div>
          <div class="time-inputs">
            <div class="form-group">
              <label for="start-time">開始時刻</label>
              <input type="time" id="start-time" v-model="startTime" required />
            </div>
            <div class="form-group">
              <label for="end-time">終了時刻</label>
              <input type="time" id="end-time" v-model="endTime" required />
            </div>
          </div>
          <div class="button-group">
            <button class="btn btn-secondary" @click="handleAdd">追加 ＋</button>
            <button class="btn btn-primary" @click="handleNext">次へ →</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/*
  このビュー全体を囲むラッパー
  flexを使い、中のスマホビューを中央に配置
*/
.page-wrapper {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: flex-start; /* 上寄せで配置 */
  padding: 2rem 0; /* 上下に少し余白 */
}

/*
  Figmaで指定された400x800のコンテナ
*/
.smartphone-view {
  width: 400px;
  height: 800px;
  background-color: #ffffff;
  border: 1px solid #e5e7eb;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column; /* 中身を縦に並べる */
  overflow: hidden;
}

/*
  ヘッダー
  高さ38px
*/
.form-header {
  width: 100%;
  height: 38px;
  background-color: #f8f9fa;
  border-bottom: 1px solid #dee2e6;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  padding-right: 20px;
  box-sizing: border-box;
}

.hamburger-menu {
  width: 24px;
  height: 24px;
}

/*
  フォーム本体を囲むエリア
  - ヘッダーとの間に48pxの余白
  - 左右に50pxの余白を作り、中のコンテンツ幅を300pxにする
*/
.form-body {
  padding: 48px 50px 0 50px;
  width: 100%;
  box-sizing: border-box;
}

.form-container {
  width: 100%; /* 親(.form-body)の幅(300px)いっぱいに広がる */
}

.title {
  font-size: 20px;
  font-weight: 500;
  height: 36px;
  line-height: 36px;
  margin: 0;
  margin-bottom: 16px;
  -webkit-text-stroke: 0.3px;
}

.form-group {
  width: 100%;
}

.form-group.date-group {
  height: 68px;
  margin-bottom: 12px;
  box-sizing: border-box;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-size: 14px;
  font-weight: 500;
  height: 20px;
  line-height: 20px;
}

.form-group input {
  width: 100%;
  height: 40px;
  padding: 0 12px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  font-size: 1rem;
  box-sizing: border-box;
}

.time-inputs {
  display: flex;
  gap: 12px;
  height: 69px;
  margin-bottom: 36px;
  box-sizing: border-box;
}

.button-group {
  display: flex;
  gap: 12px;
  height: 48px;
}

.btn {
  width: 144px;
  height: 48px;
  padding: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 6px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  border: 1px solid transparent;
  box-sizing: border-box;
}

.btn-primary {
  background-color: #3b82f6;
  color: white;
}
.btn-primary:hover {
  background-color: #2563eb;
}
.btn-secondary {
  background-color: white;
  color: #3b82f6;
  border-color: #3b82f6;
}
.btn-secondary:hover {
  background-color: #eff6ff;
}

/*
  ↓↓↓ 変更点2: 以下のスタイルを追加 ↓↓↓
  日付・時刻入力欄のプレースホルダーを非表示にする設定
*/

/* まず、デフォルトの文字色を透明に設定します */
.form-group input[type='date'],
.form-group input[type='time'] {
  color: transparent;
}

/*
  そして、以下の条件のときだけ文字色を元に戻します
  - :focus      … ユーザーが入力欄をクリックしているとき
  - :valid      … 値が正しく入力されたとき
*/
.form-group input[type='date']:focus,
.form-group input[type='time']:focus,
.form-group input[type='date']:valid,
.form-group input[type='time']:valid {
  color: #374151; /* inputにもともと設定していた文字色 */
}
</style>