<script setup lang="ts">
import { ref } from 'vue';
// icons
import { FullscreenOutlined, FullscreenExitOutlined } from '@ant-design/icons-vue';

const fullscreen = ref(false);

function toggleFullscreen() {
  fullscreen.value = !fullscreen.value;
  if (document && !document.fullscreenElement) {
    document.documentElement.requestFullscreen();
  } else if (document.exitFullscreen) {
    document.exitFullscreen();
  }
}
</script>
<template>
  <div>
    <v-btn @click="toggleFullscreen" icon class="text-secondary ml-sm-2 ml-1" color="darkText" rounded="sm" size="small">
      <FullscreenOutlined v-if="!fullscreen" :style="{ fontSize: '16px' }" />
      <FullscreenExitOutlined v-if="fullscreen" :style="{ fontSize: '16px' }" />
    </v-btn>
  </div>
</template>
