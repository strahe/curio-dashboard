<script setup>
import { computed } from 'vue';
import LogoLight from './LogoLight.vue';
import LogoDark from './LogoDark.vue';
import { useCustomizerStore } from '@/stores/customizer';

const customizer = useCustomizerStore();

const dark = computed(() => {
  if (
    customizer.actTheme === 'DarkDefaultTheme' ||
    customizer.actTheme === 'DarkTheme1' ||
    customizer.actTheme === 'DarkTheme6' ||
    customizer.actTheme === 'DarkTheme7' ||
    customizer.actTheme === 'DarkTheme8' ||
    customizer.actTheme === 'DarkTheme5' ||
    customizer.actTheme === 'DarkTheme2' ||
    customizer.actTheme === 'DarkTheme3' ||
    customizer.actTheme === 'DarkTheme4'
  ) {
    return true;
  } else {
    return false;
  }
});
</script>
<template>
  <LogoLight v-if="dark" />
  <LogoDark v-else />
</template>
