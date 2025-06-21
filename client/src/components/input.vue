<script setup lang="ts">
import { computed, ref, useTemplateRef, onMounted } from 'vue'

defineOptions({
  inheritAttrs: false
})
const {
  error = false,
  errorMessage = '',
  label = '',
  leftIcon = '',
  rightIcon = '',
  supportingText,
  size = 'large'
} = defineProps<{
  disabled?: boolean
  displaysLength?: boolean
  error?: boolean
  errorMessage?: string
  id?: string
  label?: string
  leftIcon?: string
  required?: boolean
  rightIcon?: string
  size?: 'large' | 'medium'
  supportingText?: string
}>()
const emit = defineEmits(['clickRight', 'focusin', 'blur'])
const value = defineModel<string>()

const leftIconSvg = ref('')
const rightIconSvg = ref('')

const loadSvg = async (iconName: string): Promise<string> => {
  if (!iconName) return ''
  try {
    const response = await fetch(new URL(`/src/assets/icons/${iconName}.svg`, import.meta.url).href)
    let svgContent = await response.text()

    // アイコンの色を設定（通常時は灰色）
    const fillColor = '#6B7280'
    svgContent = svgContent.replace(/fill="[^"]*"/g, `fill="${fillColor}"`)

    return svgContent
  } catch (error) {
    console.error(`Failed to load icon: ${iconName}`, error)
    return ''
  }
}

const getIconSvg = (iconName: string) => {
  if (iconName === leftIcon) return leftIconSvg.value
  if (iconName === rightIcon) return rightIconSvg.value
  return ''
}

onMounted(async () => {
  if (leftIcon) {
    leftIconSvg.value = await loadSvg(leftIcon)
  }
  if (rightIcon) {
    rightIconSvg.value = await loadSvg(rightIcon)
  }
})

const displaysError = computed(() => error || errorMessage !== '')
const displaysLeftIcon = computed(() => leftIcon !== '')
const displaysRightIcon = computed(() => rightIcon !== '')
const displaysSupportingText = computed(() => supportingText != null)
const sizeClasses = computed(() => {
  return size === 'large' ? 'w-75' : 'w-36'
})
const input = useTemplateRef('input')
const isFocused = ref<boolean>(false)
const onFocusin = (): void => {
  isFocused.value = true
  emit('focusin')
}
const onBlur = (): void => {
  isFocused.value = false
  emit('blur')
}
const onClickInnerBorder = (e: MouseEvent) => {
  e.stopPropagation()
  e.preventDefault()
  if (!isFocused.value) input.value?.focus()
}
</script>

<template>
  <div :class="sizeClasses">
    <div class="flex flex-col gap-1">
      <span v-if="label != ''" class="flex items-center gap-2">
        <label class="fontstyle-ui-control text-text-primary" :for="id">{{ label }}</label>
        <span v-if="required" class="fontstyle-ui-caption-strong text-red-500">必須</span>
      </span>
      <span
        class="flex rounded border bg-background-primary p-2"
        :class="[
          { 'outline-1': isFocused },
          { 'border-border-secondary outline-text-primary': !displaysError },
          { 'border-red-500 outline-1 outline-red-500': displaysError },
          { 'border-text-primary': isFocused && !displaysError },
          { 'bg-background-secondary': disabled }
        ]"
        @mousedown="onClickInnerBorder"
      >
        <div v-if="displaysLeftIcon" class="w-5 h-5 flex items-center justify-center">
          <div
            v-html="getIconSvg(leftIcon)"
            class="w-5 h-5"
          ></div>
        </div>
        <input
          v-bind="$attrs"
          :id="id"
          ref="input"
          v-model="value"
          :disabled="disabled"
          class="fontstyle-ui-body w-full min-w-0 bg-transparent px-2 text-text-primary outline-none placeholder:text-text-tertiary"
          @focusin="onFocusin"
          @blur="onBlur"
        />
        <span class="inline-flex items-center gap-2">
          <span v-if="displaysLength" class="fontstyle-ui-caption text-text-secondary">{{
            value?.length ?? 0
          }}</span>
          <button v-if="displaysRightIcon" type="button" @click="emit('clickRight')">
            <div class="w-5 h-5 flex items-center justify-center">
              <div
                v-html="getIconSvg(rightIcon)"
                class="w-5 h-5"
              ></div>
            </div>
          </button>
        </span>
      </span>
      <span
        v-if="displaysSupportingText"
        class="fontstyle-ui-caption whitespace-pre-line text-text-secondary"
        >{{ supportingText }}</span
      >
    </div>
    <div v-if="errorMessage != ''" class="mt-2 flex items-start gap-2 text-red-500">
      <div class="w-5 h-5 flex items-center justify-center">
        <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
          <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z"/>
        </svg>
      </div>
      <span class="fontstyle-ui-control min-w-0 break-words">{{ errorMessage }}</span>
    </div>
  </div>
</template>

<style scoped></style>
