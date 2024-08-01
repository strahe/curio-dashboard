<script setup>
import LogoLight from './LogoLight.vue'
import LogoDark from './LogoDark.vue'
import { useCustomizerStore } from '@/stores/customizer'

const customizer = useCustomizerStore()

</script>
<template>
  <LogoLight v-if="customizer.dark" />
  <LogoDark v-else />
</template>
