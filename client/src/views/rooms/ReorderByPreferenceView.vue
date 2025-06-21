<template>
  <div class="min-h-screen bg-gray-50 py-8">
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="bg-white rounded-lg shadow p-6">
        <h1 class="text-2xl font-bold text-gray-900 mb-4">優先順位設定</h1>
        <p class="text-gray-600">会議室ID: {{ $route.params.room_id }}</p>
        <p class="text-gray-600">このページでは選択肢を優先順位で並び替えます。</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'

const route = useRoute()
// 優先順位設定ページのロジックをここに追加
</script>
