<script setup>
import { computed } from 'vue'
import LogoLight from './LogoLight.vue'
import LogoDark from './LogoDark.vue'
import { useCustomizerStore } from '@/stores/customizer'

const customizer = useCustomizerStore()

const dark = computed(() => {
    return customizer.actTheme === 'DarkDefaultTheme'
})
</script>
<template>
  <LogoLight v-if="dark" />
  <LogoDark v-else />
</template>
