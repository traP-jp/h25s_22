<template>
  <div class="px-4 py-6 sm:px-0">
    <div class="border-4 border-dashed border-gray-200 rounded-lg p-8">
      <div class="text-center">
        <h1 class="text-4xl font-bold text-gray-900 mb-4">Welcome to our App</h1>
        <p class="text-lg text-gray-600 mb-8">Button Component Examples</p>

        <div class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-800 mb-4">Primary Buttons</h2>
          <div class="flex flex-col items-center space-y-4">
            <BasicButton
              text="Primary Large"
              variant="primary"
              size="large"
              left-icon="home_inv"
              right-icon="home"
            />
            <BasicButton text="Primary Medium" variant="primary" size="medium" left-icon="home" />
            <BasicButton text="Primary No Icons" variant="primary" size="large" />
          </div>
        </div>

        <div class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-800 mb-4">Input Components</h2>
          <div class="flex flex-col items-center space-y-4 max-w-md mx-auto">
            <BasicInput
              v-model="inputValue1"
              label="基本的な入力"
              placeholder="テキストを入力してください"
              left-icon="home"
              right-icon="home_inv"
            />
            <BasicInput
              v-model="inputValue2"
              label="必須項目"
              placeholder="必須項目です"
              :required="true"
              left-icon="home"
              supporting-text="これは必須項目です"
            />
            <BasicInput
              v-model="inputValue3"
              label="エラー表示"
              placeholder="エラーがある場合"
              :error="true"
              error-message="入力内容に問題があります"
              right-icon="home_inv"
            />
            <BasicInput
              v-model="inputValue4"
              label="文字数表示"
              placeholder="文字数が表示されます"
              :displays-length="true"
              left-icon="home_inv"
            />
          </div>
        </div>

        <div class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-800 mb-4">Secondary Buttons</h2>
          <div class="flex flex-col items-center space-y-4">
            <!-- Large secondary with both icons -->
            <BasicButton
              text="Secondary Large"
              variant="secondary"
              size="large"
              left-icon="home"
              right-icon="home_inv"
            />
            <BasicButton
              text="Secondary Medium"
              variant="secondary"
              size="medium"
              right-icon="home"
            />
            <BasicButton text="Secondary No Icons" variant="secondary" size="large" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import BasicButton from '@/components/BasicButton.vue'
import BasicInput from '@/components/BasicInput.vue'

const inputValue1 = ref('')
const inputValue2 = ref('')
const inputValue3 = ref('')
const inputValue4 = ref('')
</script>
