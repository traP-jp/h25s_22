<template>
  <button
    :class="[
      'px-6 py-4 rounded-lg inline-flex justify-center items-center overflow-hidden transition-colors duration-200 ease-in-out',
      sizeClasses,
      variantClasses,
    ]"
    :disabled="disabled"
  >
    <div v-if="leftIcon" class="w-6 h-6 flex items-center justify-center">
      <div v-html="getIconSvg(leftIcon)" class="w-6 h-6"></div>
    </div>

    <div class="px-2 flex justify-center items-center">
      <div
        :class="[
          'justify-center text-base font-medium font-primary leading-normal',
          textColorClasses,
        ]"
      >
        {{ text }}
      </div>
    </div>

    <!-- Right Icon -->
    <div v-if="rightIcon" class="w-6 h-6 flex items-center justify-center">
      <div v-html="getIconSvg(rightIcon)" class="w-6 h-6"></div>
    </div>
  </button>
</template>

<script setup lang="ts">
import { computed, ref, onMounted } from 'vue'

// Props definition
const props = defineProps({
  text: {
    type: String,
    required: true,
  },
  variant: {
    type: String as () => 'primary' | 'secondary',
    default: 'primary',
  },
  size: {
    type: String as () => 'large' | 'medium',
    default: 'large',
  },
  leftIcon: {
    type: String,
    default: '',
  },
  rightIcon: {
    type: String,
    default: '',
  },
  disabled: {
    type: Boolean,
    default: false,
  },
})

const leftIconSvg = ref('')
const rightIconSvg = ref('')

const loadSvg = async (iconName: string): Promise<string> => {
  if (!iconName) return ''
  try {
    const response = await fetch(new URL(`/src/assets/icons/${iconName}.svg`, import.meta.url).href)
    let svgContent = await response.text()

    const fillColor = props.variant === 'primary' ? '#ffffff' : '#38A3EE'
    svgContent = svgContent.replace(/fill="[^"]*"/g, `fill="${fillColor}"`)

    return svgContent
  } catch (error) {
    console.error(`Failed to load icon: ${iconName}`, error)
    return ''
  }
}

const getIconSvg = (iconName: string) => {
  if (iconName === props.leftIcon) return leftIconSvg.value
  if (iconName === props.rightIcon) return rightIconSvg.value
  return ''
}

onMounted(async () => {
  if (props.leftIcon) {
    leftIconSvg.value = await loadSvg(props.leftIcon)
  }
  if (props.rightIcon) {
    rightIconSvg.value = await loadSvg(props.rightIcon)
  }
})

const sizeClasses = computed(() => {
  return props.size === 'large' ? 'w-75' : 'w-36'
})

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

const textColorClasses = computed(() => {
  return props.variant === 'primary' ? 'text-white' : 'text-sky-400'
})
</script>
