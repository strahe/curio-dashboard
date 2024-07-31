<script setup lang="ts">
import { shallowRef } from 'vue';

const sources = shallowRef([
  {
    text: 'Referral',
    percent: 20,
    color: 'primary'
  },
  {
    text: 'Bounce',
    percent: 58,
    color: 'secondary'
  },
  {
    text: 'Internet',
    percent: 40,
    color: 'primary'
  },
  {
    text: 'Social',
    percent: 90,
    color: 'primary'
  }
]);
</script>

<template>
  <div class="py-3" v-for="(source, i) in sources" :key="i" :value="source.text">
    <div class="d-flex align-center justify-space-between mb-2">
      <span class="text-caption">{{ source.text }}</span>
      <span class="text-caption">{{ source.percent }}%</span>
    </div>
    <v-progress-linear :model-value="source.percent" aria-label="progressbar" height="5" :color="source.color"></v-progress-linear>
  </div>
</template>
