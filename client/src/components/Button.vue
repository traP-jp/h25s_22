<template>
  <button
    :class="[
      'px-6 py-4 rounded-lg inline-flex justify-center items-center overflow-hidden transition-colors duration-200 ease-in-out',
      sizeClasses,
      variantClasses
    ]"
    :disabled="disabled"
  >
    <div v-if="leftIcon" class="w-6 h-6 relative">
      <div class="w-6 h-6 left-0 top-0 absolute bg-zinc-300"></div>
      <div :class="[
        'left-[3px] top-[3px] absolute',
        variant === 'primary' ? 'w-4 h-4 bg-white' : 'w-4 h-4 bg-sky-400'
      ]"></div>
    </div>

    <div class="px-2 flex justify-center items-center">
      <div :class="[
        'justify-center text-base font-medium font-primary leading-normal',
        textColorClasses
      ]">
        {{ text }}
      </div>
    </div>

    <div v-if="rightIcon" class="w-6 h-6 relative">
      <div class="w-6 h-6 left-0 top-0 absolute bg-zinc-300"></div>
      <div :class="[
        'left-[5px] top-[5px] absolute',
        variant === 'primary' ? 'w-3.5 h-3.5 bg-white' : 'w-3.5 h-3.5 bg-sky-400'
      ]"></div>
    </div>
  </button>
</template>

<script setup lang="ts">
import { computed } from 'vue'

// Props definition
const props = defineProps({
  text: {
    type: String,
    required: true
  },
  variant: {
    type: String as () => 'primary' | 'secondary',
    default: 'primary'
  },
  size: {
    type: String as () => 'large' | 'medium',
    default: 'large'
  },
  leftIcon: {
    type: Boolean,
    default: false
  },
  rightIcon: {
    type: Boolean,
    default: false
  },
  disabled: {
    type: Boolean,
    default: false
  }
})

// Computed classes based on size
const sizeClasses = computed(() => {
  return props.size === 'large' ? 'w-75' : 'w-36'
})

// Computed classes based on variant
const variantClasses = computed(() => {
  const baseClasses = props.disabled ? 'cursor-not-allowed opacity-50' : 'cursor-pointer'

  if (props.variant === 'primary') {
    const hoverClasses = props.disabled ? '' : 'hover:bg-sky-500 active:bg-sky-600'
    return `bg-sky-400 ${hoverClasses} ${baseClasses}`
  } else {
    const hoverClasses = props.disabled ? '' : 'hover:bg-sky-50 active:bg-sky-100'
    return `bg-white outline outline-1 outline-offset-[-0.50px] outline-sky-400 ${hoverClasses} ${baseClasses}`
  }
})

// Computed text color classes
const textColorClasses = computed(() => {
  return props.variant === 'primary' ? 'text-white' : 'text-sky-400'
})
</script>
